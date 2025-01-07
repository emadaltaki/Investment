package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2
	var investmentAmount float64
	expectedReturn := 20.0
	years := 10
	a, v := fmt.Scan(&investmentAmount, &expectedReturn, &years)
	fmt.Println(a, v)
	futureValue := float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	futureRealValue := futureValue / (math.Pow(1+inflationRate/100, float64(years)))
	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
