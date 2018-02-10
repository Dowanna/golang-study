package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.OpenFile("result.txt", os.O_RDWR|os.O_CREATE, 644)
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprint(os.Stderr, "%v\n", err)
	}

	fetchAll(strings.Split(string(data), "\n"), f)
	defer f.Close()
}

func fetchAll(urls []string, f *os.File) {
	// get urls from args
	start := time.Now()
	ch := make(chan string)

	// call fetch from args
	for _, url := range urls {
		go fetch(url, ch)
	}
	for range urls {
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
		ch <- fmt.Sprintf("error while reading %s: %v", url, err)
		return
	}

	// get elapsed time, throw to channel
	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs fromTheUrl: %s \n%s", secs, url, bytes)
}
