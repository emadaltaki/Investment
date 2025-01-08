package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2
	expectedReturn := 5.5
	var investmentAmount float64
	var years float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter the expected Return Rate: ")
	fmt.Scan(&expectedReturn)
	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	futureRealValue := futureValue / (math.Pow(1+inflationRate/100, float64(years)))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
