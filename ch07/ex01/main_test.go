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
		{[]byte("fuga hoge                  piyo"), 3},
		{[]byte("こんにちは　おはよう　おやすみ　あ"), 4},
		{[]byte("onegiganticlongword a s h o r t w o r d"), 11},
	} {
		var c WordCounter
		c.Write(td.input)
		if c != WordCounter(td.expected) {
			t.Errorf("error! expected word count %d, got %d. Input was '%v'\n", td.expected, c, string(td.input))
		}
	}
}

func TestLine(t *testing.T) {
	for _, td := range []struct {
		input    []byte
		expected int
	}{
		{[]byte("あ\nい\nう\n"), 3},
		{[]byte("あ\t\nい\t\nうううう\n"), 3},
		{[]byte("あ\n\n\nう\n"), 4},
	} {
		var c LineCounter
		c.Write(td.input)
		if c != LineCounter(td.expected) {
			t.Errorf("error! expected line count %d, got %d. Input was '%v'\n", td.expected, c, string(td.input))
		}
	}
}
