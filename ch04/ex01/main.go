package main

import (
  "os";
  "fmt";
  "log";
  "errors";
  "crypto/sha256"
)

func main() {
  if len(os.Args) != 3 {
    log.Fatal(errors.New("enter two arguments"))
  }
  c1 := sha256.Sum256([]byte(os.Args[1]))
  c2 := sha256.Sum256([]byte(os.Args[2]))

  fmt.Printf("you typed: \n%s\n%s\n", os.Args[1],os.Args[2])
  fmt.Printf("difference of these two: %d\n%x\n%x\n", countSha256Diff(c1,c2), c1, c2)
}

func countSha256Diff(a,b [32]byte) int {
  sum := 0

  for i := range a {
    if a[i] != b[i] {
      sum += 1
    }
  }

  return sum
}
