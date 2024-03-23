package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	// investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.2f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	// fmt.Printf("Future Value: %v\nFuture Value (adjusted for inflation): %v", futureValue, futureRealValue)
	// o code below has backticks (``), they serve so that you can break your code into multiple lines
	// fmt.Printf(`Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f`, futureValue, futureRealValue)
	// fmt.Printf("Future Value: %.2f\nFuture Value (adjusted for inflation): %.2f", futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

	// Ch2. Lessons 24 and beyond
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	// return fv, rfv
	return
}

// Pulei 27
