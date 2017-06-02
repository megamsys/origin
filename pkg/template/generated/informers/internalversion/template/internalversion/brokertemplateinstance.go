// This file was automatically generated by informer-gen

package internalversion

import (
	api "github.com/openshift/origin/pkg/template/api"
	internalinterfaces "github.com/openshift/origin/pkg/template/generated/informers/internalversion/internalinterfaces"
	internalclientset "github.com/openshift/origin/pkg/template/generated/internalclientset"
	internalversion "github.com/openshift/origin/pkg/template/generated/listers/template/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// BrokerTemplateInstanceInformer provides access to a shared informer and lister for
// BrokerTemplateInstances.
type BrokerTemplateInstanceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.BrokerTemplateInstanceLister
}

type brokerTemplateInstanceInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newBrokerTemplateInstanceInformer(client internalclientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.Template().BrokerTemplateInstances().List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.Template().BrokerTemplateInstances().Watch(options)
			},
		},
		&api.BrokerTemplateInstance{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *brokerTemplateInstanceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api.BrokerTemplateInstance{}, newBrokerTemplateInstanceInformer)
}

func (f *brokerTemplateInstanceInformer) Lister() internalversion.BrokerTemplateInstanceLister {
	return internalversion.NewBrokerTemplateInstanceLister(f.Informer().GetIndexer())
}
