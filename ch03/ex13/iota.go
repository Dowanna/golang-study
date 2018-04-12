package main

import "fmt"

func main() {
	const (
		kb = 1000
		mb = kb * kb
		gb = mb * kb
		tb = gb * kb //以下同文
	)
	fmt.Printf("%d", kb)
	fmt.Printf("%d", mb)
	fmt.Printf("%d", gb)
	fmt.Printf("%d", tb)
}
