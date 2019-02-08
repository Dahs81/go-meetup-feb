package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Using the default ServeMux
*/

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple example")
}

type mySecondHandler int

func (h mySecondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the second handler output")
}

func main() {
	var h myHandler
	var hs mySecondHandler

	// Note: http has a default mux Handle function
	http.Handle("/first/", h)
	http.Handle("/second", hs)
	log.Fatal(http.ListenAndServe(":8085", nil)) // Passing nil uses the default mux
}
