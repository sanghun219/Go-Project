package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("Request Fail")

func main() {
	var result = make(map[string]string)
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
		res := "OK"
		err := hitURL(url)
		if err != nil {
			res = "FAIL"
		}
		result[url] = res
	}
	for url, res := range result {
		fmt.Println(url, res)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking URL : ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >= 400 {
		return errRequestFail
	} else {
		return nil
	}

}
