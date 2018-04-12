package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "1234"
	snew := commaByte(s)
	fmt.Printf("%v", snew)
}

func commaByte(s string) string {
	const (
		segment = 3
	)

	var buffer bytes.Buffer

	length := len(s)
	remainder := length % segment

	if length < segment {
		return s
	}

	if remainder != 0 {
		buffer.WriteString(s[:remainder])
		buffer.WriteByte(',')
	}

	for i := 0; i < length-remainder; i++ {
		if i != 0 && i%segment == 0 {
			buffer.WriteByte(',')
		}
		buffer.WriteByte(s[remainder+i])
	}
	return buffer.String()
}

// これ無限ループになるのでは・・・？
func comma(s string) string {
	fmt.Printf("called comma!")
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3] + "," + s[n-3:])
}
