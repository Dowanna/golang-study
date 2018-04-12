package main

import "fmt"

func main() {
	randomString := "hoge"
	if randomString[4] != "" {
		fmt.Printf("%v", randomString[4])
	}
}
