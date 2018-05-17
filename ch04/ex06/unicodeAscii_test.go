package ex06

import (
  "testing"
)

func TestUnicodeSpaceToAscii(t *testing.T) {
  unicodes := "a\tb\t\tc  d e"
  expected := []byte("a b  c  d e")
  asciis := UnicodeSpaceToAscii([]byte(unicodes))
  if !sliceEqual(asciis, expected) {
    t.Errorf("error!\n got     : %s\n expected: %s", asciis, expected)
  }
}

func sliceEqual(slice []byte, expected []byte) bool {
  for i, r := range slice {
    if r != expected[i] {
      return false
    }
  }
  return true
}
