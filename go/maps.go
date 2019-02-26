package main

import "fmt"

func main() {

	// maps are just key-value pairs, like python dicts
	// varName := make(map[keyType] valueType)
	presAge := make(map[string] int)
	presAge["roosevelt"] = 42
	fmt.Println(presAge["roosevelt"])

	// adding another one and checking how many pairs we have
	presAge["jfk"] = 43
	fmt.Println(len(presAge))

	// delete pair
	// delete(map, key)
	delete(presAge, "jfk")

}
