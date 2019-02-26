package main

import ("fmt"
	"log"
	"io/ioutil"
	"os") // os package

func main() {

	// use os.Create to create a file
	file, err := os.Create("samp.txt")

	// if error occurred, log it
	if err != nil {
		log.Fatal(err)
	}
	// write to the file and close it
	file.WriteString("This is text")
	file.Close()

	// start reading the file
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	// set string to the stream and print it out
	readString := string(stream)
	fmt.Println(readString)

}
