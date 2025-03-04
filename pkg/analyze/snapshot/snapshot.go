package snapshot

import (
	"context"
	"fmt"
	scyllaversioned "github.com/scylladb/scylla-operator/pkg/client/scylla/clientset/versioned"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/pager"
	"reflect"
)

// Może zamienić interface na runtime.Object
type Snapshot interface {
	Add(obj interface{})
	List(objType reflect.Type) []interface{}
	All() map[reflect.Type][]interface{}
}

type defaultSnapshot struct {
	objects map[reflect.Type][]interface{}
}

func NewEmptySnapshot() defaultSnapshot {
	ds := defaultSnapshot{
		objects: make(map[reflect.Type][]interface{}),
	}
	return ds
}

func NewSnapshot(objs map[reflect.Type][]interface{}) defaultSnapshot {
	ds := defaultSnapshot{
		objects: objs,
	}
	return ds
}

func (ds *defaultSnapshot) Add(obj interface{}) {
	t := reflect.TypeOf(obj)
	if _, exists := ds.objects[t]; !exists {
		ds.objects[t] = make([]interface{}, 0)
	}
	ds.objects[t] = append(ds.objects[t], obj)
}

func (ds *defaultSnapshot) List(objType reflect.Type) []interface{} {
	list, exists := ds.objects[objType]
	if !exists {
		return make([]interface{}, 0)
	}
	return list
}

func (ds *defaultSnapshot) All() map[reflect.Type][]interface{} {
	return ds.objects
}

func BuildListerWithOptions[T any](
	ctx context.Context,
	factory func(cache.Indexer) T,
	listFunc func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error),
	options metav1.ListOptions,
) (T, error) {
	p := pager.New(pager.SimplePageFunc(func(opts metav1.ListOptions) (runtime.Object, error) {
		return listFunc(ctx, opts)
	}))

	// Prevent users from providing unwanted ones or tempering options that pager controls
	options = metav1.ListOptions{
		LabelSelector: options.LabelSelector,
		FieldSelector: options.FieldSelector,
	}

	indexer := cache.NewIndexer(cache.MetaNamespaceKeyFunc, cache.Indexers{"NamespaceIndex": cache.MetaNamespaceIndexFunc})
	err := p.EachListItemWithAlloc(ctx, options, func(obj runtime.Object) error {
		err := indexer.Add(obj)
		if err != nil {
			return fmt.Errorf("can't add object to indexer %v: %w", obj, err)
		}
		return nil
	})
	if err != nil {
		return *new(T), fmt.Errorf("can't iterate over list items: %w", err)
	}

	return factory(indexer), nil
}

func BuildLister[T any](ctx context.Context, factory func(cache.Indexer) T, listFunc func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error)) (T, error) {
	return BuildListerWithOptions[T](ctx, factory, listFunc, metav1.ListOptions{})
}

func BuildList(
	ctx context.Context,
	ds Snapshot,
	listFunc func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error),
) error {
	p := pager.New(pager.SimplePageFunc(func(opts metav1.ListOptions) (runtime.Object, error) {
		return listFunc(ctx, opts)
	}))

	options := metav1.ListOptions{}

	// Prevent users from providing unwanted ones or tempering options that pager controls
	options = metav1.ListOptions{
		LabelSelector: options.LabelSelector,
		FieldSelector: options.FieldSelector,
	}

	err := p.EachListItemWithAlloc(ctx, options, func(obj runtime.Object) error {
		ds.Add(obj)
		return nil
	})
	if err != nil {
		return fmt.Errorf("can't iterate over list items: %w", err)
	}

	return nil
}

func NewSnapshotFromListers(
	ctx context.Context,
	kubeClient kubernetes.Interface,
	scyllaClient scyllaversioned.Interface,
) (Snapshot, error) {
	ds := NewEmptySnapshot()

	err := BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return kubeClient.CoreV1().Pods(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build pod lister: %w", err)
	}

	err = BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return kubeClient.CoreV1().Services(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build service lister: %w", err)
	}

	err = BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return kubeClient.CoreV1().Secrets(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build secret lister: %w", err)
	}

	err = BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return kubeClient.CoreV1().ConfigMaps(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build config map lister: %w", err)
	}

	err = BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return kubeClient.CoreV1().ServiceAccounts(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build service account lister: %w", err)
	}

	err = BuildList(ctx, &ds, func(ctx context.Context, options metav1.ListOptions) (runtime.Object, error) {
		return scyllaClient.ScyllaV1().ScyllaClusters(corev1.NamespaceAll).List(ctx, options)
	})
	if err != nil {
		return nil, fmt.Errorf("can't build scylla cluster lister: %w", err)
	}

	return &ds, nil
}
