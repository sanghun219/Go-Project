package main

import (
	"fmt"

	accounts "github.com/learngo/bbd/Accounts"
)

func main() {

	Acc := accounts.NewAccount("sanghun")
	fmt.Println(Acc)
}
