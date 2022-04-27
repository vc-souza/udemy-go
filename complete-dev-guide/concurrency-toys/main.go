package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkUrl(c chan string, url string) {
	res, err := http.Get(url)
	status := "OK"

	if err != nil || res.StatusCode < 200 || res.StatusCode >= 300 {
		status = "NOT OK"
	}

	fmt.Printf("%v: %v\n", url, status)

	c <- url
}

func main() {
	urls := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	c := make(chan string)

	for _, url := range urls {
		go checkUrl(c, url)
	}

	for msg := range c {
		go func(url string) {
			time.Sleep(time.Second * 5)
			checkUrl(c, url)
		}(msg)
	}
}
