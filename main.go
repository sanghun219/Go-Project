package main

import (
	"errors"
	"fmt"
	"net/http"
)

type result struct {
	url    string
	status string
}

var errRequestFail = errors.New("Request Fail")

func main() {
	var results = make(map[string]string)
	ReceiveResult := make(chan result)

	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitURL(url, ReceiveResult)
	}
	for i := 0; i < len(urls); i++ {
		result := <-ReceiveResult
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)

	}
}

func hitURL(url string, c chan<- result) {
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		c <- result{url: url, status: "FAIL"}
	} else {
		c <- result{url: url, status: "SUCCESS"}
	}

}
