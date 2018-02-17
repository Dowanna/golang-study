package main

import (
	"fmt"
	tempconv "goLight/ch02/ex01/tempconv"
)

func main() {
	fmt.Printf("%v", tempconv.CToF(tempconv.AbsoluteZeroC).String())
	fmt.Printf("%v", tempconv.CToK(tempconv.AbsoluteZeroC).String())
}
