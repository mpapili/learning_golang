package main

import "fmt"

func main() {

	yourAge := 18

	// conditionals are very traditional here
	if yourAge >= 16 {
		fmt.Println("You can drive")
	} else if yourAge >= 18 {
		fmt.Println("You can drive and vote")
	} else {
		fmt.Println("You can't drive")
	}

}
