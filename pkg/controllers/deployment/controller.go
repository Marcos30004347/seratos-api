package deploymentcontroller

import (
	"fmt"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	appsinformers "k8s.io/client-go/informers/apps/v1"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	appslisters "k8s.io/client-go/listers/apps/v1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/wait"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	clientset "github.com/Marcos30004347/seratos-api/pkg/generated/clientset/versioned"
)

const controllerAgentName = "deployment-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Deployment is synced
	SuccessSynced = "Synced"

	// ErrResourceExists is used as part of the Event 'reason' when a Deployment fails
	// to sync due to a Deployment of the same name already existing.
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by Deployment"

	// MessageResourceSynced is the message used for an Event fired when a Deployment
	// is synced successfully
	MessageResourceSynced = "Deployment synced successfully"
)

// Controller is the controller implementation for Deployment resources
type Controller struct {
	// kubeclientset is a standard kubernetes clientset
	kubeclientset kubernetes.Interface
	// seratosclientset is a clientset for our own API group
	seratosclientset clientset.Interface

	deploymentsLister appslisters.DeploymentLister
	deploymentsSynced cache.InformerSynced
	deploymentLister  appslisters.DeploymentLister
	synced            cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
}

// enqueueMicroservice takes a Deployment resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than Deployment.
func (c *Controller) enqueueDeployment(obj interface{}) {
	var key string
	var err error
	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}
	c.workqueue.Add(key)
}

// handleObject will take any resource implementing metav1.Object and attempt
// to find the Deployment resource that 'owns' it. It does this by looking at the
// objects metadata.ownerReferences field for an appropriate OwnerReference.
// It then enqueues that Deployment resource to be processed. If the object does not
// have an appropriate OwnerReference, it will simply be skipped.
func (c *Controller) handleObject(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		// klog.V(4).Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}

	// klog.V(4).Infof("Processing object: %s", object.GetName())
	if ownerRef := metav1.GetControllerOf(object); ownerRef != nil {
		// If this object is not owned by a Deployment, we should not do anything more
		// with it.
		if ownerRef.Kind != "Deployment" {
			return
		}

		deployment, err := c.deploymentLister.Deployments(object.GetNamespace()).Get(ownerRef.Name)
		if err != nil {
			// klog.V(4).Infof("ignoring orphaned object '%s' of deployment '%s'", object.GetSelfLink(), ownerRef.Name)
			return
		}

		c.enqueueDeployment(deployment)
		return
	}
}

// microserviceControllerHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the Deployment resource
// with the current status of the resource.
func (c *Controller) microserviceControllerHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the Deployment resource with this namespace/name
	deployment, err := c.deploymentLister.Deployments(namespace).Get(name)

	if err != nil {
		// The Deployment resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("deployment '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	c.recorder.Event(deployment, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the microserviceControllerHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool
		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}
		// Run the microserviceControllerHandler, passing it the namespace/name string of the
		// Deployment resource to be synced.
		if err := c.microserviceControllerHandler(key); err != nil {
			// Put the item back on the workqueue to handle any transient errors.
			c.workqueue.AddRateLimited(key)
			return fmt.Errorf("error syncing '%s': %s, requeuing", key, err.Error())
		}
		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		// klog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	// klog.Info("Starting Deployment controller")

	// Wait for the caches to be synced before starting workers
	// klog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.deploymentsSynced, c.synced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	// klog.Info("Starting workers")
	// Launch two workers to process Deployment resources
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	// klog.Info("Started workers")
	<-stopCh
	// klog.Info("Shutting down workers")

	return nil
}

// NewController returns a new deployment controller
func NewController(
	kubeclientset kubernetes.Interface,
	seratosclientset clientset.Interface,
	deploymentInformer appsinformers.DeploymentInformer,
) *Controller {
	// Create event broadcaster
	// Add deployment-controller types to the default Kubernetes Scheme so Events can be
	// logged for deployment-controller types.

	// THIS MAYBE SHOULD NOT BE COMMENTED
	// utilruntime.Must(bazscheme.AddToScheme(scheme.Scheme))

	// klog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})
	controller := &Controller{
		kubeclientset:     kubeclientset,
		seratosclientset:  seratosclientset,
		deploymentsLister: deploymentInformer.Lister(),
		deploymentsSynced: deploymentInformer.Informer().HasSynced,
		deploymentLister:  deploymentInformer.Lister(),
		synced:            deploymentInformer.Informer().HasSynced,
		workqueue:         workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Microservices"),
		recorder:          recorder,
	}

	// klog.Info("Setting up event handlers")
	// Set up an event handler for when Deployment resources change
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueDeployment,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueDeployment(new)
		},
	})
	// Set up an event handler for when Deployment resources change. This
	// handler will lookup the owner of the given Deployment, and if it is
	// owned by a Deployment resource will enqueue that Deployment resource for
	// processing. This way, we don't need to implement custom logic for
	// handling Deployment resources. More info on this pattern:
	// https://github.com/kubernetes/community/blob/8cafef897a22026d42f5e5bb3f104febe7e29830/contributors/devel/controllers.md
	deploymentInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.handleObject,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*appsv1.Deployment)
			oldDepl := old.(*appsv1.Deployment)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			controller.handleObject(new)
		},
		DeleteFunc: controller.handleObject,
	})

	return controller
}
