package microservices

import (
	"context"
	"fmt"
	"io"

	"k8s.io/apimachinery/pkg/api/errors"
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
	microserviceLister listers.MicroserviceLister
	sidecarLister      listers.SidecarLister
}

var _ = initializer.WantsSeratosInformerFactory(&Plugin{})

// Admit ensures that the object in-flight is of kind Microservice.
// In addition checks that the sidecar are known.
func (d *Plugin) Admit(ctx context.Context, a admission.Attributes, oi admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != seratos.Kind("Microservice") {
		return nil
	}

	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	obj := a.GetObject()

	microservice := obj.(*seratos.Microservice)

	for _, sidecar := range microservice.Spec.Sidecars {
		if _, err := d.sidecarLister.Sidecars(a.GetNamespace()).Get(sidecar); err != nil && errors.IsNotFound(err) {
			return errors.NewForbidden(
				a.GetResource().GroupResource(),
				a.GetName(),
				fmt.Errorf("Unknown Sidecar: %s, you may have not deployed, or it may be in another namespace", sidecar),
			)
		}
	}

	return nil
}

// SetSeratosInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Bar.
func (d *Plugin) SetSeratosInformerFactory(f informers.SharedInformerFactory) {
	d.microserviceLister = f.Seratos().V1beta1().Microservices().Lister()
	d.sidecarLister = f.Seratos().V1beta1().Sidecars().Lister()

	d.SetReadyFunc(f.Seratos().V1beta1().Sidecars().Informer().HasSynced)
}

// ValidateInitialization checks whether the plugin was correctly initialized.
func (d *Plugin) ValidateInitialization() error {
	if d.microserviceLister == nil {
		return fmt.Errorf("Microservices Plugin missing microservices policy lister")
	}
	if d.sidecarLister == nil {
		return fmt.Errorf("Microservices Plugin missing sidecars policy lister")
	}
	return nil
}

// New creates a new  admission plugin
func New() (*Plugin, error) {
	return &Plugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
