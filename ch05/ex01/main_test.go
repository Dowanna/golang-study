package main

import (
	"strings"
	"testing"

	"github.com/golang-study/ch05/testCommon/equals"
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

		if !equals.SliceEquals(res, expected) {
			t.Errorf("expected: %v\n got: %v\n", expected, res)
		}
	}
}
