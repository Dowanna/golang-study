package ex03

import (
  "testing"
)

func TestReverse(t *testing.T) {
  numbers := []int{1,2,3,4}
  expected := []int{4,2,2,1}

  Reverse(&numbers)

  for i, r := range numbers {
    if r != expected[i] {
      t.Errorf("reverse failed!\n got: %d\n expected: %d\n", numbers, expected)
    }
  }
}
