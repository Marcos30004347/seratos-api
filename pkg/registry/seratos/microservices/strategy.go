package microservices

import (
	"context"
	"fmt"

	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"
)

// NewStrategy creates and returns a MicroservicesStrategy instance
func NewStrategy(typer runtime.ObjectTyper) MicroservicesStrategy {
	return MicroservicesStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Bar
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*seratos.Microservice)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Bar")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchMicroservices is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchMicroservices(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *seratos.Microservice) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type MicroservicesStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (MicroservicesStrategy) NamespaceScoped() bool {
	return true
}

func (MicroservicesStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (MicroservicesStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

// Here is where we actually use the Validate Function defined in the api
func (MicroservicesStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}

}

func (MicroservicesStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (MicroservicesStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (MicroservicesStrategy) Canonicalize(obj runtime.Object) {
}

func (MicroservicesStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
