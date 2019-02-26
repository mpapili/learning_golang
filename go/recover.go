package main

import "fmt"

func main() {

	// do a divide that should crash (div by zero)
	fmt.Println(safeDivide(3,0))
	// make it to the safe divide since we recover
	fmt.Println(safeDivide(6,2))

}

// normal function for division
func safeDivide(num1 int, num2 int) int {

	// putting a recover() in a defer will stop it from failing
	defer func() {
		recover()
	}()

	solution := num1 / num2
	return solution

}
