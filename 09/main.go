package main

import (
	"fmt"
	"log"
	"net/http"
)

// Middleware example

func main() {
	var h myHandler
	var hs mySecondHandler

	mux := http.NewServeMux()
	mux.Handle("/first/", simpleMw(h))
	mux.Handle("/second", hs)
	log.Fatal(http.ListenAndServe(":8085", mux))
}

// simple logging middleware - Takes a Handler, does some work, and returns a Handler
func simpleMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("\nRequestURI: %v\n\n", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

type myHandler int

func (h myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple example")
}

type mySecondHandler int

func (h mySecondHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the second handler output")
}
