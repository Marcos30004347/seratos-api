package controller

import (
	"fmt"
	"time"

	corev1 "k8s.io/api/core/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	v1beta1 "github.com/Marcos30004347/seratos-api/pkg/apis/seratos/v1beta1"
	"github.com/Marcos30004347/seratos-api/pkg/apiserver"
	listers "github.com/Marcos30004347/seratos-api/pkg/generated/listers/seratos/v1beta1"
	"github.com/Marcos30004347/seratos-api/pkg/runtime"
)

// Controller is the controller implementation for Microservice resources
type Controller struct {
	kubeRuntime        *runtime.KubeRuntime
	seratosRuntime     *runtime.SeratosRuntime
	microserviceLister listers.MicroserviceLister
	sidecarsLister     listers.SidecarLister
	synced             cache.InformerSynced
	queue              workqueue.RateLimitingInterface
	recorder           record.EventRecorder
}

const (
	// CREATED value is used inside the the Microservice Event to show that a Event is of CREATED type.
	CREATED = "CREATED"

	// UPDATED is used inside the the Microservice Event to show that a Event is of UPDATED type.
	UPDATED = "UPDATED"

	// DELETED is used inside the the Microservice Event to show that a Event is of DELETED type.
	DELETED = "DELETED"
)

// MicroserviceEvent is the Event object for Microservice Resources.
type MicroserviceEvent struct {
	EventType string
	Created   *MicroserviceCreatedEvent
	Updated   *MicroserviceUpdatedEvent
	Deleted   *MicroserviceDeletedEvent
}

// MicroserviceCreatedEvent is the Event dispatched when a Microservice is Created
type MicroserviceCreatedEvent struct {
	Microservice *v1beta1.Microservice
}

// MicroserviceUpdatedEvent is the Event dispatched when a Microservice is Updated
type MicroserviceUpdatedEvent struct {
	NewMicroservice *v1beta1.Microservice
	OldMicroservice *v1beta1.Microservice
}

// MicroserviceDeletedEvent is the Event dispatched when a Microservice is Deleted
type MicroserviceDeletedEvent struct {
	Microservice *v1beta1.Microservice
}

func (c *Controller) enqueueResource(obj interface{}) {
	var key string
	var err error

	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		utilruntime.HandleError(err)
		return
	}

	c.queue.Add(key)
}

func (c *Controller) controllerHandler(event MicroserviceEvent) error {
	switch event.EventType {
	case CREATED:
		return microserviceCreated(c, event.Created.Microservice)
	case UPDATED:
		return microserviceUpdated(c, event.Updated.OldMicroservice, event.Updated.NewMicroservice)
	case DELETED:
		return microserviceDeleted(c, event.Deleted.Microservice)
	default:
		return nil
	}
}

func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.queue.Get()

	if shutdown {
		return false
	}

	err := func(obj interface{}) error {

		defer c.queue.Done(obj)
		var event MicroserviceEvent
		var ok bool

		if event, ok = obj.(MicroserviceEvent); !ok {

			c.queue.Forget(obj)
			utilruntime.HandleError(fmt.Errorf("expected string in queue but got %#v", obj))
			return nil
		}

		if err := c.controllerHandler(event); err != nil {
			c.queue.AddRateLimited(event)
			return fmt.Errorf("error syncing '%#v': %s, requeuing", event, err.Error())
		}
		c.queue.Forget(obj)
		return nil
	}(obj)

	if err != nil {
		utilruntime.HandleError(err)
		return true
	}

	return true
}

func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// Run tuns the controller
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer utilruntime.HandleCrash()
	defer c.queue.ShutDown()

	if ok := cache.WaitForCacheSync(stopCh, c.synced, c.synced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	<-stopCh

	return nil
}

// NewController creates a new controller
func NewController(
	kube *runtime.KubeRuntime,
	seratosRuntime *runtime.SeratosRuntime,
) *Controller {
	microservicesAPI := seratosRuntime.Informers.Seratos().V1beta1().Microservices()
	sidecarsAPI := seratosRuntime.Informers.Seratos().V1beta1().Sidecars()

	eventBroadcaster := record.NewBroadcaster()

	eventBroadcaster.StartStructuredLogging(0)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kube.Clientset.CoreV1().Events("")})

	recorder := eventBroadcaster.NewRecorder(apiserver.GetRuntimeScheme(), corev1.EventSource{Component: "microservices-controller"})

	controller := &Controller{
		kubeRuntime:        kube,
		seratosRuntime:     seratosRuntime,
		microserviceLister: microservicesAPI.Lister(),
		sidecarsLister:     sidecarsAPI.Lister(),
		synced:             microservicesAPI.Informer().HasSynced,
		queue:              workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Microservices"),
		recorder:           recorder,
	}

	microservicesAPI.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			event := MicroserviceEvent{
				EventType: CREATED,
				Created: &MicroserviceCreatedEvent{
					Microservice: obj.(*v1beta1.Microservice),
				},
				Updated: nil,
				Deleted: nil,
			}

			controller.queue.Add(event)
		},

		UpdateFunc: func(old interface{}, new interface{}) {
			event := MicroserviceEvent{
				EventType: UPDATED,
				Created:   nil,
				Updated: &MicroserviceUpdatedEvent{
					OldMicroservice: old.(*v1beta1.Microservice),
					NewMicroservice: new.(*v1beta1.Microservice),
				},
				Deleted: nil,
			}

			controller.queue.Add(event)
		},

		DeleteFunc: func(obj interface{}) {
			event := MicroserviceEvent{
				EventType: DELETED,
				Created:   nil,
				Updated:   nil,
				Deleted: &MicroserviceDeletedEvent{
					Microservice: obj.(*v1beta1.Microservice),
				},
			}

			controller.queue.Add(event)
		},
	})

	return controller
}
