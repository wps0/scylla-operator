package rules

import (
	"github.com/scylladb/scylla-operator/pkg/analyze/selectors"
	"github.com/scylladb/scylla-operator/pkg/analyze/symptoms"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
)

var DummySymptoms = []symptoms.SymptomTreeNode{
	buildBasicDummySymptoms(),
}

func buildBasicDummySymptoms() symptoms.SymptomTreeNode {
	emptyCluster := symptoms.NewSymptom("cluster", "cluster diagnosis", "cluster suggestion",
		selectors.
			Select("cluster", selectors.Type[*scyllav1.ScyllaCluster]()).
			Filter("cluster", func(c *scyllav1.ScyllaCluster) bool { return c != nil }).
			Collect(symptoms.DefaultLimit))
	basicNode := symptoms.NewSymptomTreeLeaf("basic", emptyCluster)

	return basicNode
}
