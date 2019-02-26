package main

import "fmt"

func main() {

	listNums := []float64{1,2,3,4,5}
	// calling a function to print its return
	fmt.Println("Sum: ", addThemUp(listNums))

	// get multiple returns (like a tuple almost
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)

}

// fun funcName(argName type) returnType
func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}

	return sum
}

// here we return two ints, so the return is (int, int)
func next2Values(number int)  (int, int) {
	return number+1, number+2
}
