package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/",hello)
	http.HandleFunc("/hi",hi)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func hi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi")
}