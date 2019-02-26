package main

import "fmt"

// &var = pointer of variable
// *var = value of pointer

func main() {

	x := 0
	// pass VALUE of x
	changeXVal(x)
	// it doesn't change in main()...
	fmt.Println("x = ", x)

	// pass MEMORY ADDRESS of x (so the real thing)
	// to pass a pointer, put "&" in front of the variable
	changeX(&x)
	fmt.Println("x = ", x)

	// we can declare a blank integer and assign it elsewhere
	y := new(int)	// this is JUST a pointer for an int
	changeYVal(y)	// no "&" since it's already a pointer
	fmt.Println(*y) // print value of y-pointer
}

// change actual x from main using its memory address
func changeX(x *int) {
	// the REAL memory address for x gets altered now
	*x = 2
}

// function that fails to change
func changeXVal(x int) {
	x = 2
}

// set y from outside without return
func changeYVal(y *int) {
	*y = 2
}
