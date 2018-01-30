package main

import (
  "fmt"
  "os"
  "bufio"
)

func main() {
  counts := make(map[string]int)
  filenames := os.Args[1:]
  //for _, line := range filenames {
  //  fmt.Println(line)
  //}
  if len(filenames) == 0 {
    countLines(os.Stdin, counts)
  } else {
    for _, arg := range filenames {
      f, err := os.Open(arg)
      if err != nil {
        fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
        continue
      }
      countLines(f, counts)
      f.Close()
    }
  }
  for line, n := range counts {
    fmt.Printf("%d\t%s\n", n, line)
  }
}

func countLines(f *os.File, counts map[string]int) {
  input := bufio.NewScanner(f)
  for input.Scan() {
    counts[input.Text()]++
  }
}
