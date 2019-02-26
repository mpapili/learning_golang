package main

import ("fmt"
	"time")

// function that will count to ten over ten seconds
	// while printint out its thread id
func count(id int){

	for i:=0; i < 10; i++{
		fmt.Println(id, ":", i)

		time.Sleep(time.Millisecond * 1000) // one second pause
	}

}

func main() {

	fire off ten goroutines counting at the SAME TIME
	for i := 0; i < 10; i++ {
		go count(i)

	}
	// wait long-enough for them to finish
	time.Sleep(time.Millisecond * 11000)

}
