package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan bool)
	people := [5]string{"sang", "jang", "pang", "kong", "tong"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
