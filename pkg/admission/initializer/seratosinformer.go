package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/Marcos30004347/seratos-api/pkg/generated/informers/externalversions"
)

type seratosInformerPluginInitializer struct {
	informers informers.SharedInformerFactory
}

var _ admission.PluginInitializer = seratosInformerPluginInitializer{}

// New creates an instance of custom admission plugins initializer.
func New(informers informers.SharedInformerFactory) seratosInformerPluginInitializer {
	return seratosInformerPluginInitializer{
		informers: informers,
	}
}

// Initialize checks the initialization interfaces implemented by a plugin
// and provide the appropriate initialization data
func (i seratosInformerPluginInitializer) Initialize(plugin admission.Interface) {
	if wants, ok := plugin.(WantsSeratosInformerFactory); ok {
		wants.SetSeratosInformerFactory(i.informers)
	}
}
