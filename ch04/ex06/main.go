package ex06

import (
  "fmt";
  "unicode";
  "unicode/utf8"
)

func main() {
  unicodes := "a\tb\t\tc  d e"
  asciis := UnicodeSpaceToAscii([]byte(unicodes))
  fmt.Printf("%s", asciis)
}

func UnicodeSpaceToAscii(unicodes []byte) []byte {
  var size int
  var asciis []byte

  for i := 0; i < len(unicodes); i += size {
    rRune, s := utf8.DecodeRune(unicodes[i:])
    size = s

    if unicode.IsSpace(rRune) {
      asciis = append(asciis, byte(0x20))
    } else {
      asciis = append(asciis, unicodes[i:i+s]...)
    }
  }

  return asciis
}
