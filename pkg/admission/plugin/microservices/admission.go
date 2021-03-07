package microservices

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apiserver/pkg/admission"

	"github.com/Marcos30004347/seratos-api/pkg/admission/initializer"
	"github.com/Marcos30004347/seratos-api/pkg/apis/seratos"

	informers "github.com/Marcos30004347/seratos-api/pkg/generated/informers/externalversions"
	listers "github.com/Marcos30004347/seratos-api/pkg/generated/listers/seratos/v1beta1"
)

// Register registers a plugin
func Register(plugins *admission.Plugins) {
	plugins.Register("Microservices", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

// The Plugin structure
type Plugin struct {
	*admission.Handler
	fooLister           listers.FooLister
	microservicesLister listers.MicroserviceLister
}

var _ = initializer.WantsSeratosInformerFactory(&Plugin{})

func AdmitFoo(ctx context.Context, d *Plugin, a admission.Attributes, oi admission.ObjectInterfaces) error {
	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	return nil
}

func AdmitMicroservices(ctx context.Context, d *Plugin, a admission.Attributes, oi admission.ObjectInterfaces) error {
	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}
	return nil
}

// Admit ensures that the object in-flight is of kind Foo.
// In addition checks that the bar are known.
func (d *Plugin) Admit(ctx context.Context, a admission.Attributes, oi admission.ObjectInterfaces) error {
	switch a.GetKind().GroupKind() {
	case seratos.Kind("Foo"):
		return AdmitFoo(ctx, d, a, oi)
	case seratos.Kind("Microservices"):
		return AdmitMicroservices(ctx, d, a, oi)
	default:
		if !d.WaitForReady() {
			return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
		}
		return nil
	}
}

// SetSeratosInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Bar.
func (d *Plugin) SetSeratosInformerFactory(f informers.SharedInformerFactory) {
	d.fooLister = f.Seratos().V1beta1().Foos().Lister()
	d.microservicesLister = f.Seratos().V1beta1().Microservices().Lister()

	d.SetReadyFunc(f.Seratos().V1beta1().Microservices().Informer().HasSynced)
}

// ValidateInitialization checks whether the plugin was correctly initialized.
func (d *Plugin) ValidateInitialization() error {
	if d.fooLister == nil {
		return fmt.Errorf("missing policy lister")
	}
	return nil
}

// New creates a new ban foo topping admission plugin
func New() (*Plugin, error) {
	return &Plugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
