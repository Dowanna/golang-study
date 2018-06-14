package ex15

import (
	"testing"
)

func TestMax(t *testing.T) {
	td := []struct {
		vals     []int
		expected int
	}{
		{[]int{1, 2, 3}, 3},
		{[]int{10, 2, 3}, 10},
		{[]int{-1, -2, -3}, -1},
	}

	for _, test := range td {
		res := Max(test.vals[0], test.vals[1:]...)
		if res != test.expected {
			t.Errorf("expected %d got %d\n", test.expected, res)
		}
	}
}

func TestMaxRequireArgs(t *testing.T) {
	td := []struct {
		vals       []int
		expected   int
		shouldFail bool
	}{
		{[]int{1, 2, 3}, 3, false},
		{[]int{10, 2, 3}, 10, false},
		{[]int{-1, -2, -3}, -1, false},
		{[]int{}, 0, true},
	}

	for _, test := range td {
		res, e := MaxRequireArgs(test.vals[0:]...)
		if test.shouldFail && e == nil {
			t.Error("dataset should fail because no arguments were given", test.vals)
		}
		if res != test.expected {
			t.Errorf("expected %d got %d\n", test.expected, res)
		}
	}
}

func TestMin(t *testing.T) {
	td := []struct {
		vals     []int
		expected int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{10, 2, 3}, 2},
		{[]int{-1, -2, -3}, -3},
	}

	for _, test := range td {
		res := Min(test.vals[0], test.vals[1:]...)
		if res != test.expected {
			t.Errorf("expected %d got %d\n", test.expected, res)
		}
	}
}
