package rules

import (
	"github.com/scylladb/scylla-operator/pkg/analyze/selector"
	"github.com/scylladb/scylla-operator/pkg/analyze/symptoms"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
)

var DummySymptoms = []symptoms.SymptomTreeNode{
	buildBasicDummySymptoms(),
}

func buildBasicDummySymptoms() symptoms.SymptomTreeNode {
	emptyCluster := symptoms.NewSymptom("cluster", "cluster diagnosis", "cluster suggestion",
		selector.Select("cluster", selector.Type[*scyllav1.ScyllaCluster](), nil))
	basicNode := symptoms.NewSymptomTreeLeaf("basic", emptyCluster)

	return basicNode
}
