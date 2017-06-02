// This file was automatically generated by lister-gen

package v1

import (
	api "github.com/openshift/origin/pkg/build/api"
	v1 "github.com/openshift/origin/pkg/build/api/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// BuildConfigLister helps list BuildConfigs.
type BuildConfigLister interface {
	// List lists all BuildConfigs in the indexer.
	List(selector labels.Selector) (ret []*v1.BuildConfig, err error)
	// BuildConfigs returns an object that can list and get BuildConfigs.
	BuildConfigs(namespace string) BuildConfigNamespaceLister
	BuildConfigListerExpansion
}

// buildConfigLister implements the BuildConfigLister interface.
type buildConfigLister struct {
	indexer cache.Indexer
}

// NewBuildConfigLister returns a new BuildConfigLister.
func NewBuildConfigLister(indexer cache.Indexer) BuildConfigLister {
	return &buildConfigLister{indexer: indexer}
}

// List lists all BuildConfigs in the indexer.
func (s *buildConfigLister) List(selector labels.Selector) (ret []*v1.BuildConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BuildConfig))
	})
	return ret, err
}

// BuildConfigs returns an object that can list and get BuildConfigs.
func (s *buildConfigLister) BuildConfigs(namespace string) BuildConfigNamespaceLister {
	return buildConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// BuildConfigNamespaceLister helps list and get BuildConfigs.
type BuildConfigNamespaceLister interface {
	// List lists all BuildConfigs in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.BuildConfig, err error)
	// Get retrieves the BuildConfig from the indexer for a given namespace and name.
	Get(name string) (*v1.BuildConfig, error)
	BuildConfigNamespaceListerExpansion
}

// buildConfigNamespaceLister implements the BuildConfigNamespaceLister
// interface.
type buildConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all BuildConfigs in the indexer for a given namespace.
func (s buildConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.BuildConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BuildConfig))
	})
	return ret, err
}

// Get retrieves the BuildConfig from the indexer for a given namespace and name.
func (s buildConfigNamespaceLister) Get(name string) (*v1.BuildConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(api.Resource("buildconfig"), name)
	}
	return obj.(*v1.BuildConfig), nil
}
