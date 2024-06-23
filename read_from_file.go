package main

import (
	"errors"
	"os"
	"strconv"
)

func read_from_file() (float64, error) {
	data, err := os.ReadFile("balance.txt")

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse balance value")
	}

	return balance, nil
}
