package main

import (
	"fmt"
	"math"
)

func main() {
	var inflationRate float64
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Please enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Please enter expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Please enter expected inflation: ")
	fmt.Scan(&inflationRate)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
