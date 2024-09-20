package main

import "fmt"

func main() {

	var number1, number2 int

	println("Inform the first number: ")
	_, _ = fmt.Scan(&number1)
	println("Inform the first number: ")
	_, _ = fmt.Scan(&number2)

	println("The sum is: ", number1+number2)
}
