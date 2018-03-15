package main

import (
	"fmt"
	"goLight/ch02/ex05/popcount"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCount(uint64(312)))
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCountNew(uint64(312)))
}
