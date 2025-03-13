package selector

import (
	"fmt"
	"github.com/scylladb/scylla-operator/pkg/analyze/selector/engine"
	"github.com/scylladb/scylla-operator/pkg/analyze/snapshot"
)

// TODO: come up with a better name
type Match = []map[string]any

type Collector interface {
	Collect() (*Match, error)
	Take(int) (*Match, error)
	ForEach(func(map[string]any) (bool, error)) error
	Clone() Collector
}

type ExecEnvironment struct {
	iterator *engine.Iterator
}

func NewCollector(ss snapshot.Snapshot, selector *Selector) (Collector, error) {
	values := make(map[string][]any)
	resources := selector.relations.Resources()

	for name, resourceType := range resources {
		if selector.nilable[name] {
			values[name] = []any{nil}
		} else {
			filter, err := engine.NewPredicate("o", func(o any) (bool, error) { return true, nil })
			if err != nil {
				return nil, fmt.Errorf("failed to create identity filter: %w", err)
			}
			if pred, exists := selector.filter[name]; exists {
				filter = pred
			}

			ofType := ss.List(resourceType)
			filtered, err := Filter(ofType, filter)
			if err != nil {
				return nil, fmt.Errorf("failed to filter resources for %s: %w", name, err)
			}
			values[name] = filtered
		}
	}

	iterator, err := engine.NewIterator(selector.relations, values)
	if err != nil {
		return nil, fmt.Errorf("failed to create iterator: %w", err)
	}

	var collector Collector = &ExecEnvironment{
		iterator: iterator,
	}
	return collector, nil
}

// Collect Constructs a Selector which returns a slice of matches
func (s *ExecEnvironment) Collect() (*Match, error) {
	matched := make([]map[string]any, 0)

	err := s.iterator.ForEach(func(match map[string]any) (bool, error) {
		matched = append(matched, match)
		return true, nil
	})

	return &matched, err
}

// Take Constructs a Selector which returns a slice of at most limit matches
func (s *ExecEnvironment) Take(limit int) (*Match, error) {
	matched := make([]map[string]any, 0)

	err := s.iterator.ForEach(func(match map[string]any) (bool, error) {
		matched = append(matched, match)
		if len(matched) >= limit {
			return false, nil
		}
		return true, nil
	})

	return &matched, err
}

// ForEach Constructs a Selector which calls callback for every match
func (s *ExecEnvironment) ForEach(callback func(map[string]any) (bool, error)) error {
	return s.iterator.ForEach(callback)
}

func (s *ExecEnvironment) Clone() Collector {
	// TODO: implement
	panic("unimplemented")
	return &ExecEnvironment{}
}
