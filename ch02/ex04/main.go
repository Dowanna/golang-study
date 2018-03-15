package main

import (
	"fmt"
	"goLight/ch02/ex04/popcount"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCount(uint64(255)))
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCountNew(uint64(255)))
}
