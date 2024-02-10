package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	hello := []byte("<html><body><h1>情報ビジネス</h1></body></html>")
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
	
}

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	hello := []byte(`aaa`)
	_, err := w.Write(hello)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/kon", helloHandler2)
	fmt.Println("Server Start Up........")
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
