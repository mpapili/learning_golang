package main

import "fmt"

func main() {

	var yourAge int = 18

	//switches are another way to do conditionals
	switch yourAge {
		case 16 : fmt.Println("Go drive")
		case 18 : fmt.Println("Go vote")
		default : fmt.Println("Go have fun")
	}
	
}
