package main

import (
	"fmt"
	"math"
)

func main() {
	var investmentAmount float64 = 1000
	expectedReturn := 5.0
	years := 10
	var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	fmt.Println(futureValue)
}
