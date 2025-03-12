package engine

import (
	"maps"
	"reflect"
	"slices"
)

// A Relations is a container for Relation instances
// It enforces types of Relation's arguments and a single relation per
// unordered pair of kinds
type Relations struct {
	types     map[string]reflect.Type
	relations map[string]map[string]Relation
}

func NewRelations() Relations {
	return Relations{
		types:     make(map[string]reflect.Type),
		relations: make(map[string]map[string]Relation),
	}
}

func (r Relations) Add(name string, typ reflect.Type) bool {
	if _, contains := r.types[name]; contains {
		return false
	}

	r.types[name] = typ
	r.relations[name] = make(map[string]Relation)

	return true
}

func (r *Relations) List() []string {
	return slices.Collect(maps.Keys(r.types))
}

func (r *Relations) Resources() map[string]reflect.Type {
	return r.types
}

func (r Relations) Relate(relation Relation) bool {
	if relation == nil {
		return false
	}

	firstName, firstType := relation.FirstParameter()
	secondName, secondType := relation.SecondParameter()

	if typ, exists := r.types[firstName]; !exists || firstType != typ {
		return false
	}

	if typ, exists := r.types[secondName]; !exists || secondType != typ {
		return false
	}

	if firstName > secondName {
		firstName, secondName = secondName, firstName
	}

	if _, exists := r.relations[firstName][secondName]; exists {
		return false
	}

	r.relations[firstName][secondName] = relation

	return true
}

func (r Relations) Relation(first, second string) Relation {
	if first > second {
		first, second = second, first
	}

	return r.relations[first][second]
}
