package ex07

import (
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestOutline(t *testing.T) {
	res, _ := Outline("https://golang.org")
	if res == "" {
		t.Error()
	}
	r := strings.NewReader(res)
	_, err := html.Parse(r)
	if err != nil {
		t.Error()
	}

	// fmt.Printf("%s", doc)
}
