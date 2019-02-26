package main

import ("fmt"
	"strconv")

func main() {
	randInt := 5
	randString := "100"
	randString2 := "200.5"
	// cast int to float
	fmt.Println(float64(randInt))
	
	// int to string
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

        // flat to string
        newFloat, _ := strconv.ParseFloat(randString2, 64)
        fmt.Println(newFloat)


}
