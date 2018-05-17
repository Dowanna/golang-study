package ex05  

import (
  "fmt"
)

func main() {
  s := []string{"hoge", "piyo", "hoge", "hoge","fuga"}
  s = RemoveDup(s)
  fmt.Printf("%s", s)
}

func RemoveDup(s []string) []string {
  for i, _ := range s {
    // 重複する要素がある限りはループを回し続ける
    for {
      if i < len(s)-1 && s[i] == s[i+1] {
        s = remove(s, i)
      } else {
        break
      }
    }
  }
  return s
}

func remove(slice []string, i int) []string {
  copy(slice[i:], slice[i+1:])
  return slice[:len(slice)-1]
}
