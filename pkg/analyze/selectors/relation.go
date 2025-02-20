package selectors

type relation struct {
	function[bool]
}

func newRelation(lhs, rhs string, f any) *relation {
	function := newFunction[bool]([]string{lhs, rhs}, f)

	if function == nil {
		return nil
	}

	return &relation{function: *function}
}

func (r *relation) Labels() (string, string) {
	labels := make([]string, 0, 2)
	for label, _ := range r.function.Labels() {
		labels = append(labels, label)
	}
	return labels[0], labels[1]
}

func (r *relation) Check(lhs, rhs labeled[any]) bool {
	return r.Call(map[string]any{lhs.Label: lhs.Value, rhs.Label: rhs.Value})
}
