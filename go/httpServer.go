package main

import ("fmt"
	"net/http")

func main() {

	// bind root directory to handler
	http.HandleFunc("/", handler)
	// bind /earth directory to handler2
	http.HandleFunc("/earth", handler2)

	// listen for requests and serve responses
	http.ListenAndServe(":8080", nil)

	// this is insane - go to localhost:8080 and you'll see your response
	// go to localhost:8080/earth and you'll see your response from handler2

}

func handler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintf(w, "Hellooooo Earth\n")

}

func handler2(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Helllooo earth againnn")
}
