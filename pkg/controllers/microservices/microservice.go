package controller

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	v1beta1 "github.com/Marcos30004347/seratos-api/pkg/apis/seratos/v1beta1"
)

// microserviceCreated is responsable for handling the creation of a microservice Resources
func microserviceCreated(c *Controller, obj *v1beta1.Microservice) error {
	microservice, err := c.microserviceLister.Microservices(obj.Namespace).Get(obj.Name)

	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("microservice '%s' in work queue no longer exists", obj.Namespace+"/"+obj.Name))
			return nil
		}
		return err
	}

	var sidecars []*v1beta1.Sidecar

	for _, sidecarname := range microservice.Spec.Sidecars {
		sidecar, err := c.sidecarsLister.Sidecars(microservice.Namespace).Get(sidecarname)

		if err != nil {
			if errors.IsNotFound(err) {
				utilruntime.HandleError(fmt.Errorf("sidecars '%s' in microservice no longer exists", obj.Namespace+"/"+obj.Name))
				return nil
			}
			return err
		}

		sidecars = append(sidecars, sidecar)
	}

	_, err = c.kubeRuntime.Clientset.AppsV1().Deployments(microservice.Namespace).Create(
		context.TODO(),
		createDeployment(microservice, sidecars),
		metav1.CreateOptions{},
	)

	if err != nil {
		utilruntime.HandleError(err)
		return err
	}

	c.recorder.Event(microservice, corev1.EventTypeNormal, "MicroserviceCreated", "Microservice Created")
	return nil
}

// microserviceUpdated is responsable for handling the update of a microservice resource
func microserviceUpdated(c *Controller, old *v1beta1.Microservice, new *v1beta1.Microservice) error {
	microservice, err := c.microserviceLister.Microservices(new.Namespace).Get(new.Name)

	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("microservice '%s' in work queue no longer exists", old.Namespace+"/"+old.Name))
			return nil
		}
		return err
	}

	var sidecars []*v1beta1.Sidecar

	for _, sidecarname := range microservice.Spec.Sidecars {
		sidecar, err := c.sidecarsLister.Sidecars(microservice.Namespace).Get(
			sidecarname,
		)

		if err != nil {
			if errors.IsNotFound(err) {
				utilruntime.HandleError(fmt.Errorf("sidecars '%s' in microservice no longer exists", old.Namespace+"/"+old.Name))
				return nil
			}
			return err
		}

		sidecars = append(sidecars, sidecar)
	}

	_, err = c.kubeRuntime.Clientset.AppsV1().Deployments(microservice.Namespace).Update(
		context.TODO(),
		createDeployment(new, sidecars),
		metav1.UpdateOptions{},
	)

	if err != nil {
		utilruntime.HandleError(err)
		return err
	}

	c.recorder.Event(microservice, corev1.EventTypeNormal, "MicroserviceUpdated", "Microservice Updated")
	return nil
}

// microserviceDeleted is responsably for handling the deletion of a Microservice Resource
func microserviceDeleted(c *Controller, obj *v1beta1.Microservice) error {
	microservice, err := c.microserviceLister.Microservices(obj.Namespace).Get(obj.Name)

	if err != nil {
		if errors.IsNotFound(err) {
			utilruntime.HandleError(fmt.Errorf("microservice '%s' in work queue no longer exists", obj.Namespace+"/"+obj.Name))
			return nil
		}
		return err
	}

	err = c.kubeRuntime.Clientset.AppsV1().Deployments(microservice.Namespace).Delete(
		context.TODO(),
		microservice.Name,
		metav1.DeleteOptions{},
	)

	if err != nil {
		utilruntime.HandleError(err)
		return err
	}

	c.recorder.Event(microservice, corev1.EventTypeNormal, "MicroserviceDeleted", "Microservice Deleted")
	return nil
}
