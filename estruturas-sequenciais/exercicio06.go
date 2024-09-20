package main

import (
	"fmt"
)

func main() {

	var radio float64
	println("Insert the radio of a circle: ")

	_, _ = fmt.Scan(&radio)
	println(fmt.Sprintf("The area of the circle is %.2f", 3.14*(radio*radio))) // :D

}
