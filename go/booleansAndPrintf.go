package main

import "fmt"

func main() {

	// booleans are just "bool"
	// booleans are lowercase true/false
	var isOver40 bool = true

	// print out a float
	var pi float64 = 3.14159
	fmt.Printf("%f \n", pi)
	// print out first three decimal places of float
	fmt.Printf("%.3f \n", pi)

	// print TYPE of variable
	fmt.Printf("%T\n", pi)

	//print boolean
	fmt.Printf("%t\n", isOver40)

	// Printf %* codes:
	// %d for integer
	// %b for binary
	// %c for character code
	// %x for hexcodes
	// %e for scientific notations

}
