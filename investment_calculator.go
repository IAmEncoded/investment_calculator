package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	//fmt.Print("Please Enter your Investment Amount: ")
	outputText("Please Enter your Investment Amount: ")
	fmt.Scan(&investmentAmount)

	outputText("Please Enter your years: ")
	fmt.Scan(&years)

	outputText("Please Enter your Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFFV := fmt.Sprintf("Future Real Value: %.1f\n", futureRealValue)
	// fmt.Println("Future Value: ", futureValue)
	// fmt.Printf("Future Value: %.1f\nFuture Real Value: %.1f", futureValue, futureRealValue)

	fmt.Print(formattedFV)
	fmt.Print(formattedFFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv = fv / math.Pow(1 + inflationRate / 100, years)

	return fv, rfv
}