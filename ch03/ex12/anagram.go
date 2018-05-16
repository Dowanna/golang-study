package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(anagram(os.Args[1]))
}

func anagram(sOrg string) bool {
	s := []rune(sOrg)
	if len(s) <= 1 {
		return true
	}

	j := len(s) - 1

	for i := 0; i < len(s); i++ {
		if i >= j {
			return true
		}

		if s[i] != s[j] {
			fmt.Printf("mismatch characters were found! %v and %v\n", s[i], s[j])
			return false
		}

		j--
		continue
	}
	return true
}

// func anagram(s string) bool {
// 	if len(s) <= 1 {
// 		return true
// 	}
//
// 	j := len(s) - 1
//
// 	for i := 0; i < len(s); i++ {
// 		if i >= j {
// 			return true
// 		}
//
// 		if s[i] != s[j] {
// 			fmt.Printf("mismatch characters were found! %v and %v\n", s[i], s[j])
// 			return false
// 		}
//
// 		j--
// 		continue
// 	}
// 	return true
// }
