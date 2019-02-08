package main

import (
	"fmt"
	"net/http"
)

/*
Handler Interface
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/

/*
Any type that implements the Handler interface is a handler and can
be passed to ListenAndServer. We're using an int to prove that it
doesn't matter what the type is.
*/
type myhandler int

func (h myhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a simple example")
}

func main() {
	var h myhandler
	http.ListenAndServe(":8085", h)
}
