package main

import "fmt"

func main() {

	var taxrate, revenue, expense, ebt, profit float64
	fmt.Print("TaxRate % =")
	fmt.Scan(&taxrate)

	fmt.Print(" Revenue =")
	fmt.Scan(&revenue)

	fmt.Print("Expenses =")
	fmt.Scan(&expense)

	ebt = revenue - expense
	profit = ebt * (1 - taxrate/100)

	fmt.Printf("EBT is : %.2f\n", ebt)
	//fmt.Scan(&ebt)

	fmt.Printf("Profit is : %.2f\n", profit)
	//fmt.Scan(&profit)
}
