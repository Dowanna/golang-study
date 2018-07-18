package main

import (
	"testing"
)

func TestIntersectWith(t *testing.T) {
	// setup data
	type dataSet struct {
		input    [][]int
		expected []int
	}
	var tds []dataSet

	tds = append(tds, dataSet{[][]int{{1, 2, 3}, {1, 3, 4}}, []int{1, 3}})
	tds = append(tds, dataSet{[][]int{{1, 5, 8, 9, 10}, {1, 5}}, []int{1, 5}})
	tds = append(tds, dataSet{[][]int{{1, 63, 128, 190}, {63, 32, 16}}, []int{63}})
	tds = append(tds, dataSet{[][]int{{67, 312, 118}, {67}}, []int{67}})

	for _, td := range tds {
		var testIntSet, tempIntSet IntSet

		// setup test data
		for i, input := range td.input {
			if i == 0 {
				testIntSet.AddAll(input...)
			} else {
				tempIntSet.AddAll(input...)
				testIntSet.IntersectWith(&tempIntSet)
			}
		}

		// execute test
		for _, v := range td.expected {
			if !testIntSet.Has(v) {
				t.Errorf("error! couldn't find %d", v)
			}
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	for _, td := range []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1}},
		{[]int{3, 4, 5, 6, 7, 8, 15, 20}, []int{3, 4, 5, 6, 7}, []int{8, 15, 20}},
	} {
		var a IntSet
		var b IntSet
		a.AddAll(td.a...)
		b.AddAll(td.b...)

		a.DifferenceWith(&b)

		for _, v := range td.expected {
			if !a.Has(v) {
				t.Errorf("error! couldn't find %d", v)
			}
		}
	}
}
func TestSymmetricDifferenceWith(t *testing.T) {
	for _, td := range []struct {
		a        []int
		b        []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{2, 3, 4}, []int{1, 4}},
		{[]int{68, 10, 25}, []int{1, 10, 24}, []int{1, 24, 25, 68}},
	} {
		var a IntSet
		var b IntSet
		a.AddAll(td.a...)
		b.AddAll(td.b...)

		a.SymmetricDifferenceWith(&b)

		for _, v := range td.expected {
			if !a.Has(v) {
				t.Errorf("error! couldn't find %v in %v", v, a.words)
			}
		}
	}
}
