package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	urls := getUrls(os.Args[1])

	start := time.Now()
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch) // ゴルーチンを開始
	}

	for range urls {
		fmt.Println(<-ch) // ch チャネルから受信
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
func fetch(url string, ch chan<- string) {
	if !strings.HasPrefix(url, "https://") {
		url = "https://www." + url
	}

	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err) // ch チャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

func getUrls(fileName string) []string {
	urls, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Fprint(os.Stdout, "んなファイルはねぇ\n")
		os.Exit(1)
	}

	return strings.Split(string(urls), "\n")
}
