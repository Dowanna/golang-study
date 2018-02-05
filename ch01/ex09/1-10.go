package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	// create output file
	f, err := os.OpenFile("hoge.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
		return
	}
	defer f.Close()

	// get urls from args
	start := time.Now()
	ch := make(chan string)

	// call fetch from args
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		// fmt.Println(<-ch)
		f.WriteString(<-ch)
		f.WriteString("\n")
	}
	f.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
	// get current time
	start := time.Now()

	// get response, check error
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	// get bytesize with io.Copy (1st arg is ioutil.Discard), check error
	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	// get elapsed time, throw to channel
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs fromTheUrl: %s \n%s", secs, url, bytes)
}
