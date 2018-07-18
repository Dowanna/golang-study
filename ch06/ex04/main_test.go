package ex04

import (
	"testing"
)

func TestElems(t *testing.T) {
	for _, td := range []struct {
		expected []int
	}{
		{[]int{1, 2, 3}},
		{[]int{1, 65, 89, 9888}},
	} {
		var a IntSet
		a.AddAll(td.expected...)

		elems := a.Elems()

		for _, v := range td.expected {
			if !a.Has(v) {
				t.Errorf("error! couldn't find %v in %v", v, elems)
			}
		}
	}
}
