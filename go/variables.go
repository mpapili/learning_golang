package main

import "fmt"

func main() {

	// Go is strongly-typed, we set it with "var <name> <type>
	var age int = 40
	var favNum float64 = 1.6180339

	// but the option exists to have Go figure out what type it wants
	randNum := 1

	fmt.Println(age, favNum, randNum)

	var numOne = 1.00
	var num99 = 0.999

	// do some basic artihmatic 
	fmt.Println(numOne - num99)

}
