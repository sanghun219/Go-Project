package main

import (
	"fmt"

	accounts "github.com/learngo/bbd/Accounts"
)

func main() {

	Acc := accounts.NewAccount("sanghun")
	Acc.Deposit(10)

	fmt.Println(Acc.Balance())
	fmt.Println(Acc)
}
