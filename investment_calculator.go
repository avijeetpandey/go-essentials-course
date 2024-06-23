package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount = 1000
	years := 5.5
	expectedReturnRate := 10.0

	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years))
	fmt.Println(futureValue)
}
