package main

import "fmt"

func main() {

	// var <arrayName[size]> type
	var favNums2[5] float64

	favNums2[0] = 163
	favNums2[1] = 1234
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

	fmt.Println(favNums2)

	// faster way
	// arrayName := [size]type  {element, element, elemtn, etc...}
	favNums3 := [5]float64 {1, 2, 3, 5, 5}
	fmt.Println(favNums3)

	// iterate through array
	for i, value := range favNums3 {
		fmt.Println(value, i)
	}

}
