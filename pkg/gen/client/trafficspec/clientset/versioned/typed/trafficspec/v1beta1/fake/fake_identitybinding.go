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
	v1beta1 "github.com/deislabs/smi-sdk-go/pkg/apis/trafficspec/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIdentityBindings implements IdentityBindingInterface
type FakeIdentityBindings struct {
	Fake *FakeSmispecV1beta1
	ns   string
}

var identitybindingsResource = schema.GroupVersionResource{Group: "smispec.io", Version: "v1beta1", Resource: "identitybindings"}

var identitybindingsKind = schema.GroupVersionKind{Group: "smispec.io", Version: "v1beta1", Kind: "IdentityBinding"}

// Get takes name of the identityBinding, and returns the corresponding identityBinding object, and an error if there is any.
func (c *FakeIdentityBindings) Get(name string, options v1.GetOptions) (result *v1beta1.IdentityBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(identitybindingsResource, c.ns, name), &v1beta1.IdentityBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityBinding), err
}

// List takes label and field selectors, and returns the list of IdentityBindings that match those selectors.
func (c *FakeIdentityBindings) List(opts v1.ListOptions) (result *v1beta1.IdentityBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(identitybindingsResource, identitybindingsKind, c.ns, opts), &v1beta1.IdentityBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.IdentityBindingList{ListMeta: obj.(*v1beta1.IdentityBindingList).ListMeta}
	for _, item := range obj.(*v1beta1.IdentityBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested identityBindings.
func (c *FakeIdentityBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(identitybindingsResource, c.ns, opts))

}

// Create takes the representation of a identityBinding and creates it.  Returns the server's representation of the identityBinding, and an error, if there is any.
func (c *FakeIdentityBindings) Create(identityBinding *v1beta1.IdentityBinding) (result *v1beta1.IdentityBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(identitybindingsResource, c.ns, identityBinding), &v1beta1.IdentityBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityBinding), err
}

// Update takes the representation of a identityBinding and updates it. Returns the server's representation of the identityBinding, and an error, if there is any.
func (c *FakeIdentityBindings) Update(identityBinding *v1beta1.IdentityBinding) (result *v1beta1.IdentityBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(identitybindingsResource, c.ns, identityBinding), &v1beta1.IdentityBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityBinding), err
}

// Delete takes name of the identityBinding and deletes it. Returns an error if one occurs.
func (c *FakeIdentityBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(identitybindingsResource, c.ns, name), &v1beta1.IdentityBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIdentityBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(identitybindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.IdentityBindingList{})
	return err
}

// Patch applies the patch and returns the patched identityBinding.
func (c *FakeIdentityBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.IdentityBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(identitybindingsResource, c.ns, name, pt, data, subresources...), &v1beta1.IdentityBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.IdentityBinding), err
}
