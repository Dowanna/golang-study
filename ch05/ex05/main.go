package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	. "github.com/golang-study/ch05/errorHandling"
	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	CheckError(err)

	counts := CountWordsAndImages(map[string]int{"words": 0, "images": 0}, doc)

	for e, count := range counts {
		fmt.Printf("%s: %d\n", e, count)
	}
}

func CountWordsAndImages(counts map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return counts
	}

	if n.Type == html.TextNode {
		s := bufio.NewScanner(strings.NewReader(n.Data))
		s.Split(bufio.ScanWords)

		for s.Scan() {
			counts["words"]++
		}
	}

	if n.Type == html.ElementNode && n.Data == "img" {
		counts["images"]++
	}

	if !(n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style")) {
		counts = CountWordsAndImages(counts, n.FirstChild)
	}

	return CountWordsAndImages(counts, n.NextSibling)
}
