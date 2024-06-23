package main

import (
	"fmt"
	"go-essentials/fileops"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	balance, err := fileops.ReadFloatFromFile("balance.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(balance)

	fileops.WriteFloatToFile(100000)

	// instantiating user

	appUser := user{
		firstName: "Avijeet",
		lastName:  "pandey",
		birthDate: "11/08/1998",
		createdAt: time.Now(),
	}

	fmt.Println(appUser)
}
