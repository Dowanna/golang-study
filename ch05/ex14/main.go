package main

import (
	"fmt"
)

var host string
var paths []string

const depth = 3 //深く潜りすぎてテザリング容量が死にかけるので制限かける

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures": {"discrete math"},
	// "databases":             {"data structures": true},
	// "discrete math":         {"intro to programming": true},
	// "formal languages":      {"discrete math": true},
	// "networks":              {"operating systems": true},
	// "operating systems":     {"data structures": true, "computer organization": true},
	// "programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	host = "https://golang.org"
	breadthFirst(crawl, []string{host})
}

func crawl(courses map[string][]string) []string {

	fmt.Println("%s", courses)
	// if len(paths) > depth {
	// 	return nil
	// }
	//
	// list, _ := links.Extract(url)

	return prereqs[course]
}

func breadthFirst(f func(item map[string][]string) []string, worklist map[string][]string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for key, item := range items {
			if !seen[key] {
				seen[key] = true
				worklist[key] = f(item)
			}
			// worklist = append(worklist, f(item))
		}
	}
	// seen := make(map[string]bool)
	// for len(worklist) > 0 {
	// 	items := worklist
	// 	worklist = nil
	// 	for _, item := range items {
	// 		if !seen[item] {
	// 			seen[item] = true
	// 			worklist = append(worklist, f(item)...)
	// 		}
	// 	}
	// }
}
