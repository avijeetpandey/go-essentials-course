package main

import (
	"fmt"
)

func takeIntInput(message string, value *int) {
	fmt.Print(message)
	fmt.Scan(&value)
}

func calculate_profit() {
	var expenses int
	var revenue int
	var taxRate int

	takeIntInput("Enter the expenses: ", &expenses)
	takeIntInput("Enter the revenue: ", &revenue)
	takeIntInput("Enter the tax rate: ", &taxRate)

	ebt := revenue - expenses

	fmt.Println(ebt)
}
