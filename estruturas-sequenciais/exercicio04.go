package main

import "fmt"

func main() {

	var grade1, grade2, grade3, grade4 float32
	println("Inform four grades: ")

	_, _ = fmt.Scan(&grade1, &grade2, &grade3, &grade4)

	println(fmt.Sprintf("The average grade is: %.2f", (grade1+grade2+grade3+grade4)/4))

}
