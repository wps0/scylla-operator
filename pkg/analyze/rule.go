package analyze

import (
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

type InterfaceIter = func() interface{}

type Condition interface {
	Evaluate(ds *DataSource) InterfaceIter
	EvaluateOnResource(res interface{}) bool
	Kind() string
}

type ResourceCondition[T any] struct {
	ResKind   string
	Condition func(T) bool
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

type FunctionalEqualStringFieldsRelation[L, R any] struct {
	LhsExtractor func(lhs L) []string
	RhsExtractor func(lhs R) []string
}

func (f *FunctionalEqualStringFieldsRelation[L, R]) EvaluateOn(a interface{}, b interface{}) (bool, error) {
	if a != nil {
		for _, lhs := range f.LhsExtractor(a.(L)) {
			if b != nil {
				for _, rhs := range f.RhsExtractor(b.(R)) {
					if lhs == rhs {
						return true, nil
					}
				}
			}
		}
	}
	return false, nil
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
		&FunctionalCondition[*v1.Pod]{
			ResKind:   "Pod",
			Condition: func(c *v1.Pod) bool { return true },
		},
		&FunctionalCondition[*storagev1.StorageClass]{
			ResKind:   "StorageClass",
			Condition: func(c *storagev1.StorageClass) bool { return true },
		},
		&FunctionalCondition[*storagev1.CSIDriver]{
			ResKind:   "CSIDriver",
			Condition: func(c *storagev1.CSIDriver) bool { return true },
		},
	},
	Relations: []resourceConnection{
		{
			Lhs:    0,
			Rhs:    1,
			Exists: true,
			Rel: &FunctionalEqualStringFieldsRelation[*scyllav1.ScyllaCluster, *v1.Pod]{
				LhsExtractor: func(lhs *scyllav1.ScyllaCluster) []string { return []string{lhs.Name} },
				RhsExtractor: func(rhs *v1.Pod) []string { return []string{rhs.Labels["scylla/cluster"]} },
			},
		},
		{
			Lhs:    0,
			Rhs:    2,
			Exists: true,
			Rel: &FunctionalEqualStringFieldsRelation[*scyllav1.ScyllaCluster, *storagev1.StorageClass]{
				LhsExtractor: func(lhs *scyllav1.ScyllaCluster) []string {
					classes := make([]string, 0)
					for _, rack := range lhs.Spec.Datacenter.Racks {
						classes = append(classes, *rack.Storage.StorageClassName)
					}
					return classes
				},
				RhsExtractor: func(rhs *storagev1.StorageClass) []string { return []string{rhs.Name} },
			},
		},
		{
			Lhs:    2,
			Rhs:    3,
			Exists: false,
			Rel: &FunctionalEqualStringFieldsRelation[*storagev1.StorageClass, *storagev1.CSIDriver]{
				LhsExtractor: func(lhs *storagev1.StorageClass) []string { return []string{lhs.Provisioner} },
				RhsExtractor: func(rhs *storagev1.CSIDriver) []string { return []string{rhs.Name} },
			},
		},
	},
}
