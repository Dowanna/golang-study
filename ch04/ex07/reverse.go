package ex07

import (
	"unicode/utf8"
)

func RevUtf8(b []byte) {
	for i := 0; i < len(b); {
		_, s := utf8.DecodeRune(b[i:])
		rev(b[i : i+s])
		i += s
	}
	rev(b)
}

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}
