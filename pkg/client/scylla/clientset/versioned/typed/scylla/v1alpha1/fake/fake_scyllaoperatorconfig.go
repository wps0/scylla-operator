// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeScyllaOperatorConfigs implements ScyllaOperatorConfigInterface
type FakeScyllaOperatorConfigs struct {
	Fake *FakeScyllaV1alpha1
}

var scyllaoperatorconfigsResource = v1alpha1.SchemeGroupVersion.WithResource("scyllaoperatorconfigs")

var scyllaoperatorconfigsKind = v1alpha1.SchemeGroupVersion.WithKind("ScyllaOperatorConfig")

// Get takes name of the scyllaOperatorConfig, and returns the corresponding scyllaOperatorConfig object, and an error if there is any.
func (c *FakeScyllaOperatorConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ScyllaOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(scyllaoperatorconfigsResource, name), &v1alpha1.ScyllaOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), err
}

// List takes label and field selectors, and returns the list of ScyllaOperatorConfigs that match those selectors.
func (c *FakeScyllaOperatorConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ScyllaOperatorConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(scyllaoperatorconfigsResource, scyllaoperatorconfigsKind, opts), &v1alpha1.ScyllaOperatorConfigList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ScyllaOperatorConfigList{ListMeta: obj.(*v1alpha1.ScyllaOperatorConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ScyllaOperatorConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested scyllaOperatorConfigs.
func (c *FakeScyllaOperatorConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(scyllaoperatorconfigsResource, opts))
}

// Create takes the representation of a scyllaOperatorConfig and creates it.  Returns the server's representation of the scyllaOperatorConfig, and an error, if there is any.
func (c *FakeScyllaOperatorConfigs) Create(ctx context.Context, scyllaOperatorConfig *v1alpha1.ScyllaOperatorConfig, opts v1.CreateOptions) (result *v1alpha1.ScyllaOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(scyllaoperatorconfigsResource, scyllaOperatorConfig), &v1alpha1.ScyllaOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), err
}

// Update takes the representation of a scyllaOperatorConfig and updates it. Returns the server's representation of the scyllaOperatorConfig, and an error, if there is any.
func (c *FakeScyllaOperatorConfigs) Update(ctx context.Context, scyllaOperatorConfig *v1alpha1.ScyllaOperatorConfig, opts v1.UpdateOptions) (result *v1alpha1.ScyllaOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(scyllaoperatorconfigsResource, scyllaOperatorConfig), &v1alpha1.ScyllaOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeScyllaOperatorConfigs) UpdateStatus(ctx context.Context, scyllaOperatorConfig *v1alpha1.ScyllaOperatorConfig, opts v1.UpdateOptions) (*v1alpha1.ScyllaOperatorConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(scyllaoperatorconfigsResource, "status", scyllaOperatorConfig), &v1alpha1.ScyllaOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), err
}

// Delete takes name of the scyllaOperatorConfig and deletes it. Returns an error if one occurs.
func (c *FakeScyllaOperatorConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(scyllaoperatorconfigsResource, name, opts), &v1alpha1.ScyllaOperatorConfig{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeScyllaOperatorConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(scyllaoperatorconfigsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ScyllaOperatorConfigList{})
	return err
}

// Patch applies the patch and returns the patched scyllaOperatorConfig.
func (c *FakeScyllaOperatorConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ScyllaOperatorConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(scyllaoperatorconfigsResource, name, pt, data, subresources...), &v1alpha1.ScyllaOperatorConfig{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ScyllaOperatorConfig), err
}
