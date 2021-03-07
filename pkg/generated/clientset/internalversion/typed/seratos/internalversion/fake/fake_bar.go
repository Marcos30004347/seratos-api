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

package fake

import (
	"context"

	seratos "github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBars implements BarInterface
type FakeBars struct {
	Fake *FakeSeratos
}

var barsResource = schema.GroupVersionResource{Group: "seratos.microservices", Version: "", Resource: "bars"}

var barsKind = schema.GroupVersionKind{Group: "seratos.microservices", Version: "", Kind: "Bar"}

// Get takes name of the bar, and returns the corresponding bar object, and an error if there is any.
func (c *FakeBars) Get(ctx context.Context, name string, options v1.GetOptions) (result *seratos.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(barsResource, name), &seratos.Bar{})
	if obj == nil {
		return nil, err
	}
	return obj.(*seratos.Bar), err
}

// List takes label and field selectors, and returns the list of Bars that match those selectors.
func (c *FakeBars) List(ctx context.Context, opts v1.ListOptions) (result *seratos.BarList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(barsResource, barsKind, opts), &seratos.BarList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &seratos.BarList{ListMeta: obj.(*seratos.BarList).ListMeta}
	for _, item := range obj.(*seratos.BarList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bars.
func (c *FakeBars) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(barsResource, opts))
}

// Create takes the representation of a bar and creates it.  Returns the server's representation of the bar, and an error, if there is any.
func (c *FakeBars) Create(ctx context.Context, bar *seratos.Bar, opts v1.CreateOptions) (result *seratos.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(barsResource, bar), &seratos.Bar{})
	if obj == nil {
		return nil, err
	}
	return obj.(*seratos.Bar), err
}

// Update takes the representation of a bar and updates it. Returns the server's representation of the bar, and an error, if there is any.
func (c *FakeBars) Update(ctx context.Context, bar *seratos.Bar, opts v1.UpdateOptions) (result *seratos.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(barsResource, bar), &seratos.Bar{})
	if obj == nil {
		return nil, err
	}
	return obj.(*seratos.Bar), err
}

// Delete takes name of the bar and deletes it. Returns an error if one occurs.
func (c *FakeBars) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(barsResource, name), &seratos.Bar{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBars) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(barsResource, listOpts)

	_, err := c.Fake.Invokes(action, &seratos.BarList{})
	return err
}

// Patch applies the patch and returns the patched bar.
func (c *FakeBars) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *seratos.Bar, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(barsResource, name, pt, data, subresources...), &seratos.Bar{})
	if obj == nil {
		return nil, err
	}
	return obj.(*seratos.Bar), err
}
