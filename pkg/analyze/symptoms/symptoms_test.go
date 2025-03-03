package symptoms

import (
	"github.com/scylladb/scylla-operator/pkg/analyze/snapshot"
	scyllav1 "github.com/scylladb/scylla-operator/pkg/api/scylla/v1"
	v1 "k8s.io/api/core/v1"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"reflect"
	"testing"
)

type dummySymptom struct {
	name          string
	diagnoses     []string
	suggestions   []string
	matchCallback func(snapshot.Snapshot) ([]Issue, error)
}

func newEmptyFakeSymptom(name string) Symptom {
	return &dummySymptom{
		name:        name,
		diagnoses:   []string{"diagnoses"},
		suggestions: []string{"suggestions"},
		matchCallback: func(s snapshot.Snapshot) ([]Issue, error) {
			return nil, nil
		},
	}
}

func newFakeSymptom(name string, match func(snapshot.Snapshot) ([]Issue, error)) Symptom {
	return &dummySymptom{
		name:          name,
		diagnoses:     []string{"diagnoses"},
		suggestions:   []string{"suggestions"},
		matchCallback: match,
	}
}

func (d *dummySymptom) Name() string {
	return d.name
}

func (d *dummySymptom) Diagnoses() []string {
	return d.diagnoses
}

func (d *dummySymptom) Suggestions() []string {
	return d.suggestions
}

func (d *dummySymptom) Match(snapshot snapshot.Snapshot) ([]Issue, error) {
	return d.matchCallback(snapshot)
}

func proxySelector(pairing map[string]string) func(snapshot.Snapshot) []map[string]any {
	return func(s snapshot.Snapshot) []map[string]any {
		objects := make(map[string]any)
		for k, v := range pairing {
			vals := s.All()

			var found any = nil
			for _, objs := range vals {
				for _, obj := range objs {
					val := reflect.ValueOf(obj)
					if val.IsValid() {
						if val.Kind() == reflect.Ptr {
							val = val.Elem()
						}
						name := val.FieldByName("Name")
						if name.IsValid() && name.String() == v {
							found = obj
							break
						}
					}
				}
				if found != nil {
					break
				}
			}

			if found != nil {
				objects[k] = found
			}
		}
		if len(objects) == 0 {
			return nil
		} else {
			return []map[string]any{objects}
		}
	}
}

type fakeSnapshot struct {
	objects map[reflect.Type][]any
}

func (m *fakeSnapshot) Add(obj interface{}) {
	m.objects[reflect.TypeOf(obj)] = append(m.objects[reflect.TypeOf(obj)], obj)
}

func (m *fakeSnapshot) List(objType reflect.Type) []interface{} {
	return m.objects[objType]
}

func (m *fakeSnapshot) All() map[reflect.Type][]interface{} {
	return m.objects
}

func makeNicer(m map[string]any) map[string]string {
	nicer := make(map[string]string)
	for k, v := range m {
		val := reflect.ValueOf(v)
		if val.IsValid() {
			id := ""
			if val.Kind() == reflect.Ptr {
				val = val.Elem()
			}
			fieldNamespace := val.FieldByName("Namespace")
			if fieldNamespace.IsValid() && fieldNamespace.String() == v {
				id = fieldNamespace.String()
			}
			fieldName := val.FieldByName("Name")
			if fieldName.IsValid() && fieldName.String() == v {
				id = fieldNamespace.String() + "." + fieldName.String()
			}
			nicer[k] = id
		}
	}
	return nicer
}

