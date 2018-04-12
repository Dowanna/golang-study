package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

const (
	decimal = "decimal"
	symbol  = "symbol"
	integer = "integer"
)

func main() {
	var buf bytes.Buffer

	separatedString := separateNumbers(os.Args[1])

	buf.WriteString(separatedString[symbol])
	buf.WriteString(separatedString[integer])
	buf.WriteString(separatedString[decimal])

	fmt.Printf("%v", buf.String())
}

func separateNumbers(s string) map[string]string {
	separatedString := make(map[string]string, 0)
	sloc := s
	if strings.HasPrefix(sloc, "-") {
		separatedString[symbol] = "-"
		sloc = sloc[1:]
	}

	decimalIndex := strings.Index(sloc, ".")

	if decimalIndex > -1 {
		separatedString[decimal] = sloc[decimalIndex:]
		sloc = sloc[:decimalIndex]
	}
	separatedString[integer] = commaByte(sloc)
	return separatedString
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
