package sidecars

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
	plugins.Register("Sidecars", func(config io.Reader) (admission.Interface, error) {
		return New()
	})
}

// The Plugin structure
type Plugin struct {
	*admission.Handler
	lister listers.SidecarLister
}

var _ = initializer.WantsSeratosInformerFactory(&Plugin{})

// Admit ensures that the object in-flight is of kind Foo.
// In addition checks that the bar are known.
func (d *Plugin) Admit(ctx context.Context, a admission.Attributes, oi admission.ObjectInterfaces) error {
	if a.GetKind().GroupKind() != seratos.Kind("Sidecar") {
		return nil
	}
	if !d.WaitForReady() {
		return admission.NewForbidden(a, fmt.Errorf("not yet ready to handle request"))
	}

	return nil
}

// SetSeratosInformerFactory gets Lister from SharedInformerFactory.
// The lister knows how to lists Bar.
func (d *Plugin) SetSeratosInformerFactory(f informers.SharedInformerFactory) {
	d.lister = f.Seratos().V1beta1().Sidecars().Lister()

	d.SetReadyFunc(f.Seratos().V1beta1().Sidecars().Informer().HasSynced)
}

// ValidateInitialization checks whether the plugin was correctly initialized.
func (d *Plugin) ValidateInitialization() error {
	if d.lister == nil {
		return fmt.Errorf("Sidecars Plugin missing policy lister")
	}
	return nil
}

// New creates a new  admission plugin
func New() (*Plugin, error) {
	return &Plugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}, nil
}
