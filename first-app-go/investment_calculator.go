package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	fmt.Println("Hello, World!")
	var investmentAmount float64
	expectedReturnRate := 5.5
	var years float64

	fmt.Printf("Investment Amount:")
	fmt.Scanln(&investmentAmount)
	fmt.Printf("Expected Annual Return Rate (%%):")
	fmt.Scanln(&expectedReturnRate)
	fmt.Printf("Number of Years:")
	fmt.Scanln(&years)

	// futureValue := calculateFutureValue(investmentAmount, annualInterestRate, years)
	// fmt.Printf("Future Value of Investment: %.2f\n", futureValue)
	futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)
	fmt.Printf("Future Real Value of Investment: %.2f\n", futureRealValue)
}

func calculateFutureValue(principal float64, annualInterestRate float64, years float64) float64 {
	futureValue := principal * math.Pow(1+annualInterestRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
	return futureRealValue
}
