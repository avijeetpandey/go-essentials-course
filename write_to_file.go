package main

import (
	"fmt"
	"os"
)

func write_to_file(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}
