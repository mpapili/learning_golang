package main

import ("fmt"
	"strings"
	"sort")
	//"os"
	//"log"
	//"io/ioutil"
	//"strconv")

func main() {

	sampString := "Hello World"
	// print if substring is in string
	fmt.Println(strings.Contains(sampString, "lo"))
	// return index that substring occurs in string
	fmt.Println(strings.Index(sampString, "lo"))
	// count all of the l's
	fmt.Println(strings.Count(sampString, "l"))
	// replace first 3 l's with x's
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	// turn comma-separated into csv
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	// sort a string from a list 
	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Letters: ", listOfLetters)

	// join array into single string
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")
	fmt.Println(listOfNums)


}
