package ex02

import (
	"testing"
)

func TestAddAll(t *testing.T) {
	var testDatas = [][]int{
		{1, 2, 3},
		{100, 200, 300, 400, 1, 1000, 2, 5},
		{1, 1, 1},
	}

	for _, td := range testDatas {
		var input, expected IntSet

		// setup data
		input.AddAll(td...)
		for _, datum := range td {
			expected.Add(datum)
		}

		// execute test
		if expected.String() != input.String() {
			t.Errorf("error! expected %s, got %s", expected.String(), input.String())
		}
	}

}
