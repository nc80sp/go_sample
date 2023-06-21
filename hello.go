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
	hello := []byte('
	<!doctype html>
	<html dir="ltr" lang="ja">
	  <head>
		<meta charset="utf-8">
		<title>新しいタブ</title>
		<style>
		  body {
			background: #FFFFFF;
			margin: 0;
		  }
	
		  #backgroundImage {
			border: none;
			height: 100%;
			pointer-events: none;
			position: fixed;
			top: 0;
			visibility: hidden;
			width: 100%;
		  }
	
		  [show-background-image] #backgroundImage {
			visibility: visible;
		  }
		</style>
	  </head>
	  <body>
		<iframe id="backgroundImage" src=""></iframe>
		<ntp-app></ntp-app>
		<script type="module" src="new_tab_page.js"></script>
		<link rel="stylesheet" href="chrome://resources/css/text_defaults_md.css">
		<link rel="stylesheet" href="chrome://theme/colors.css?sets=ui,chrome">
		<link rel="stylesheet" href="shared_vars.css">
	  </body>
	</html>
	')
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
