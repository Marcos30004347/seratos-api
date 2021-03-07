package foo

import (
	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	"github.com/Marcos30004347/seratos-api/pkg/registry"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/registry/generic"
	genericregistry "k8s.io/apiserver/pkg/registry/generic/registry"
	"k8s.io/apiserver/pkg/registry/rest"
)

// NewREST returns a RESTStorage object that will work against API services.
func NewREST(scheme *runtime.Scheme, optsGetter generic.RESTOptionsGetter) (*registry.REST, error) {
	strategy := NewStrategy(scheme)

	store := &genericregistry.Store{
		// Here is where you set that the bars objets are of kind Bar
		NewFunc:     func() runtime.Object { return &seratos.Foo{} },
		NewListFunc: func() runtime.Object { return &seratos.FooList{} },

		PredicateFunc:            MatchFoo,
		DefaultQualifiedResource: seratos.Resource("foos"),

		CreateStrategy: strategy,
		UpdateStrategy: strategy,
		DeleteStrategy: strategy,

		TableConvertor: rest.NewDefaultTableConvertor(seratos.Resource("foos")),
	}
	options := &generic.StoreOptions{RESTOptions: optsGetter, AttrFunc: GetAttrs}
	if err := store.CompleteWithOptions(options); err != nil {
		return nil, err
	}
	return &registry.REST{store}, nil
}
