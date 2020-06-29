package main

import (
	"fmt"

	"github.com/learngo/bbd/mydict"
)

func main() {
	dict := mydict.Dictionary{}
	err := dict.Add("hello", "greeting")
	if err != nil {
		fmt.Println("error")
	}
	word, err := dict.Search("hello")
	fmt.Println(word)
}
