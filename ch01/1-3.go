package main

import (
	"fmt"
	"os"
	"time"	
	"strings"
	"strconv"
)

func main() {
	slow()
	fast()
}

func fast() {
	start := time.Now()
	fmt.Println(strings.Join(os.Args[1:],""))
	fmt.Println(strconv.Itoa(int(time.Since(start).Nanoseconds())) + " nanoseconds")
}

func slow() {
	s, sep, start := "", "", time.Now()
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(strconv.Itoa(int(time.Since(start).Nanoseconds())) + " nanoseconds")
}
