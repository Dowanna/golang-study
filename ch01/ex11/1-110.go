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
	start := time.Now()
	ch := make(chan string)

	f, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer f.Close()

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
	}

	urls := strings.Split(string(data), "\n")

	for _, url := range urls {
		fmt.Fprint(os.Stdout, "%s\n", url)
		go fetch(url, ch) // ゴルーチンを開始
	}

	for range urls {
		fmt.Fprintf(os.Stdout, "write response to file!\n")
		f.WriteString(<-ch)
		f.WriteString("\n")
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "https://") {
		url = "https://" + url
	}

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
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

// func main() {
//

// fetchAll()
// 	ch := make(chan string)
// 	for _, url := range os.Args[1:] {
// 		fmt.Fprintf(os.Stdout, "%s\n", fmt.Sprintf("%s", url))
// 		go fetch(fmt.Sprintf("%s", url), ch)
// 	}
//
// 	defer f.Close()
// }
//
// func fetchAll(urls []string, f *os.File) {
// 	// get urls from args
// 	start := time.Now()
// 	ch := make(chan string)
//
// 	// call fetch from args
// 	for _, url := range urls {
// 		go fetch(url, ch)
// 	}
// 	for range urls {
// 		f.WriteString(<-ch)
// 		f.WriteString("\n")
// 	}
// 	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
// }
//
// func fetch(url string, ch chan<- string) {
// 	// get current time
// 	start := time.Now()
//
// 	// get response, check error
// 	// resp, err := http.Get(fmt.Sprintf("http://%s", url))
// 	resp, err := http.Get(url)
//
// 	if err != nil {
// 		ch <- fmt.Sprint(err)
// 		return
// 	}
//
// 	// get bytesize with io.Copy (1st arg is ioutil.Discard), check error
// 	bytes, err := ioutil.ReadAll(resp.Body)
// 	resp.Body.Close()
//
// 	if err != nil {
// 		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
// 		return
// 	}
//
// 	// get elapsed time, throw to channel
// 	secs := time.Since(start).Seconds()
//
// 	ch <- fmt.Sprintf("%.2fs fromTheUrl: %s \n%s", secs, url, bytes)
// }
