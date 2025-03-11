package selector

import (
	"fmt"
	"github.com/scylladb/scylla-operator/pkg/analyze/selector/engine"
	"reflect"
)

type Selector struct {
	relations engine.Relations
	filter    map[string]*engine.Predicate
	nilable   map[string]bool
}

func Type[T any]() reflect.Type {
	return reflect.TypeFor[T]()
}

func New() *Selector {
	return &Selector{relations: engine.NewRelations()}
}

func Select(name string, typ reflect.Type, filter any) *Selector {
	return New().Select(name, typ, filter)
}

func SelectWithNil(name string, typ reflect.Type, filter any) *Selector {
	return New().SelectWithNil(name, typ, filter)
}

func (s *Selector) Select(name string, typ reflect.Type, filter any) *Selector {
	if !s.relations.Add(name, typ) {
		panic(fmt.Sprintf("%s already defined", name))
	}

	if filter != nil {
		p, err := engine.NewPredicate(name, filter)
		if err != nil {
			panic(err)
		}

		s.filter[name] = p
	}

	s.nilable[name] = false

	return s
}

func (s *Selector) SelectWithNil(name string, typ reflect.Type, filter any) *Selector {
	s.Select(name, typ, filter)

	s.nilable[name] = true

	return s
}

func (s *Selector) Relate(first, second string, lambda any) *Selector {
	relation, err := engine.NewRelation(first, second, lambda)
	if err != nil {
		panic(err)
	}

	if !s.relations.Relate(relation) {
		panic("Invalid relation")
	}

	return s
}

func (s *Selector) Where(name string, lambda any) *Selector {
	predicate, err := engine.NewPredicate(name, lambda)
	if err != nil {
		panic(err)
	}

	if !s.relations.Relate(predicate) {
		panic("Invalid predicate")
	}

	return s
}

type Dummy struct{}

// Constructs a Selector which returns a slice of matches
func (s *Selector) Collect() *Dummy {
	// TODO: Implement

	return &Dummy{}
}

// Constructs a Selector which returns a slice of at most limit matches
func (s *Selector) Take(limit int) *Dummy {
	// TODO: Implement

	return &Dummy{}
}

// Constructs a Selector which calls callback for every match
func (s *Selector) ForEach(callback func(map[string]any) (bool, error)) *Dummy {
	// TODO: Implement

	return &Dummy{}
}
