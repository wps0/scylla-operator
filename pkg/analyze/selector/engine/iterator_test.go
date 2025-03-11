package engine

import (
	"cmp"
	"fmt"
	"maps"
	"reflect"
	"slices"
	"testing"
)

func NewPredicateUnchecked(label string, f any) *Predicate {
	predicate, err := NewPredicate(label, f)
	if err != nil {
		panic(err)
	}

	return predicate
}

func NewRelationUnchecked(lhs, rhs string, f any) Relation {
	relation, err := NewRelation(lhs, rhs, f)
	if err != nil {
		panic(err)
	}

	return relation
}

func MakeRelations(
	types map[string]reflect.Type,
	relations []Relation,
) Relations {
	result := NewRelations()

	for name, typ := range types {
		if !result.Add(name, typ) {
			panic("Invalid field")
		}
	}

	for _, relation := range relations {
		if !result.Relate(relation) {
			panic("Invalid relation")
		}
	}

	return result
}

func CompareMaps(x, y map[string]any) int {
	result := cmp.Compare(len(x), len(y))
	if result != 0 {
		return result
	}

	keys := slices.Sorted(maps.Keys(x))
	otherKeys := slices.Sorted(maps.Keys(y))
	result = slices.Compare(keys, otherKeys)

	if result != 0 {
		return result
	}

	for _, key := range keys {
		result = cmp.Compare(
			fmt.Sprintf("%+v", x[key]),
			fmt.Sprintf("%+v", y[key]),
		)
		if result != 0 {
			return result
		}
	}

	return 0
}

type ForEachTest struct {
	name      string
	relations Relations
	values    map[string][]any
	expected  []map[string]any
}

func TestForEach(t *testing.T) {
	tests := []ForEachTest{
		{
			name: "no relations",
			relations: MakeRelations(map[string]reflect.Type{
				"A": reflect.TypeFor[int](),
				"B": reflect.TypeFor[bool](),
			}, []Relation{}),
			values: map[string][]any{
				"A": []any{1, 2, 3},
				"B": []any{false, true},
			},
			expected: []map[string]any{
				{"A": 1, "B": false},
				{"A": 1, "B": true},
				{"A": 2, "B": false},
				{"A": 2, "B": true},
				{"A": 3, "B": false},
				{"A": 3, "B": true},
			},
		},
		{
			name: "single relation",
			relations: MakeRelations(map[string]reflect.Type{
				"A": reflect.TypeFor[int](),
				"B": reflect.TypeFor[bool](),
			}, []Relation{
				NewRelationUnchecked("A", "B",
					func(a int, b bool) (bool, error) {
						return (a%2 == 0) == b, nil
					}),
			}),
			values: map[string][]any{
				"A": []any{1, 2, 3},
				"B": []any{false, true},
			},
			expected: []map[string]any{
				{"A": 1, "B": false},
				{"A": 2, "B": true},
				{"A": 3, "B": false},
			},
		},
		{
			name: "single predicate",
			relations: MakeRelations(map[string]reflect.Type{
				"A": reflect.TypeFor[int](),
				"B": reflect.TypeFor[bool](),
			}, []Relation{
				NewPredicateUnchecked("A",
					func(a int) (bool, error) {
						return a%2 != 0, nil
					}),
			}),
			values: map[string][]any{
				"A": []any{1, 2, 3},
				"B": []any{false, true},
			},
			expected: []map[string]any{
				{"A": 1, "B": false},
				{"A": 1, "B": true},
				{"A": 3, "B": false},
				{"A": 3, "B": true},
			},
		},
	}

	for _, test := range tests {
		it, err := NewIterator(test.relations, test.values)

		if it == nil || err != nil {
			t.Errorf("%s: Unexpected error constructing object: %s", test.name, err)
			continue
		}

		result := make([]map[string]any, 0, len(test.expected))
		err = it.ForEach(func(values map[string]any) (bool, error) {
			result = append(result, values)
			return true, nil
		})

		if err != nil {
			t.Errorf("%s: Unexpected error from ForEach: %s", test.name, err)
		}

		slices.SortFunc(test.expected, CompareMaps)
		slices.SortFunc(result, CompareMaps)

		if !slices.EqualFunc(test.expected, result, maps.Equal) {
			t.Errorf("%s: Fail", test.name)

			for i, match := range result {
				t.Logf("%d: %+v", i, match)
			}
		}
	}
}
