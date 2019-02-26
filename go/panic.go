package main

import "fmt"

func main() {

	// panic is a forced error we define (think, "raise")
	demPanic()

}

func demPanic() {

	defer func(){
		fmt.Println(recover())
	}()

	// panic(errormessage)
	panic("PANIC")

}
