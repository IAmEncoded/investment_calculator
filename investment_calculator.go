package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Please Enter your Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Please Enter your years: ")
	fmt.Scan(&years)

	fmt.Print("Please Enter your Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	var futureValue = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)

	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}