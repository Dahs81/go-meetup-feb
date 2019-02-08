package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
HandleFunc
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

If you have used javascript's express framework or similar libraries, this
should be pretty familiar.
*/

func h(w http.ResponseWriter, r *http.Request) {
	// You could inspect r.Method and r.URL and do
	// 
	fmt.Fprintln(w, "This is a simple example")
}

func hs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the second handler output")
}

func main() {
	// Custom ServeMux
	mux := http.NewServeMux()
	mux.HandleFunc("/first/", h)
	mux.HandleFunc("/second", hs)
	log.Fatal(http.ListenAndServe(":8085", mux)) // Passing in mux
}
