package main

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	testData := make(map[string][]string)

	testData["<html><body></body></html>"] = []string{}
	testData["<html><body><a href='hoge' /></body></html>"] = []string{"hoge"}
	testData["<html><body><div><div><a href='hoge' /></div><a href='fuga' /></div></body></html>"] = []string{"hoge", "fuga"}

	for d, expected := range testData {
		doc, _ := html.Parse(strings.NewReader(d))

		res := Visit(nil, doc)

		if !equals(res, expected) {
			t.Errorf("expected: %v\n got: %v\n", expected, res)
		}
	}
}

func equals(res, expected []string) bool {
	if len(res) != len(expected) {
		return false
	}

	for i, _ := range expected {
		if res[i] != expected[i] {
			return false
		}
	}
	return true
}
