package main

import "fmt"

func main() {

	fmt.Println(subtractThem(1,2,3,4,5))

}

// function that accept any number of args
// func funcName(argName ...type) returnType {}
func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue
}
