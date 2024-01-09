// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/scylladb/scylla-operator/pkg/externalapi/monitoring/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePrometheuses implements PrometheusInterface
type FakePrometheuses struct {
	Fake *FakeMonitoringV1
	ns   string
}

var prometheusesResource = v1.SchemeGroupVersion.WithResource("prometheuses")

var prometheusesKind = v1.SchemeGroupVersion.WithKind("Prometheus")

// Get takes name of the prometheus, and returns the corresponding prometheus object, and an error if there is any.
func (c *FakePrometheuses) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(prometheusesResource, c.ns, name), &v1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Prometheus), err
}

// List takes label and field selectors, and returns the list of Prometheuses that match those selectors.
func (c *FakePrometheuses) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PrometheusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(prometheusesResource, prometheusesKind, c.ns, opts), &v1.PrometheusList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.PrometheusList{ListMeta: obj.(*v1.PrometheusList).ListMeta}
	for _, item := range obj.(*v1.PrometheusList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested prometheuses.
func (c *FakePrometheuses) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(prometheusesResource, c.ns, opts))

}

// Create takes the representation of a prometheus and creates it.  Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Create(ctx context.Context, prometheus *v1.Prometheus, opts metav1.CreateOptions) (result *v1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(prometheusesResource, c.ns, prometheus), &v1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Prometheus), err
}

// Update takes the representation of a prometheus and updates it. Returns the server's representation of the prometheus, and an error, if there is any.
func (c *FakePrometheuses) Update(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (result *v1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(prometheusesResource, c.ns, prometheus), &v1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Prometheus), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePrometheuses) UpdateStatus(ctx context.Context, prometheus *v1.Prometheus, opts metav1.UpdateOptions) (*v1.Prometheus, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(prometheusesResource, "status", c.ns, prometheus), &v1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Prometheus), err
}

// Delete takes name of the prometheus and deletes it. Returns an error if one occurs.
func (c *FakePrometheuses) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(prometheusesResource, c.ns, name, opts), &v1.Prometheus{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePrometheuses) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(prometheusesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1.PrometheusList{})
	return err
}

// Patch applies the patch and returns the patched prometheus.
func (c *FakePrometheuses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Prometheus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(prometheusesResource, c.ns, name, pt, data, subresources...), &v1.Prometheus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Prometheus), err
}
