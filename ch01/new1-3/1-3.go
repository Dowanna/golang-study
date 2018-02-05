package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fast(os.Args[1:])
	slow(os.Args[1:])
}

func fast(args []string) {
	//start := time.Now()
	fmt.Println(strings.Join(args, " "))
	//fmt.Println(strconv.Itoa(int(time.Since(start).Nanoseconds())) + " nanoseconds")
}

func slow(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	//fmt.Println(strconv.Itoa(int(time.Since(start).Nanoseconds())) + " nanoseconds")
}
