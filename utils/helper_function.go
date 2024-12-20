package utils

import "fmt"

func HelperFunction() string {
	return "Helper function executed!"
}

func AddNumbers(num1 uint, num2 uint) uint {
	return num1 + num2
}

func DivideNumbers(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}
