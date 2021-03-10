package controller

import (
	"context"
	"fmt"

	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Marcos30004347/seratos-api/pkg/runtime"
)

// RunEvent creates a new controller
func ReadEvents(kube *runtime.KubeRuntime) {
	watcher, err := kube.Clientset.CoreV1().Events("").Watch(context.TODO(), metav1.ListOptions{})

	if err != nil {
		utilruntime.HandleError(fmt.Errorf("Cant listen events"))
	}

	ch := watcher.ResultChan()

	for event := range ch {
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("EVENT")
		klog.V(4).Infof("%#v", event)
	}
}
