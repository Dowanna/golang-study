package main

import (
	"strings"
	"testing"

	"github.com/golang-study/ch05/testCommon/equals"
	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {
	testData := make(map[string]map[string]int, 0)

	testData["<html><div>'hogehoge and piyo'</div><img src='here.gif' /></html>"] = map[string]int{
		"words":  3,
		"images": 1,
	}

	testData[`
  <html>
    <div>
      <div>'inner nested string'</div>
      <img src='innerImage.jpg'/>
      'hogehoge and piyo'
    </div>
    <img src='here.gif' />
  </html>`] = map[string]int{
		"words":  6,
		"images": 2,
	}

	for in, exp := range testData {
		doc, _ := html.Parse(strings.NewReader(in))
		res := CountWordsAndImages(make(map[string]int, 0), doc)

		if !equals.MapEquals(res, exp) {
			t.Errorf("failed! expected: %v\n got: %v", exp, res)
		}
	}
}
