package main

import (
	"io"
	"log"
	"net/http"
)

type myhandler int

/*
Of course, this is a terrible idea and there are much better ways to
do routing in Go. Let's move on!
*/

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Simple Routing
	switch r.URL.Path {
	case "/":
		io.WriteString(w, "Home Page")
	case "/apple":
		io.WriteString(w, "This is the Apple path!")
	case "/pear":
		io.WriteString(w, "This is the Pear path!")
	default:
		io.WriteString(w, "This is a catch-all for any other path.")
	}
}

func main() {
	var h myhandler
	log.Fatal(http.ListenAndServe(":8085", h))
}
