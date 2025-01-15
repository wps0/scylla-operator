package analyze

import (
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
)

type ResourceCondition struct {
	Kind      string
	Condition fields.Selector
}

type EqualFieldsRelation struct {
	LhsPath string
	RhsPath string
}

type Rule struct {
	Diagnosis   string
	Suggestions string
	Resources   []func(dataSource *DataSource) []interface{}
	Relations   []Relation
}

type Relation struct {
	Lhs  int
	Rhs  int
	Eval func(a interface{}, b interface{}) bool
}

var CsiDriverMissingFunctional = Rule{
	Diagnosis:   "local-csi-driver CSIDriver, referenced by <NAME> StorageClass, is missing",
	Suggestions: "deploy local-csi-driver provisioner",
	Resources: []func(dataSource *DataSource) []interface{}{
		func(ds *DataSource) []interface{} {
			clusters, err := ds.ScyllaClusterLister.List(labels.Everything())
			if err != nil {
				panic("err")
			}
			res := make([]interface{}, 0)
			for _, c := range clusters {
				rack := false
				for _, r := range c.Spec.Datacenter.Racks {
					if r.Storage.StorageClassName != nil && *r.Storage.StorageClassName == "scylladb-local-xfs" {
						rack = true
					}
				}

				sscp := false
				prog := false
				for _, cond := range c.Status.Conditions {
					if cond.Type == "StatefulSetControllerProgressing" {
						sscp = true
					} else if cond.Type == "Progressing" {
						prog = true
					}
				}

				if rack && sscp && prog {
					res = append(res, c)
				}
			}
			return res
		},
	},
	Relations: []Relation{
		{
			Lhs: 0,
			Rhs: 1,
			Eval: func(a interface{}, b interface{}) bool {

			},
		},
	},
}
