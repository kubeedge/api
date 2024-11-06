/*
Copyright The KubeEdge Authors.

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

	v1 "github.com/kubeedge/api/apis/rules/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRules implements RuleInterface
type FakeRules struct {
	Fake *FakeRulesV1
	ns   string
}

var rulesResource = v1.SchemeGroupVersion.WithResource("rules")

var rulesKind = v1.SchemeGroupVersion.WithKind("Rule")

// Get takes name of the rule, and returns the corresponding rule object, and an error if there is any.
func (c *FakeRules) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rulesResource, c.ns, name), &v1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Rule), err
}

// List takes label and field selectors, and returns the list of Rules that match those selectors.
func (c *FakeRules) List(ctx context.Context, opts metav1.ListOptions) (result *v1.RuleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rulesResource, rulesKind, c.ns, opts), &v1.RuleList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.RuleList{ListMeta: obj.(*v1.RuleList).ListMeta}
	for _, item := range obj.(*v1.RuleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rules.
func (c *FakeRules) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rulesResource, c.ns, opts))

}

// Create takes the representation of a rule and creates it.  Returns the server's representation of the rule, and an error, if there is any.
func (c *FakeRules) Create(ctx context.Context, rule *v1.Rule, opts metav1.CreateOptions) (result *v1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rulesResource, c.ns, rule), &v1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Rule), err
}

// Update takes the representation of a rule and updates it. Returns the server's representation of the rule, and an error, if there is any.
func (c *FakeRules) Update(ctx context.Context, rule *v1.Rule, opts metav1.UpdateOptions) (result *v1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rulesResource, c.ns, rule), &v1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Rule), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeRules) UpdateStatus(ctx context.Context, rule *v1.Rule, opts metav1.UpdateOptions) (*v1.Rule, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(rulesResource, "status", c.ns, rule), &v1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Rule), err
}

// Delete takes name of the rule and deletes it. Returns an error if one occurs.
func (c *FakeRules) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rulesResource, c.ns, name, opts), &v1.Rule{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRules) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rulesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.RuleList{})
	return err
}

// Patch applies the patch and returns the patched rule.
func (c *FakeRules) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Rule, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rulesResource, c.ns, name, pt, data, subresources...), &v1.Rule{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Rule), err
}
