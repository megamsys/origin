// This file was automatically generated by informer-gen

package v1

import (
	api_v1 "github.com/openshift/origin/pkg/deploy/api/v1"
	clientset "github.com/openshift/origin/pkg/deploy/generated/clientset"
	internalinterfaces "github.com/openshift/origin/pkg/deploy/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/origin/pkg/deploy/generated/listers/apps/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// DeploymentConfigInformer provides access to a shared informer and lister for
// DeploymentConfigs.
type DeploymentConfigInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.DeploymentConfigLister
}

type deploymentConfigInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

func newDeploymentConfigInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	sharedIndexInformer := cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.AppsV1().DeploymentConfigs(meta_v1.NamespaceAll).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.AppsV1().DeploymentConfigs(meta_v1.NamespaceAll).Watch(options)
			},
		},
		&api_v1.DeploymentConfig{},
		resyncPeriod,
		cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc},
	)

	return sharedIndexInformer
}

func (f *deploymentConfigInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&api_v1.DeploymentConfig{}, newDeploymentConfigInformer)
}

func (f *deploymentConfigInformer) Lister() v1.DeploymentConfigLister {
	return v1.NewDeploymentConfigLister(f.Informer().GetIndexer())
}
