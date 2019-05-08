/*
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

package v1beta1

import (
	time "time"

	trafficspecv1beta1 "github.com/deislabs/smi-sdk-go/pkg/apis/trafficspec/v1beta1"
	versioned "github.com/deislabs/smi-sdk-go/pkg/gen/client/trafficspec/clientset/versioned"
	internalinterfaces "github.com/deislabs/smi-sdk-go/pkg/gen/client/trafficspec/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/deislabs/smi-sdk-go/pkg/gen/client/trafficspec/listers/trafficspec/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// IdentityBindingInformer provides access to a shared informer and lister for
// IdentityBindings.
type IdentityBindingInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.IdentityBindingLister
}

type identityBindingInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIdentityBindingInformer constructs a new informer for IdentityBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIdentityBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIdentityBindingInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIdentityBindingInformer constructs a new informer for IdentityBinding type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIdentityBindingInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SmispecV1beta1().IdentityBindings(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SmispecV1beta1().IdentityBindings(namespace).Watch(options)
			},
		},
		&trafficspecv1beta1.IdentityBinding{},
		resyncPeriod,
		indexers,
	)
}

func (f *identityBindingInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIdentityBindingInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *identityBindingInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&trafficspecv1beta1.IdentityBinding{}, f.defaultInformer)
}

func (f *identityBindingInformer) Lister() v1beta1.IdentityBindingLister {
	return v1beta1.NewIdentityBindingLister(f.Informer().GetIndexer())
}
