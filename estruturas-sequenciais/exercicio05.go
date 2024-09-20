package main

import "fmt"

func main() {
	var meters float32
	println("Insert the distance in meters: ")
	_, _ = fmt.Scan(&meters)

	println(fmt.Sprintf("The distance in centimeters is: %.2f", meters*100))

}
