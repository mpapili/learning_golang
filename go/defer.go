package main

import "fmt"

func main() {

	// defer <method to defer>
	defer printTwo()
	printOne()

	// defer does not run until it's enclosing function (here "main()" finishes)
	// defer lets you use one last function as a cleanup function without having to struggle to make sure that it's going at the end
	// maybe "closing a file" ?
}

// define two tiny one-liner functions
func printOne() { fmt.Println(1)}
func printTwo() { fmt.Println(2)}
