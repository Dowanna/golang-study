package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("hoge.txt")
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	CountWords(scanner)
}

func CountWords(scanner *bufio.Scanner) {
	counts := make(map[string]int)

	for scanner.Scan() {
		counts[scanner.Text()]++
	}

	fmt.Println("words\tcount")

	for s, i := range counts {
		fmt.Printf("%s\t%d\n", s, i)
	}
}
