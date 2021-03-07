/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	"time"

	seratos "github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	scheme "github.com/Marcos30004347/seratos-api/pkg/generated/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// BarsGetter has a method to return a BarInterface.
// A group's client should implement this interface.
type BarsGetter interface {
	Bars() BarInterface
}

// BarInterface has methods to work with Bar resources.
type BarInterface interface {
	Create(ctx context.Context, bar *seratos.Bar, opts v1.CreateOptions) (*seratos.Bar, error)
	Update(ctx context.Context, bar *seratos.Bar, opts v1.UpdateOptions) (*seratos.Bar, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*seratos.Bar, error)
	List(ctx context.Context, opts v1.ListOptions) (*seratos.BarList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *seratos.Bar, err error)
	BarExpansion
}

// bars implements BarInterface
type bars struct {
	client rest.Interface
}

// newBars returns a Bars
func newBars(c *SeratosClient) *bars {
	return &bars{
		client: c.RESTClient(),
	}
}

// Get takes name of the bar, and returns the corresponding bar object, and an error if there is any.
func (c *bars) Get(ctx context.Context, name string, options v1.GetOptions) (result *seratos.Bar, err error) {
	result = &seratos.Bar{}
	err = c.client.Get().
		Resource("bars").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Bars that match those selectors.
func (c *bars) List(ctx context.Context, opts v1.ListOptions) (result *seratos.BarList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &seratos.BarList{}
	err = c.client.Get().
		Resource("bars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested bars.
func (c *bars) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("bars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a bar and creates it.  Returns the server's representation of the bar, and an error, if there is any.
func (c *bars) Create(ctx context.Context, bar *seratos.Bar, opts v1.CreateOptions) (result *seratos.Bar, err error) {
	result = &seratos.Bar{}
	err = c.client.Post().
		Resource("bars").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bar).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a bar and updates it. Returns the server's representation of the bar, and an error, if there is any.
func (c *bars) Update(ctx context.Context, bar *seratos.Bar, opts v1.UpdateOptions) (result *seratos.Bar, err error) {
	result = &seratos.Bar{}
	err = c.client.Put().
		Resource("bars").
		Name(bar.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(bar).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the bar and deletes it. Returns an error if one occurs.
func (c *bars) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("bars").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *bars) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("bars").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched bar.
func (c *bars) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *seratos.Bar, err error) {
	result = &seratos.Bar{}
	err = c.client.Patch(pt).
		Resource("bars").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}