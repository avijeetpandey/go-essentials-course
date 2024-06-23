package main

import "fmt"

func using_switch_statements() {
	var number int
	fmt.Print("Please enter a number: ")
	fmt.Scan(&number)

	switch number {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Oops wrong number")
	}
}
