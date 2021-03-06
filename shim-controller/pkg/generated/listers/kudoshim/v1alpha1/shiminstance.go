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
// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/kudobuilder/shim/shim-controller/pkg/apis/kudoshim/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ShimInstanceLister helps list ShimInstances.
type ShimInstanceLister interface {
	// List lists all ShimInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ShimInstance, err error)
	// ShimInstances returns an object that can list and get ShimInstances.
	ShimInstances(namespace string) ShimInstanceNamespaceLister
	ShimInstanceListerExpansion
}

// shimInstanceLister implements the ShimInstanceLister interface.
type shimInstanceLister struct {
	indexer cache.Indexer
}

// NewShimInstanceLister returns a new ShimInstanceLister.
func NewShimInstanceLister(indexer cache.Indexer) ShimInstanceLister {
	return &shimInstanceLister{indexer: indexer}
}

// List lists all ShimInstances in the indexer.
func (s *shimInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ShimInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShimInstance))
	})
	return ret, err
}

// ShimInstances returns an object that can list and get ShimInstances.
func (s *shimInstanceLister) ShimInstances(namespace string) ShimInstanceNamespaceLister {
	return shimInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShimInstanceNamespaceLister helps list and get ShimInstances.
type ShimInstanceNamespaceLister interface {
	// List lists all ShimInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ShimInstance, err error)
	// Get retrieves the ShimInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ShimInstance, error)
	ShimInstanceNamespaceListerExpansion
}

// shimInstanceNamespaceLister implements the ShimInstanceNamespaceLister
// interface.
type shimInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ShimInstances in the indexer for a given namespace.
func (s shimInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ShimInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ShimInstance))
	})
	return ret, err
}

// Get retrieves the ShimInstance from the indexer for a given namespace and name.
func (s shimInstanceNamespaceLister) Get(name string) (*v1alpha1.ShimInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shiminstance"), name)
	}
	return obj.(*v1alpha1.ShimInstance), nil
}
