package ex09

import (
	"testing"
)

func TestExpand(t *testing.T) {
	type TestData struct {
		input    string
		expected string
	}

	tds := []TestData{
		{
			input:    "$hoge$fuga",
			expected: "!hoge?!fuga?",
		},
		{
			input:    "hoge($fuga)",
			expected: "hoge(!fuga?)",
		},
		{
			input:    "here comes an exclamation $here foooo",
			expected: "here comes an exclamation !here? foooo",
		},
	}
	for _, td := range tds {
		res := Expand(td.input, func(s string) string {
			return "!" + s[1:] + "?"
		})
		if res != td.expected {
			t.Errorf("expected %s, got %s", td.expected, res)
		}
	}
}
