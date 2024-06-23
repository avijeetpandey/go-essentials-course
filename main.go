package main

import (
	"fmt"
	"go-essentials/fileops"
)

func main() {
	balance, err := fileops.ReadFloatFromFile("balance.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)

	fileops.WriteFloatToFile(100000)
}
