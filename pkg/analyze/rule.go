package analyze

import (
	"errors"
	"fmt"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"reflect"
	"strings"
)

type InterfaceIter = func() interface{}

type Condition interface {
	Evaluate(ds *DataSource) InterfaceIter
	EvaluateOnResource(res interface{}) bool
	Kind() string
}

type ResourceCondition struct {
	ResKind   string
	Condition fields.Selector
}

func (r *ResourceCondition) requirementMatches(obj interface{}, requirement fields.Requirement) bool {
	it := GetFieldValueIterator(requirement.Field, obj)
	val := it()
	for val != nil && val.String() != requirement.Value {
		val = it()
	}
	return val != nil
}

func (r *ResourceCondition) Evaluate(ds *DataSource) InterfaceIter {
	return func() interface{} { return nil }
}

func (r *ResourceCondition) EvaluateOnResource(res interface{}) bool {
	reqs := r.Condition.Requirements()
	for _, req := range reqs {
		if !r.requirementMatches(res, req) {
			return false
		}
	}
	return true
}

func (r *ResourceCondition) Kind() string {
	return r.ResKind
}

type FunctionalCondition[T any] struct {
	ResKind   string
	Condition func(T) bool
}

func (r *FunctionalCondition[T]) Evaluate(ds *DataSource) InterfaceIter {
	return func() interface{} { return nil }
}

func (r *FunctionalCondition[T]) EvaluateOnResource(res interface{}) bool {
	return r.Condition(res.(T))
}

func (r *FunctionalCondition[T]) Kind() string {
	//return reflect.TypeFor[T]().Name()
	return r.ResKind
}

type Relation interface {
	EvaluateOn(a interface{}, b interface{}) (bool, error)
}

type resourceConnection struct {
	Lhs int
	Rhs int
	Rel Relation
	// True if this relation exists.
	Exists bool
}

type EqualFieldsRelation struct {
	LhsPath string
	RhsPath string
}

func (r *EqualFieldsRelation) EvaluateOn(a interface{}, b interface{}) (bool, error) {
	lhsIter := GetFieldValueIterator(r.LhsPath, a)
	rhsIter := GetFieldValueIterator(r.RhsPath, b)
	lhsAll := make([]*reflect.Value, 0)
	for lhs := lhsIter(); lhs != nil; lhs = lhsIter() {
		lhsAll = append(lhsAll, lhs)
	}
	for rhs := rhsIter(); rhs != nil; rhs = rhsIter() {
		for _, lhs := range lhsAll {
			if lhs.String() == rhs.String() {
				return true, nil
			}
		}
	}
	return false, nil
}

type FunctionalEqualStringFieldsRelation[L, R any] struct {
	LhsExtractor func(lhs L) string
	RhsExtractor func(lhs R) string
}

func (f *FunctionalEqualStringFieldsRelation[L, R]) EvaluateOn(a interface{}, b interface{}) (bool, error) {
	return f.LhsExtractor(a.(L)) == f.RhsExtractor(b.(R)), nil
}

func getFieldValueIterator(node reflect.Value, fields []string) func() *reflect.Value {
	switch node.Kind() {
	case reflect.Ptr:
		return getFieldValueIterator(node.Elem(), fields)
	}

	if len(fields) == 0 {
		called := false
		return func() *reflect.Value {
			if !called {
				called = true
				return &node
			}
			return nil
		}
	}

	switch node.Kind() {
	case reflect.Map:
		return getFieldValueIterator(node.MapIndex(reflect.ValueOf(fields[0])), fields[1:])
	case reflect.Struct:
		return getFieldValueIterator(node.FieldByName(fields[0]), fields[1:])
	case reflect.Slice | reflect.Array:
		i := -1
		iter := func() *reflect.Value { return nil }
		return func() *reflect.Value {
			val := iter()
			for val == nil && i+1 < node.Len() {
				i++
				iter = getFieldValueIterator(node.Index(i), fields)
				val = iter()
			}
			return val
		}
	case reflect.Invalid:
		return func() *reflect.Value { return nil }
	default:
		panic(errors.New(fmt.Sprintf("unknown field type %s for %v", node, node)))
	}
	return nil
}

func GetFieldValueIterator(path string, obj interface{}) func() *reflect.Value {
	path = strings.Map(func(ch rune) rune {
		if ch == ' ' {
			return -1
		}
		return ch
	}, path)
	fieldNames := strings.Split(path, ".")
	if fieldNames[0] == "Metadata" {
		fieldNames = fieldNames[1:]
	}
	return getFieldValueIterator(reflect.ValueOf(obj), fieldNames)
}

type Rule struct {
	Diagnosis   string
	Suggestions string
	Resources   []Condition
	Relations   []resourceConnection
}

var CsiDriverMissing = Rule{
	Diagnosis:   "local-csi-driver CSIDriver, referenced by <NAME> StorageClass, is missing",
	Suggestions: "deploy local-csi-driver provisioner",
	Resources: []Condition{
		&FunctionalCondition[*scyllav1.ScyllaCluster]{
			ResKind: "ScyllaCluster",
			//fields.AndSelectors(
			//	fields.ParseSelectorOrDie("Status.Conditions.Type=StatefulSetControllerProgressing"),
			//	fields.ParseSelectorOrDie("Status.Conditions.Type=Progressing"),
			//	fields.ParseSelectorOrDie("Spec.Datacenter.Racks.Storage.StorageClassName=scylladb-local-xfs"))
			Condition: func(c *scyllav1.ScyllaCluster) bool {
				storageClassXfs := false
				conditionControllerProgressing := false
				conditionProgressing := false
				for _, rack := range c.Spec.Datacenter.Racks {
					if *rack.Storage.StorageClassName == "scylladb-local-xfs" {
						storageClassXfs = true
					}
				}
				for _, cond := range c.Status.Conditions {
					if cond.Type == "StatefulSetControllerProgressing" {
						conditionControllerProgressing = true
					} else if cond.Type == "Progressing" {
						conditionProgressing = true
					}
				}
				return storageClassXfs && conditionProgressing && conditionControllerProgressing
			},
		},
		&ResourceCondition{
			ResKind:   "Pod",
			Condition: fields.Everything(),
		},
		&ResourceCondition{
			ResKind:   "StorageClass",
			Condition: fields.Everything(),
		},
		&ResourceCondition{
			ResKind:   "CSIDriver",
			Condition: fields.Everything(),
		},
	},
	Relations: []resourceConnection{
		{
			Lhs:    0,
			Rhs:    1,
			Exists: true,
			Rel: &FunctionalEqualStringFieldsRelation[*scyllav1.ScyllaCluster, *v1.Pod]{
				LhsExtractor: func(lhs *scyllav1.ScyllaCluster) string { return lhs.Name },     // LhsPath: "Metadata.Name",
				RhsExtractor: func(rhs *v1.Pod) string { return rhs.Labels["scylla/cluster"] }, // RhsPath: "Metadata.Labels.scylla/cluster",
			},
		},
		{
			Lhs:    0,
			Rhs:    2,
			Exists: true,
			Rel: &EqualFieldsRelation{
				LhsPath: "Spec.Datacenter.Racks.Storage.StorageClassName",
				RhsPath: "Metadata.Name",
			},
		},
		{
			Lhs:    2,
			Rhs:    3,
			Exists: false,
			Rel: &EqualFieldsRelation{
				LhsPath: "Provisioner",
				RhsPath: "Metadata.Name",
			},
		},
	},
}
