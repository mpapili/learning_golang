package main

import "fmt"

func main() {

	fmt.Println(factorial(6))

}

// recursive function set up similarly
func factorial(num int) int{
	// define our "out"
	if num == 0 {
		return 1
	}
	return num * factorial(num - 1)	// keep drilling down until one finally returns 1 for the multiplier

}
