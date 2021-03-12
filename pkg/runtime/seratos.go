package runtime

import (
	"time"

	"k8s.io/client-go/tools/clientcmd"

	clientset "github.com/Marcos30004347/seratos-api/pkg/generated/clientset/versioned"
	informers "github.com/Marcos30004347/seratos-api/pkg/generated/informers/externalversions"
	"k8s.io/client-go/rest"
)

type SeratosRuntime struct {
	ClientSet  clientset.Interface
	RESTConfig *rest.Config
	Informers  informers.SharedInformerFactory
}

func (r *SeratosRuntime) Run(stopCh <-chan struct{}) {
	r.Informers.Start(stopCh)
}
func (r *SeratosRuntime) Cleanup() {
}

func NewSeratosRuntime(
	kubeconfig string,
	master string,
) (*SeratosRuntime, error) {
	cfg, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)

	if err != nil {
		return nil, err
	}

	clientset, err := clientset.NewForConfig(cfg)

	if err != nil {
		return nil, err
	}

	return &SeratosRuntime{
		ClientSet:  clientset,
		RESTConfig: cfg,
		Informers:  informers.NewSharedInformerFactory(clientset, time.Second*30),
	}, nil
}
