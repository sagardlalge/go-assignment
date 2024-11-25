package main

import "fmt"

func main() {
	var taxrate, Revenue, expenses float64
	var ebt, profit float64

	fmt.Print("Enter the taxrate (%) : ")
	fmt.Scanln(&taxrate)

	fmt.Print("Enter the Revenue =")
	fmt.Scanln(&Revenue)

	fmt.Print("Enter the expenses = ")
	fmt.Scanln(&expenses)

	ebt = Revenue - expenses

	profit = ebt * (1 - taxrate/100)

	fmt.Printf("EBT (Earnings Before Tax): %.3f\n", ebt)
	fmt.Printf("Profit after tax: %.3f\n", profit)

}
