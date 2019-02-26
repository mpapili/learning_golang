package main

import "fmt"

func main() {

	// create something to iterate over
	i := 1

	for i <= 5 {
		fmt.Println(i)
		i++	// i = i + 1
	}

	fmt.Println("\n\n")

	// or all at once, similar to c++
	for j := 0; j < 5; j++ {
		fmt.Println(j)
	}

}
