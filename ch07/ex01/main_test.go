package ex01

import (
	"testing"
)

func TestWrite(t *testing.T) {
	for _, td := range []struct {
		input    []byte
		expected int
	}{
		{[]byte("hello world"), 2},
		{[]byte("fuga hoge piyo"), 3},
		{[]byte("こんにちは　おはよう　おやすみ　あ"), 4},
	} {
		var c WordCounter
		c.Write(td.input)
		if c != WordCounter(td.expected) {
			t.Errorf("error! expected word count %d, got %d. Input was '%v'\n", td.expected, c, string(td.input))
		}
	}
}
