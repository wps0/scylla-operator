package analyze

import (
	"context"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	scyllav1alpha1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1alpha1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	"log"
)

type Condition interface {
	EvaluateOn(res interface{}) bool
	Kind() string
}

type Logger interface {
	Info(string, ...any)
}

type TypedCond[T any] struct {
	ResKind   string
	Condition func(context.Context, Logger, T) bool
}

type dummyLogger struct {
}

func (d dummyLogger) Info(s string, a ...any) {
	log.Default().Printf(s, a...)
}

func (r *TypedCond[T]) EvaluateOn(res interface{}) bool {
	return r.Condition(context.Background(), dummyLogger{}, res.(T))
}

func (r *TypedCond[T]) Kind() string {
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

type FuncRelation[L, R any] struct {
	Evaluate func(lhs L, rhs R) bool
}

func (f *FuncRelation[L, R]) EvaluateOn(a interface{}, b interface{}) (bool, error) {
	if a != nil {
		return f.Evaluate(a.(L), b.(R)), nil
	}
	return false, nil
}

type Symptom struct {
	Diagnosis   string
	Suggestions string
	Resources   []Condition
	Relations   []resourceConnection
}

var (
	SCYLLADB_LOCAL_XFS_STORAGECLASS = "scylladb-local-xfs"
)

var NodeConfigNonexistentStorageDevice = Symptom{
	Diagnosis:   "",
	Suggestions: "",
	Resources: []Condition{
		&TypedCond[*scyllav1.ScyllaCluster]{
			ResKind: "ScyllaCluster",
			Condition: func(ctx context.Context, log Logger, cluster *scyllav1.ScyllaCluster) bool {
				containsNotReady := false
				for k, v := range cluster.Status.Racks {
					if v.ReadyMembers < v.Members {
						log.Info("Rack [%s]: only [%d] out of [%d] member(s) are ready.", k, v.ReadyMembers, v.Members)
						containsNotReady = true
					}
				}
				return containsNotReady
			},
		},
		&TypedCond[*v1.Pod]{
			ResKind:   "Pod",
			Condition: func(ctx context.Context, log Logger, _ *v1.Pod) bool { return true },
		},
		&TypedCond[*v1.PersistentVolumeClaim]{
			ResKind: "PersistentVolumeClaim",
			Condition: func(ctx context.Context, log Logger, pvc *v1.PersistentVolumeClaim) bool {
				return pvc.Status.Phase == "Pending"
			},
		},
		&TypedCond[*storagev1.StorageClass]{ // 3
			ResKind: "StorageClass",
			Condition: func(ctx context.Context, log Logger, sc *storagev1.StorageClass) bool {
				return sc.Name == SCYLLADB_LOCAL_XFS_STORAGECLASS
			},
		},
		&TypedCond[*storagev1.CSIDriver]{ // 4
			ResKind: "CSIDriver",
			Condition: func(ctx context.Context, log Logger, csi *storagev1.CSIDriver) bool {
				return true
			},
		},
		&TypedCond[*v1.Pod]{ // 5 - CSIDriver Pod
			ResKind: "Pod",
			Condition: func(ctx context.Context, log Logger, pod *v1.Pod) bool {
				return pod.Labels["app.kubernetes.io/name"] == "local-csi-driver"
			},
		},
		&TypedCond[*scyllav1alpha1.NodeConfig]{ // 6
			ResKind: "NodeConfig",
			Condition: func(ctx context.Context, log Logger, nco *scyllav1alpha1.NodeConfig) bool {
				return nco.Spec.Placement.NodeSelector["scylla.scylladb.com/node-type"] == "scylla"
			},
		},
	},
	Relations: []resourceConnection{
		{
			Lhs: 0,
			Rhs: 1,
			Rel: &FuncRelation[*scyllav1.ScyllaCluster, *v1.Pod]{
				Evaluate: func(lhs *scyllav1.ScyllaCluster, rhs *v1.Pod) bool {
					return lhs.Name == rhs.Labels["scylla/cluster"]
				},
			},
			Exists: true,
		},
		{
			Lhs: 1,
			Rhs: 2,
			Rel: &FuncRelation[*v1.Pod, *v1.PersistentVolumeClaim]{
				Evaluate: func(lhs *v1.Pod, rhs *v1.PersistentVolumeClaim) bool {
					for _, vol := range lhs.Spec.Volumes {
						if vol.PersistentVolumeClaim != nil && vol.PersistentVolumeClaim.ClaimName == rhs.Name {
							return true
						}
					}
					return false
				},
			},
			Exists: true,
		},
		{
			Lhs: 2,
			Rhs: 3,
			Rel: &FuncRelation[*v1.PersistentVolumeClaim, *storagev1.StorageClass]{
				Evaluate: func(lhs *v1.PersistentVolumeClaim, rhs *storagev1.StorageClass) bool {
					return *lhs.Spec.StorageClassName == rhs.Name
				},
			},
			Exists: true,
		},
		{
			Lhs: 3,
			Rhs: 4,
			Rel: &FuncRelation[*storagev1.StorageClass, *storagev1.CSIDriver]{
				Evaluate: func(lhs *storagev1.StorageClass, rhs *storagev1.CSIDriver) bool {
					return lhs.Provisioner == rhs.Name
				},
			},
			Exists: true,
		},
		{
			Lhs: 5,
			Rhs: 6,
			Rel: &FuncRelation[*v1.Pod, *scyllav1alpha1.NodeConfig]{
				Evaluate: func(csiPod *v1.Pod, nodeConfig *scyllav1alpha1.NodeConfig) bool {
					dirsMounted := false
					for _, vol := range csiPod.Spec.Volumes {
						if vol.Name == "volumes-dir" {
							for _, mp := range nodeConfig.Spec.LocalDiskSetup.Mounts {
								if (*vol.HostPath).Path == mp.MountPoint {
									dirsMounted = true
								}
							}
						}
					}
					return dirsMounted
				},
			},
			Exists: false,
		},
	},
}
