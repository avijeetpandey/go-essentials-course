package main

import "fmt"

func main() {
	balance, err := read_from_file()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)

	write_to_file(100000)
}
