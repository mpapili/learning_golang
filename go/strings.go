package main

import "fmt"

func main() {
	var myName string = "mickey mouse"

	fmt.Println(myName)
	// print length of your string
	fmt.Println("my name has", len(myName), "characters")

	// concat strings together
	fmt.Println("i like\n" + myName)
}
