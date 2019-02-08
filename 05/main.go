package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
Using a custom ServeMux from the http package
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

	// Custom ServeMux
	mux := http.NewServeMux()
	mux.Handle("/first/", h)
	mux.Handle("/second", hs)
	log.Fatal(http.ListenAndServe(":8085", mux))
}
