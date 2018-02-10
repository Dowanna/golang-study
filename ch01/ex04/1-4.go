package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	sources := make(map[string][]string)

	// get fileName, first
	fileNames := os.Args[1:]

	for _, fileName := range fileNames {
		// open file and get all contents with ioutil (separated by \n)
		data, err := ioutil.ReadFile(fileName)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		}

		// range and loop through data, and add to map
		for _, content := range strings.Split(string(data), "\n") {
			// ignore empty line
			if content == "" {
				continue
			}

			counts[content]++

			// if sourceFileName does not exist, append
			if !contains(sources[content], fileName) {
				sources[content] = append(sources[content], fileName)
			}
		}
	}

	// range and loop through map, and print
	for line, n := range counts {
		fmt.Printf("%d\t%s\t%v\n", n, line, sources[line])
	}
}

func contains(array []string, s string) bool {
	for _, v := range array {
		if v == s {
			return true
		}
	}
	return false
}
