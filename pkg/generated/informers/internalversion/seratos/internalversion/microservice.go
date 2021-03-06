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

// Code generated by informer-gen. DO NOT EDIT.

package internalversion

import (
	"context"
	time "time"

	seratos "github.com/Marcos30004347/seratos-api/pkg/apis/seratos"
	clientsetinternalversion "github.com/Marcos30004347/seratos-api/pkg/generated/clientset/internalversion"
	internalinterfaces "github.com/Marcos30004347/seratos-api/pkg/generated/informers/internalversion/internalinterfaces"
	internalversion "github.com/Marcos30004347/seratos-api/pkg/generated/listers/seratos/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MicroserviceInformer provides access to a shared informer and lister for
// Microservices.
type MicroserviceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.MicroserviceLister
}

type microserviceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMicroserviceInformer constructs a new informer for Microservice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMicroserviceInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMicroserviceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMicroserviceInformer constructs a new informer for Microservice type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMicroserviceInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Seratos().Microservices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Seratos().Microservices(namespace).Watch(context.TODO(), options)
			},
		},
		&seratos.Microservice{},
		resyncPeriod,
		indexers,
	)
}

func (f *microserviceInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMicroserviceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *microserviceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&seratos.Microservice{}, f.defaultInformer)
}

func (f *microserviceInformer) Lister() internalversion.MicroserviceLister {
	return internalversion.NewMicroserviceLister(f.Informer().GetIndexer())
}
