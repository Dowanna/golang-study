package main

import (
"fmt"
"os"
"strconv"
)

func main() {
	for index, element := range os.Args[0:] {
		fmt.Println(strconv.Itoa(index) + ":" + element)
	}
}
