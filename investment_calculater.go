package main

import (
	"fmt"
)

func main() {
	// Invetsment Calculater
	// const inflationRate float64 = 2
	// expectedReturn := 5.5
	// var investmentAmount float64
	// var years float64

	// fmt.Print("Enter the investment amount: ")
	// fmt.Scan(&investmentAmount)
	// fmt.Print("Enter the expected Return Rate: ")
	// fmt.Scan(&expectedReturn)
	// fmt.Print("Enter the number of years: ")
	// fmt.Scan(&years)

	// futureValue := float64(investmentAmount) * math.Pow(1+expectedReturn/100, float64(years))
	// futureRealValue := futureValue / (math.Pow(1+inflationRate/100, float64(years)))
	// fmt.Println(futureValue)
	// fmt.Println(futureRealValue)
	/////////////////////////
	//Tax Calculater

	var revenue float64
	var expenses float64
	var taxRate float64

	outputText("Enter the revenue:")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	// fmt.Println("ebt value: ", ebt)
	// fmt.Println("Profit Value", profit)
	// fmt.Println(ratio)
	/// Note Use Printf to format the output use %v to print the value
	// `` use this for lines
	fmt.Printf(`Profit Value:
	
	%v ratio Value %v1`, profit, ratio)
	/// Note Use Printf to change the dicimal value to integer  /// .0f is no dicimal value .1f is one dicimal value
	// fmt.Printf("profit Value: %.0f\nratio Value %.0f", profit, ratio)

	// //fmt.Sprintf adds a return
	// formatedFV := fmt.Sprintf("\nprofit Value %.0f", profit)
	// fmt.Println(formatedFV)
}

// define a function to output text
func outputText(text string) {
	fmt.Print(text)
}
