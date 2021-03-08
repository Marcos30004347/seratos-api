package event

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

// NewStrategy creates and returns a EventStrategy instance
func NewStrategy(typer runtime.ObjectTyper) EventStrategy {
	return EventStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Event
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, error) {
	apiserver, ok := obj.(*seratos.Event)
	if !ok {
		return nil, nil, fmt.Errorf("given object is not a Event")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), nil
}

// MatchEvent is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchEvent(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *seratos.Event) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type EventStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (EventStrategy) NamespaceScoped() bool {
	return true
}

func (EventStrategy) PrepareForCreate(ctx context.Context, obj runtime.Object) {
}

func (EventStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
}

// Here is where we actually use the Validate Function defined in the api
func (EventStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}

}

func (EventStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (EventStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (EventStrategy) Canonicalize(obj runtime.Object) {
}

func (EventStrategy) ValidateUpdate(ctx context.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}
