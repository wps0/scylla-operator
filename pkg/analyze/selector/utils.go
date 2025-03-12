package selector

import "github.com/scylladb/scylla-operator/pkg/analyze/selector/engine"

func Filter[T any](list []T, filter *engine.Predicate) ([]T, error) {
	filtered := make([]T, 0)
	for _, el := range list {
		r, err := filter.Test(el)
		if err != nil {
			return nil, err
		}

		if r {
			filtered = append(filtered, el)
		}
	}
	return filtered, nil
}
