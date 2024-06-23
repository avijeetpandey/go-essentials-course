package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse value")
	}

	return value, nil
}

func WriteFloatToFile(value float64) {
	valueText := fmt.Sprint(value)
	os.WriteFile("balance.txt", []byte(valueText), 0644)
}
