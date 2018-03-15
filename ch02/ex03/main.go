package main

import (
	"fmt"
	"goLight/ch02/ex03/popcount"
	"os"
)

func main() {
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCount(uint64(228)))
	fmt.Fprintf(os.Stdout, "%d\n", popcount.PopCountNew(uint64(228)))
}
