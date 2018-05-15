package main

import (
  "os";
  "fmt";
  "flag";
  "crypto/sha256";
  "crypto/sha512"
)

func main() {
  var sha = flag.Int("sha", 256, "specify sha to use for encoding")
  flag.Parse()

  printSha(getSha(*sha, os.Args[len(os.Args)-1]))
}

func getSha(sha int, text string) string {
  textByte := []byte(text)

  switch sha {
  case 256:
    fmt.Println("you selected sha256 and your text was", text)
    convertedByte := sha256.Sum256(textByte)
    return BytesToString(convertedByte[:])

    //こう書けないのはなぜ・・・？
    //return BytesToString(sha256.Sum256(textByte)[:]) //./main.go:27:49: invalid operation sha256.Sum256(textByte)[:] (slice of unaddressable value)
  case 384:
    fmt.Println("you selected sha384 and your text was", text)
    convertedByte := sha512.Sum384(textByte)
    return BytesToString(convertedByte[:])
  case 512:
    fmt.Println("you selected sha512 and your text was", text)
    convertedByte := sha512.Sum512(textByte)
    return BytesToString(convertedByte[:])
  default:
    fmt.Println("you selected sha256 and your text was", text)
    convertedByte := (sha256.Sum256(textByte))
    return BytesToString(convertedByte[:])
  }
}

func printSha(text string) {
  fmt.Printf("your text encoded: %x\n", text)
}

func BytesToString(data []byte) string {
	return string(data[:])
}
