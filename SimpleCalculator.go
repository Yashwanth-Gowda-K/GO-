package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	for {
		fmt.Println("Simple Calculator in Go")
		fmt.Println("------------------------")

		// Get first number
		fmt.Print("Enter first number: ")
		_, err1 := fmt.Scanln(&num1)
		if err1 != nil {
			fmt.Println("Invalid input for number. Please try again.\n")
			continue // skip to next loop
		}

		// Get operator
		fmt.Print("Enter operator (+, -, *, /): ")
		_, err2 := fmt.Scanln(&operator)
		if err2 != nil {
			fmt.Println("Invalid input for operator. Please try again.\n")
			continue
		}

		// Get second number
		fmt.Print("Enter second number: ")
		_, err3 := fmt.Scanln(&num2)
		if err3 != nil {
			fmt.Println("Invalid input for number. Please try again.\n")
			continue
		}

		// Perform the calculation
		switch operator {
		case "+":
			fmt.Printf("Result: %.2f\n", num1+num2)
		case "-":
			fmt.Printf("Result: %.2f\n", num1-num2)
		case "*":
			fmt.Printf("Result: %.2f\n", num1*num2)
		case "/":
			if num2 == 0 {
				fmt.Println("Error: Division by zero is not allowed. Please try again.\n")
				continue
			}
			fmt.Printf("Result: %.2f\n", num1/num2)
		default:
			fmt.Println("Error: Invalid operator. Please use +, -, *, or /. Try again.\n")
			continue
		}

		fmt.Println() // Add space before next run
	}
}
