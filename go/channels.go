package main

import ("fmt"
	"time"
	"strconv")

var pizzaNum = 0
var pizzaName = ""

func makeDough(stringChan chan string){
	pizzaNum++

	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Make Dough and Send for Sauce")

	stringChan <- pizzaName

	time.Sleep(time.Millisecond * 10) // slight sleep

}

func addSauce(stringChan chan string){
	pizza := <- stringChan

        fmt.Println("Add Sauce and Send ", pizza, "for topping")

        time.Sleep(time.Millisecond * 10) // slight sleep

}

func addToppings(stringChan chan string){
	pizza := <- stringChan

        fmt.Println("Toppings are added to ", pizza)

        time.Sleep(time.Millisecond * 10) // slight sleep

}


func main() {
 
	stringChan := make(chan string)

	// sometimes skips toppings..??
	for i:= 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)
		time.Sleep(time.Millisecond * 5000)
	}

}
