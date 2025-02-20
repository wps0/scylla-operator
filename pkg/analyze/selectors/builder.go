package selectors

import (
	"github.com/scylladb/scylla-operator/pkg/analyze/sources"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"reflect"
)

type builder struct {
	resources   map[string]reflect.Type
	constraints map[string][]*constraint
	relations   []*relation
}

func Type[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

func Select(label string, typ reflect.Type) *builder {
	return (&builder{
		resources:   make(map[string]reflect.Type),
		constraints: make(map[string][]*constraint),
		relations:   make([]*relation, 0),
	}).Select(label, typ)
}

func (b *builder) Select(label string, typ reflect.Type) *builder {
	if _, exists := b.resources[label]; exists {
		panic("TODO: Handle duplicate labels")
	}

	b.resources[label] = typ

	return b
}

func (b *builder) Filter(label string, f any) *builder {
	typ, defined := b.resources[label]
	if !defined {
		panic("TODO: Handle undefined labels in Filter")
	}

	constraint := newConstraint(label, f)
	if constraint.Labels()[label] != reflect.PointerTo(typ) {
		panic("TODO: Handle mismatched type in Filter")
	}

	b.constraints[label] = append(b.constraints[label], constraint)

	return b
}

func (b *builder) Relate(lhs, rhs string, f any) *builder {
	// TODO: Check input

	relation := newRelation(lhs, rhs, f)

	b.relations = append(b.relations, relation)

	return b
}

func eraseSliceType[T any](slice []T) []any {
	result := make([]any, len(slice))

	for i, _ := range slice {
		result[i] = slice[i]
	}

	return result
}

func fromDataSource(ds *sources.DataSource) map[reflect.Type][]any {
	clusters, _ := ds.ScyllaClusterLister.List(labels.Everything())
	pods, _ := ds.PodLister.List(labels.Everything())

	return map[reflect.Type][]any{
		reflect.TypeFor[scyllav1.ScyllaCluster](): eraseSliceType(clusters),
		reflect.TypeFor[v1.Pod]():                 eraseSliceType(pods),
	}
}

func (b *builder) Collect(labels []string, function any) func(*sources.DataSource) {
	for _, label := range labels {
		if _, contains := b.resources[label]; !contains {
			panic("TODO: Handle undefined label")
		}
	}

	callback := newFunction[bool](labels, function)
	executor := newExecutor(b.resources, b.constraints, b.relations)

	return func(ds *sources.DataSource) {
		executor.execute(fromDataSource(ds), callback)
	}
}
