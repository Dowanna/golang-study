package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Print(anagram(os.Args[1]))
}

func anagram(s string) bool {
	buff := []byte(s)

	if len(buff) <= 1 {
		return true
	}

	j := len(buff) - 1

	for i := 0; i < len(buff); i++ {

		if i >= j {
			return true
		}

		if buff[i] != buff[j] {
			fmt.Printf("mismatch characters were found! %v and %v", buff[i], buff[j])
			return false
		}

		j--
		continue
	}
	return true
}
