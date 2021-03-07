package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	informers "github.com/Marcos30004347/seratos-api/pkg/generated/informers/externalversions"
)

// WantsSeratosInformerFactory defines a function which sets InformerFactory for admission plugins that need it
type WantsSeratosInformerFactory interface {
	SetSeratosInformerFactory(informers.SharedInformerFactory)
	admission.InitializationValidator
}
