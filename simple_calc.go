package main

import (
	"fmt"
	"math"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) string {
	if b != 0 {
		return fmt.Sprintf("%.2f", a/b)
	} else {
		return "Cannot divide by zero"
	}
}

func power(a, b float64) float64 {
	return math.Pow(a, b)
}

func main() {
	var choice int
	var num1, num2 float64

	// Display the menu
	fmt.Println("Select operation:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Division")
	fmt.Println("4. Multiplication")
	fmt.Println("5. Power")

	// Input choice
	fmt.Print("Enter choice (1/2/3/4/5): ")
	fmt.Scan(&choice)

	// Input numbers
	fmt.Print("Enter first number: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second number: ")
	fmt.Scan(&num2)

	// Perform the operation based on choice
	switch choice {
	case 1:
		fmt.Printf("%.2f + %.2f = %.2f\n", num1, num2, add(num1, num2))
	case 2:
		fmt.Printf("%.2f - %.2f = %.2f\n", num1, num2, subtract(num1, num2))
	case 3:
		result := divide(num1, num2)
		fmt.Printf("%.2f / %.2f = %s\n", num1, num2, result)
	case 4:
		fmt.Printf("%.2f * %.2f = %.2f\n", num1, num2, multiply(num1, num2))
	case 5:
		fmt.Printf("%.2f ^ %.2f = %.2f\n", num1, num2, power(num1, num2))
	default:
		fmt.Println("Invalid Input")
	}
}
