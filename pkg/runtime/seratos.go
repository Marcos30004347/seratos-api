package runtime

import (
	"time"

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

func NewSeratosRuntime(cfg *rest.Config) (*SeratosRuntime, error) {
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
