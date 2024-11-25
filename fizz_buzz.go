package main

import "fmt"

func main() {
	var num int
	fmt.Print("Enter a Number :")
	fmt.Scanln(&num)

	if num%3 == 0 && num%5 == 0 {
		fmt.Println("Fizzbuzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else if num%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println(num)
	}

}
