package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello users, this is  my personal calculator in which you can perfrorm basic calculations.")
	fmt.Println("You can perform addition, subtraction, multiplication, and division.")
	fmt.Println("Please select the operation you want to perform:")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		var a, b float64
		fmt.Scan(&a, &b)
		fmt.Println("Result:", a+b)
	case 2:
		var a, b float64
		fmt.Scan(&a, &b)
		fmt.Println("Result:", a-b)
	case 3:
		var a, b float64
		fmt.Scan(&a, &b)
		fmt.Println("Result:", a*b)
	case 4:
		var a, b float64
		fmt.Scan(&a, &b)
		if b != 0 {
			fmt.Println("Result:", a/b)
		} else {
			fmt.Println("Error: Division by zero")
		}
	}
	fmt.Println("Thank you for using my calculator. Goodbye!")
}
