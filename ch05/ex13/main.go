package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/gopl.io/ch5/links"
)

var host string
var paths []string

const depth = 3 //深く潜りすぎてテザリング容量が死にかけるので制限かける

func main() {
	host = "https://golang.org"
	breadthFirst(crawl, []string{host})
	// fmt.Printf("%s", paths)
	// for _, path := range paths {
	// 	os.MkdirAll("."+path, 0777)
	// }
}

func crawl(url string) []string {

	if strings.HasPrefix(url, host) {
		if strings.HasSuffix(url, "/") {
			paths = append(paths, strings.Split(url, host)[1])
			os.MkdirAll("."+strings.Split(url, host)[1], 0777)
		} else {
			res, _ := http.Get(url)
			f, _ := os.Create("." + strings.Split(url, host)[1])

			defer f.Close()
			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)

			f.Write(body)
		}
	}

	fmt.Println("%s", paths)
	if len(paths) > depth {
		return nil
	}

	list, _ := links.Extract(url)

	return list
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}
