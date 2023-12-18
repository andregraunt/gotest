package main

import "fmt"

func main() {

	var choice int

	fmt.Println("Select operation")
	fmt.Println("1. Add")
	fmt.Println("2. Subract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")

	fmt.Scanf("%d", &choice)

	var num1, num2 float64

	fmt.Print("Enter First Number: ")
	fmt.Scanf("%f", &num1)

	fmt.Print("Enter Second Number: ")
	fmt.Scanf("%f", &num2)

	switch choice {
	case 1:
		fmt.Println("Result of Addition: ", add(num1, num2))
	case 2:
		fmt.Println("Result of Subtraction: ", subtract(num1, num2))
	case 3:
		fmt.Println("Result of Multiplication: ", multiply(num1, num2))
	case 4:
		fmt.Println("Result of Division: ", divide(num1, num2))
	default:
		fmt.Println("Invalid Choice")
	}
}

// Function to calculate add
func add(num1, num2 float64) float64 {
	return num1 + num2
}

// Function to calculate subtract
func subtract(num1, num2 float64) float64 {
	return num1 - num2
}

// Function to calculate multiply
func multiply(num1, num2 float64) float64 {
	return num1 * num2
}

// Function to calculate divide
func divide(num1, num2 float64) float64 {
	return num1 / num2
}
