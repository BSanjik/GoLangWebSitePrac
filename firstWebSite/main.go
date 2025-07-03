package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func handleRequest() {
	log.Println("Server is running ...")
	http.HandleFunc("/", homePage)
	http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}
