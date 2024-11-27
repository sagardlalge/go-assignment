package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Enter Number :")
	fmt.Scan(&num)
	if num%3 == 0 && num%5 == 0 {
		//fmt.Printf("Enter Number :")
		//fmt.Scan(&num)
		fmt.Println("FizzBuzz")
	} else if num%5 == 0 {
		//fmt.Printf("Enter a Number: ")
		//fmt.Scan(&num)
		fmt.Println("Buzz")
	} else if num%3 == 0 {
		fmt.Println("Fizz")
	} else {
		fmt.Println("NIL")
	}
}
