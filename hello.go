package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("Hello World!!!   aaaあああbbb")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	hello := []byte("<html><body><h1>こんにちは</h1></body></html>")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/kon", helloHandler)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
