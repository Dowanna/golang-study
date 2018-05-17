package ex05

import (
  "testing"
)

func TestRemoveDup(t *testing.T) {
  s := []string{"hoge", "hoge", "hoge","fuga"}
  expected := []string{"hoge", "fuga"}
  s = RemoveDup(s)
  if !sliceEqual(s, expected) {
    t.Errorf("reverse failed!\n got: %s\n expected: %s\n", s, expected)
  }
}

func sliceEqual(slice []string, expected []string) bool {
  for i, r := range slice {
    if r != expected[i] {
      return false
    }
  }
  return true
}
