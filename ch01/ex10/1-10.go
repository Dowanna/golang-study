package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Sprint(err)
	}

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}

	for range os.Args[1:] {
		f.WriteString(<-ch + "\n")
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	defer resp.Body.Close() // 資源をリークさせない

	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルへ送信
		return
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	body, err := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2fs from url: %s \n %v", secs, url, string(body))
}
