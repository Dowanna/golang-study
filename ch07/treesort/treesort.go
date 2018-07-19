package treesort

import (
	"bytes"
	"strconv"
)

//!+
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place.
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}

	var b bytes.Buffer

	if s := t.left.String(); len(s) > 0 {
		b.WriteString(s)
		b.WriteString(" ")
	}

	b.WriteString(strconv.Itoa(t.value))

	if s := t.right.String(); len(s) > 0 {
		b.WriteString(" ")
		b.WriteString(s)
	}
	return b.String()
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

//!-
