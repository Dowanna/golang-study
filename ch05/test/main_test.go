package test

import (
	"testing"

	ex07 "github.com/golang-study/ch05/ex07"
	"golang.org/x/net/html"
)

func TestReadVar(t *testing.T) {
	res, _ := ex07.Outline("https://golang.org")
	if res == "" {
		t.Error()
	}
	_, err := html.Parse(res)
	if err != nil {
		t.Error()
	}

	// fmt.Printf("%s", doc)
}