func TestSymptom_Match(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		selector func(snapshot.Snapshot) []map[string]any
		snapshot snapshot.Snapshot
		want     []Issue
		wantErr  bool
	}{
		{
			name: "no issues",
			snapshot: &fakeSnapshot{
				objects: map[reflect.Type][]any{
					reflect.TypeFor[*v1.Pod](): {
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_b",
							},
						},
					},
				},
			},
			selector: proxySelector(map[string]string{
				"pod1": "pod_a",
			}),
			want: []Issue{
				{
					Resources: map[string]any{
						"pod1": &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "duplicate resources",
			snapshot: &fakeSnapshot{
				objects: map[reflect.Type][]any{
					reflect.TypeFor[*v1.Pod](): {
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_b",
							},
						},
					},
				},
			},
			selector: proxySelector(map[string]string{
				"pod1": "pod_a",
				"pod2": "pod_a",
			}),
			want: []Issue{
				{
					Resources: map[string]any{
						"pod1": &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
						"pod2": &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "nil resources",
			snapshot: &fakeSnapshot{
				objects: map[reflect.Type][]any{
					reflect.TypeFor[*v1.Pod](): {
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
					},
				},
			},
			selector: proxySelector(map[string]string{}),
			want:     []Issue{},
			wantErr:  false,
		},
		{
			name: "many different resources should match",
			snapshot: &fakeSnapshot{
				objects: map[reflect.Type][]any{
					reflect.TypeFor[*v1.Pod](): {
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_a",
							},
						},
						&v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_b",
							},
						},
					},
					reflect.TypeFor[*scyllav1.ScyllaCluster](): {
						&scyllav1.ScyllaCluster{
							ObjectMeta: metav1.ObjectMeta{
								Name: "sc_a",
							},
						},
					},
					reflect.TypeFor[*storagev1.StorageClass](): {
						&storagev1.StorageClass{
							ObjectMeta: metav1.ObjectMeta{
								Name: "storage_a",
							},
						},
					},
					reflect.TypeFor[*v1.ServiceAccount](): {
						&v1.ServiceAccount{
							ObjectMeta: metav1.ObjectMeta{
								Name: "sa_a",
							},
						},
					},
				},
			},
			selector: proxySelector(map[string]string{
				"pod1":     "pod_b",
				"sc1":      "sc_a",
				"storage1": "storage_a",
				"sa1":      "sa_a",
			}),
			want: []Issue{
				{
					Resources: map[string]any{
						"pod1": &v1.Pod{
							ObjectMeta: metav1.ObjectMeta{
								Name: "pod_b",
							},
						},
						"sc1": &scyllav1.ScyllaCluster{
							ObjectMeta: metav1.ObjectMeta{
								Name: "sc_a",
							},
						},
						"storage1": &storagev1.StorageClass{
							ObjectMeta: metav1.ObjectMeta{
								Name: "storage_a",
							},
						},
						"sa1": &v1.ServiceAccount{
							ObjectMeta: metav1.ObjectMeta{
								Name: "sa_a",
							},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// given
			t.Parallel()
			s := NewSymptom("symptom", "diag", "suggestions", tc.selector)

			// when
			got, err := s.Match(tc.snapshot)

			// then
			if (err != nil) != tc.wantErr {
				t.Errorf("Match() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
			if len(tc.want) != len(got) {
				t.Fatalf("Match() issue length mismatch - got = %v, want %v", got, tc.want)
			}
			for i, want := range tc.want {
				if len(want.Resources) != len(got[i].Resources) {
					t.Fatalf("Match() issue resources length mismatch - got[%d] = %v, want %v", i, makeNicer(got[i].Resources), makeNicer(want.Resources))
				}
				if !reflect.DeepEqual(want.Resources, got[i].Resources) {
					t.Errorf("Match() issue resources mismatch - got = %v, want %v", got[i].Resources, want.Resources)
				}
			}
		})
	}
}

func TestNewEmptySymptomSet_IsEmpty(t *testing.T) {
	t.Parallel()
	// given
	expectedName := "dummySet"

	// when
	set := NewEmptySymptomSet(expectedName)

	// then
	if set.Name() != expectedName {
		t.Errorf("name differs - got %s, wwant %s", set.Name(), expectedName)
	}
	if len(set.Symptoms()) > 0 {
		t.Errorf("Symptoms() is not empty - got %v", set.Symptoms())
	}
	if len(set.DerivedSets()) > 0 {
		t.Errorf("DerivedSets() is not empty - got %v", set.DerivedSets())
	}
	if set.Parent() != nil {
		t.Errorf("Parent() is not nil - got %v", set.Parent())
	}
}

func TestNewEmptySymptomSet_IsNotEmpty(t *testing.T) {
	t.Parallel()
	// given
	expectedName := "dummySet"
	child1Name := "child1"
	child2Name := "child2"
	child1 := NewEmptySymptomSet(child1Name)
	child2 := NewEmptySymptomSet(child2Name)
	children := []*SymptomSet{&child1, &child2}

	// when
	set := NewSymptomSet(expectedName, children)

	// then
	if set.Name() != expectedName {
		t.Errorf("name differs - got %s, wwant %s", set.Name(), expectedName)
	}
	if len(set.Symptoms()) > 0 {
		t.Errorf("Symptoms() is not empty - got %v", set.Symptoms())
	}
	if set.Parent() != nil {
		t.Fatalf("Parent() is not nil - got %v", set.Parent())
	}

	for _, child := range children {
		found := false
		for k, ds := range set.DerivedSets() {
			if (*child).Name() == k {
				found = true
				if (*ds).Parent() == nil || (*(*ds).Parent()).Name() != set.Name() {
					t.Errorf("wrong parent for %s - got %v, want %v", k, (*ds).Parent(), &set)
				}
				if len((*ds).Symptoms()) > 0 {
					t.Errorf("Symptoms() is not empty for %s - got %v", k, set.Symptoms())
				}
				if len((*ds).DerivedSets()) > 0 {
					t.Errorf("DerivedSets() is not empty for %s - got %v", k, set.DerivedSets())
				}
			}
		}
		if !found {
			t.Errorf("child missing: %s", (*child).Name())
		}
	}
}

func TestSymptomSet_Add_ShouldAddValidSymptom(t *testing.T) {
	t.Parallel()
	// given
	ss := NewEmptySymptomSet("symptomSet")
	s := newFakeSymptom("symptom", func(snapshot.Snapshot) ([]Issue, error) { return nil, nil })

	// when
	err := ss.Add(&s)

	// then
	if err != nil {
		t.Fatalf("add shouldn't return an error %v", err)
	}
	if len(ss.Symptoms()) != 1 {
		t.Fatalf("symptoms length mismatch, got %d want 1", len(ss.Symptoms()))
	}
	if _, ok := ss.Symptoms()["symptom"]; !ok {
		t.Fatalf("symptoms should contain symptom, got %d want 1", len(ss.Symptoms()))
	}
	if (*ss.Symptoms()["symptom"]).Name() != "symptom" {
		t.Errorf("symptom name invalid, got %s want symptom", (*ss.Symptoms()["symptom"]).Name())
	}
}

func TestSymptomSet_Add_ShouldReturnAnErrorGivenNil(t *testing.T) {
	t.Parallel()
	// given
	ss := NewEmptySymptomSet("symptomSet")

	// when
	err := ss.Add(nil)

	// then
	if err == nil {
		t.Fatalf("add should return an error %v", err)
	}
}

func TestSymptomSet_AddChild_ShouldAddValidChild(t *testing.T) {
	t.Parallel()
	// given
	ss := NewEmptySymptomSet("symptomSet1")
	ss2 := NewEmptySymptomSet("symptomSet2")

	// when
	err := ss.AddChild(&ss2)

	// then
	if err != nil {
		t.Fatalf("AddChild shouldn't return an error %v", err)
	}
	if len(ss.DerivedSets()) != 1 {
		t.Fatalf("DerivedSets length mismatch, got %d want 1", len(ss.DerivedSets()))
	}
	if _, ok := ss.DerivedSets()["symptomSet2"]; !ok {
		t.Fatalf("DerivedSets should contein symptomSet2")
	}
	if (*ss.DerivedSets()["symptomSet2"]).Name() != "symptomSet2" {
		t.Errorf("DerivedSets name invalid, got %s want symptom", (*ss.DerivedSets()["symptomSet2"]).Name())
	}
}

func TestSymptomSet_AddChild_ShouldReturnAnErrorGivenNil(t *testing.T) {
	t.Parallel()
	// given
	ss := NewEmptySymptomSet("symptomSet")

	// when
	err := ss.AddChild(nil)

	// then
	if err == nil {
		t.Fatalf("add should return an error %v", err)
	}
}
