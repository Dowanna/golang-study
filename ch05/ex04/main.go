package main

import (
	"fmt"
	"os"

	. "github.com/golang-study/ch05/errorHandling"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	CheckError(err)
	for _, list := range Visit(nil, doc) {
		fmt.Fprintf(writer, "%s\n", list)
	}
}

var targetAttributes = map[string]string{
	"a":      "href",
	"script": "src",
	"img":    "src",
	"link":   "href",
}

func Visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		a := getAttribute(n, targetAttributes[n.Data])
		if a != "" {
			links = append(links, a)
		}
	}

	links = Visit(links, n.FirstChild)
	return Visit(links, n.NextSibling)
}

func getAttribute(n *html.Node, attr string) string {
	for _, a := range n.Attr {
		if a.Key == attr {
			return a.Val
		}
	}
	return ""
}
