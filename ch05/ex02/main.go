package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// name := os.Args[1]
	// f, err := os.Open(name)
	//
	// if err != nil {
	// 	fmt.Errorf("%v", err)
	// }
	//
	doc, _ := html.Parse(os.Stdin)

	fmt.Println("tag\t\tcount")
	for tag, count := range CountTags(make(map[string]int), doc) {
		fmt.Printf("%s\t\t%d\n", tag, count)
	}
}

func CountTags(count map[string]int, n *html.Node) map[string]int {
	if n == nil {
		return count
	}

	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	CountTags(count, n.FirstChild)
	CountTags(count, n.NextSibling)

	return count
}
