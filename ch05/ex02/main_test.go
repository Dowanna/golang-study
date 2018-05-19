package main

import (
	"strings"
	"testing"

	"github.com/golang-study/ch05/testCommon/equals"
	"golang.org/x/net/html"
)

func TestCountTags(t *testing.T) {
	testData := make(map[string]map[string]int)

	testData["<html><body></body></html>"] = map[string]int{
		"html": 1,
		"body": 1,
	}

	testData["<html><body><div></div><div></div><div></div><div></div></body></html>"] = map[string]int{
		"html": 1,
		"div":  4,
		"body": 1,
	}

	for d, exp := range testData {
		doc, _ := html.Parse(strings.NewReader(d))

		res := CountTags(make(map[string]int), doc)
		if !equals.MapEquals(res, exp) {
			t.Errorf("failed! expected: %v\n got: %v", exp, res)
		}
	}
}
