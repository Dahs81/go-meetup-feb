package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// gorilla mux example
// http://www.gorillatoolkit.org/pkg/mux

func main() {
	r := mux.NewRouter()

	// Note: there must be an ending / after /person
	r.HandleFunc("/person/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "person")
	})
	//.Methods("POST")

	log.Fatal(http.ListenAndServe(":8085", r))
}
