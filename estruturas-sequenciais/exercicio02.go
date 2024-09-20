package main

import "fmt"

func main() {

	println("Inform a number:")
	var number int
	_, _ = fmt.Scan(&number)

	println("The input number is: ", number)
}
