package main

import "fmt"

func main() {

	var num1, num2 float64
	fmt.Printf("Enter the first number :")
	fmt.Scanln(&num1)

	fmt.Printf("Enter the Second Number :")
	fmt.Scanln(&num2)

	//fmt.Printf("Addition is :", num1+num2)
	fmt.Printf("Addition is: %.f\n", num1+num2)
}
