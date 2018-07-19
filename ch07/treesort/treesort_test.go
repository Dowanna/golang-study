package treesort

import (
	"testing"
)

func TestTree_String(t *testing.T) {
	for _, td := range []struct {
		t        *tree
		expected string
	}{
		{t: nil, expected: ""},
		{t: &tree{value: 1}, expected: "1"},
		{t: &tree{value: 3, left: &tree{value: 4}, right: &tree{value: 5}}, expected: "4 3 5"},
		{t: &tree{value: 3, right: &tree{value: 5}}, expected: "3 5"},
	} {
		if td.t.String() != td.expected {
			t.Errorf("expected %s, got %s", td.expected, td.t.String())
		}
	}
}
