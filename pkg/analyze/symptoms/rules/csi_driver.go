package rules

import (
	"github.com/scylladb/scylla-operator/pkg/analyze/selector"
	"github.com/scylladb/scylla-operator/pkg/analyze/symptoms"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
)

var StorageSymptoms = []symptoms.SymptomTreeNode{
	buildLocalCsiDriverMissingSymptoms(),
	buildStorageClassMissingSymptoms(),
}

func buildLocalCsiDriverMissingSymptoms() symptoms.SymptomTreeNode {
	// Scenario #2: local-csi-driver CSIDriver, referenced by scylladb-local-xfs StorageClass, is missing
	csiDriverMissing := symptoms.NewSymptom("CSIDriver is missing",
		"%[csi-driver.Name]% CSIDriver, referenced by %[storage-class.Name]% StorageClass, is missing",
		"deploy %[csi-driver.Name]% provisioner (or change StorageClass)",
		selector.
			Select("scylla-cluster", selector.Type[*scyllav1.ScyllaCluster](), nil).
			Select("storage-class", selector.Type[*storagev1.StorageClass](), nil).
			SelectWithNil("csi-driver", selector.Type[*storagev1.CSIDriver](), nil).
			Relate("scylla-cluster", "storage-class", func(c *scyllav1.ScyllaCluster, sc *storagev1.StorageClass) (bool, error) {
				for _, rack := range c.Spec.Datacenter.Racks {
					if *rack.Storage.StorageClassName == sc.Name {
						return true, nil
					}
				}
				return false, nil
			}).
			Relate("storage-class", "csi-driver", func(sc *storagev1.StorageClass, d *storagev1.CSIDriver) (bool, error) {
				return sc.Provisioner == d.Name, nil
			}))

	csiDriverMissingSymptoms := symptoms.NewSymptomTreeLeaf("csi-driver-missing", csiDriverMissing)
	return csiDriverMissingSymptoms
}

func buildStorageClassMissingSymptoms() symptoms.SymptomTreeNode {
	// Scenario #1: scylladb-local-xfs StorageClass used by a ScyllaCluster is missing
	notDeployedStorageClass := symptoms.NewSymptom("StorageClass is missing",
		"%[storage-class.Name]% StorageClass used by a ScyllaCluster is missing",
		"deploy %[storage-class.Name]% StorageClass (or change StorageClass)",
		selector.
			Select("scylla-cluster", selector.Type[*scyllav1.ScyllaCluster](), func(c *scyllav1.ScyllaCluster) (bool, error) {
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
				return storageClassXfs && conditionProgressing && conditionControllerProgressing, nil
			}).
			SelectWithNil("storage-class", selector.Type[*storagev1.StorageClass](), nil).
			Select("scylla-pod", selector.Type[*v1.Pod](), nil).
			Select("csi-driver", selector.Type[*storagev1.CSIDriver](), nil).
			Relate("scylla-cluster", "scylla-pod", func(c *scyllav1.ScyllaCluster, p *v1.Pod) (bool, error) {
				return c.Name == p.Labels["scylla/cluster"], nil
			}).
			Relate("scylla-cluster", "storage-class", func(c *scyllav1.ScyllaCluster, sc *storagev1.StorageClass) (bool, error) {
				for _, rack := range c.Spec.Datacenter.Racks {
					if *rack.Storage.StorageClassName == sc.Name {
						return true, nil
					}
				}
				return false, nil
			}).
			Relate("storage-class", "csi-driver", func(sc *storagev1.StorageClass, d *storagev1.CSIDriver) (bool, error) {
				return sc.Provisioner == d.Name, nil
			}))

	storageClassMissingSymptoms := symptoms.NewSymptomTreeLeaf("StorageClass missing", notDeployedStorageClass)
	return storageClassMissingSymptoms
}
