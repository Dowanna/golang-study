package main

import (
	"fmt"
	"testing"
)

var prereqsTestData = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"linear algebra": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func TestTopoSort(t *testing.T) {
	finished := make(map[string]bool, 0)

	for i, course := range TopoSort(prereqsTestData) {
		fmt.Printf("%d:\t%s\n", i, course)
		finished[course] = true
		requiredCourses := prereqsTestData[course]

		// no pre-required courses
		if len(requiredCourses) < 1 {
			continue
		}

		for requiredCourse := range requiredCourses {
			if finished[requiredCourse] == false {
				fmt.Errorf("required course %s is not completed", requiredCourse)
			}
		}
	}

}
