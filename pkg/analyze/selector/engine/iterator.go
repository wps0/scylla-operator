// Package engine provides utilities for the parent selector package. It should
// not be used directly outside of that package.
//
// The package proviedes extracts sets of elements from provided sets of named
// sets such that all elements are in relation with one another for all
// provided relations. Those relations are represented as functions from
// pairs of elements from diffrent or same sets to a boolean value.
package engine

import (
	"cmp"
	"fmt"
	"maps"
	"slices"
)

// An Iterator provides means to iterate over all solutions.
type Iterator struct {
	labels    []string
	relations Relations
	values    map[string][]any
}

func NewIterator(relations Relations, values map[string][]any) (*Iterator, error) {
	labels := relations.List()

	for _, label := range labels {
		if _, contains := values[label]; !contains {
			return nil, fmt.Errorf("Missing key %s", label)
		}
	}

	slices.SortFunc(labels, func(lhs, rhs string) int {
		return cmp.Compare(len(values[lhs]), len(values[rhs]))
	})

	return &Iterator{
		labels:    labels,
		relations: relations,
		values:    values,
	}, nil
}

func (it *Iterator) ForEach(callback func(map[string]any) (bool, error)) error {
	_, err := it.forEach(make(map[string]any, len(it.labels)), callback)
	return err
}

func (it *Iterator) forEach(
	prefix map[string]any,
	callback func(map[string]any) (bool, error),
) (bool, error) {
	if len(prefix) >= len(it.labels) {
		return callback(maps.Clone(prefix))
	}

	label := it.labels[len(prefix)]
	for _, value := range it.values[label] {
		prefix[label] = value

		canAppend, err := it.canAppend(prefix, label, value)
		if err != nil {
			return false, err
		}

		if canAppend {
			continu, err := it.forEach(prefix, callback)

			if !continu || err != nil {
				return false, err
			}
		}

		delete(prefix, label)
	}

	return true, nil
}

func (it *Iterator) canAppend(
	selection map[string]any,
	newLabel string,
	newValue any,
) (bool, error) {
	for otherLabel, otherValue := range selection {
		relation := it.relations.Relation(otherLabel, newLabel)

		if otherValue != nil && newValue != nil {
			if relation == nil {
				continue
			}

			related, err := relation.Check(
				otherLabel, otherValue,
				newLabel, newValue,
			)
			if !related || err != nil {
				return false, err
			}
		} else if otherValue != nil && newValue == nil {
			result, err := checkRelationWithNil(
				otherLabel, otherValue, newLabel, it.values[newLabel], relation,
			)

			if !result || err != nil {
				return false, err
			}
		} else if otherValue == nil && newValue != nil {
			result, err := checkRelationWithNil(
				newLabel, newValue, otherLabel, it.values[otherLabel], relation,
			)

			if !result || err != nil {
				return false, nil
			}
		} else if relation != nil {
			return false, nil
		}
	}

	return true, nil
}

func checkRelationWithNil(
	presentLabel string,
	presentValue any,
	absentLabel string,
	absentValues []any,
	relation Relation,
) (bool, error) {
	if relation == nil {
		return true, nil
	}

	for _, absentValue := range absentValues {
		if absentValue == nil {
			continue
		}

		related, err := relation.Check(presentLabel, presentValue, absentLabel, absentValue)
		if related || err != nil {
			return false, err
		}
	}

	return true, nil
}
