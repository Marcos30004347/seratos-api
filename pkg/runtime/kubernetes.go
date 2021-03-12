package runtime

import (
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/component-base/logs"

	kubeinformers "k8s.io/client-go/informers"
)

type KubeRuntime struct {
	Clientset  kubernetes.Interface
	RESTConfig *rest.Config
	Informers  kubeinformers.SharedInformerFactory
}

func (c *KubeRuntime) Run(stopCh <-chan struct{}) {
	c.Informers.Start(stopCh)
}

func (c *KubeRuntime) Cleanup() {
	logs.FlushLogs()
}

func NewKubernetesRuntime(
	kubeconfig string,
	master string,
) (*KubeRuntime, error) {

	logs.InitLogs()
	defer logs.FlushLogs()

	cfg, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)

	if err != nil {
		return nil, err
	}

	clientset, err := kubernetes.NewForConfig(cfg)

	if err != nil {
		return nil, err
	}

	informer := kubeinformers.NewSharedInformerFactory(clientset, time.Second*30)

	return &KubeRuntime{
		Clientset:  clientset,
		RESTConfig: cfg,
		Informers:  informer,
	}, nil
}
