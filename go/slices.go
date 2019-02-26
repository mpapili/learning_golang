package main

import "fmt"

func main() {

	// slices are REFERENCE arrays
	// slices don't have a set size and can be appended to
	numSlice := []int {5, 4, 3, 2, 1}

	// make a new slice that's a subset of the old slice
	numSlice2 := numSlice[3:5]
	fmt.Println(numSlice, '\n', numSlice2)

	// print all the way up to index 3
	fmt.Println(numSlice[:2])

	// make a slice whose size isn't set
	// varName :=  make([]type, initialsize, maxSize)
	numSlice3 := make([]int, 5,  10)
	fmt.Println(numSlice3)

	// append to list
	// slice = append(slice, value, value, value....)
	numSlice3 = append(numSlice3, 0, -1)
	fmt.Println(numSlice3[6])

}
