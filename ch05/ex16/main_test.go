package ex16

import (
	"testing"
)

func TestJoin(t *testing.T) {
	td := []struct {
		sep      string
		expected string
		val      []string
	}{
		{",", "hoge,hoge,hoge", []string{"hoge", "hoge", "hoge"}},
		{"+", "hoge+fuga+1", []string{"hoge", "fuga", "1"}},
		{"+", "hoge", []string{"hoge"}},
	}

	for _, test := range td {
		res := Join(test.sep, test.val[0:]...)
		if res != test.expected {
			t.Errorf("expected: %s, got %s\n", test.expected, res)
		}
	}
}
