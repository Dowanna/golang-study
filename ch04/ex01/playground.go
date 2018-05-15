package main

import (
  "fmt"
)

func main() {
  c1 := []byte("X")
  c3 := [][]string{[]string{"hoge", "fuga", "piyo"}, []string{"moga", "moge"}}

  fmt.Printf("%s and %s", c1, c3)
}
