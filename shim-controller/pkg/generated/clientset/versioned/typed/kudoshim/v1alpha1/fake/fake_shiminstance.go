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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/kudobuilder/shim/shim-controller/pkg/apis/kudoshim/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeShimInstances implements ShimInstanceInterface
type FakeShimInstances struct {
	Fake *FakeKudoshimV1alpha1
	ns   string
}

var shiminstancesResource = schema.GroupVersionResource{Group: "kudoshim.dev", Version: "v1alpha1", Resource: "shiminstances"}

var shiminstancesKind = schema.GroupVersionKind{Group: "kudoshim.dev", Version: "v1alpha1", Kind: "ShimInstance"}

// Get takes name of the shimInstance, and returns the corresponding shimInstance object, and an error if there is any.
func (c *FakeShimInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ShimInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(shiminstancesResource, c.ns, name), &v1alpha1.ShimInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShimInstance), err
}

// List takes label and field selectors, and returns the list of ShimInstances that match those selectors.
func (c *FakeShimInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ShimInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(shiminstancesResource, shiminstancesKind, c.ns, opts), &v1alpha1.ShimInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ShimInstanceList{ListMeta: obj.(*v1alpha1.ShimInstanceList).ListMeta}
	for _, item := range obj.(*v1alpha1.ShimInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested shimInstances.
func (c *FakeShimInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(shiminstancesResource, c.ns, opts))

}

// Create takes the representation of a shimInstance and creates it.  Returns the server's representation of the shimInstance, and an error, if there is any.
func (c *FakeShimInstances) Create(ctx context.Context, shimInstance *v1alpha1.ShimInstance, opts v1.CreateOptions) (result *v1alpha1.ShimInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(shiminstancesResource, c.ns, shimInstance), &v1alpha1.ShimInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShimInstance), err
}

// Update takes the representation of a shimInstance and updates it. Returns the server's representation of the shimInstance, and an error, if there is any.
func (c *FakeShimInstances) Update(ctx context.Context, shimInstance *v1alpha1.ShimInstance, opts v1.UpdateOptions) (result *v1alpha1.ShimInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(shiminstancesResource, c.ns, shimInstance), &v1alpha1.ShimInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShimInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeShimInstances) UpdateStatus(ctx context.Context, shimInstance *v1alpha1.ShimInstance, opts v1.UpdateOptions) (*v1alpha1.ShimInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(shiminstancesResource, "status", c.ns, shimInstance), &v1alpha1.ShimInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShimInstance), err
}

// Delete takes name of the shimInstance and deletes it. Returns an error if one occurs.
func (c *FakeShimInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(shiminstancesResource, c.ns, name), &v1alpha1.ShimInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeShimInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(shiminstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ShimInstanceList{})
	return err
}

// Patch applies the patch and returns the patched shimInstance.
func (c *FakeShimInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ShimInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(shiminstancesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ShimInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ShimInstance), err
}
