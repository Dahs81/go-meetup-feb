package main

import (
	"fmt"
	"log"
	"net/http"
)

// Person is custom function that returns HandleFunc
type Person struct {
	Name string
	Age  int
	// MongoURL string
}

// Create - Returns a Handler that can be passed to http.Handle
func (p Person) Create(name string, age int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "This prints the name and age of a Person:\n\t%v\n\t%v\n", name, age)
	})
}

func main() {
	bill := Person{
		Name: "Bill",
		Age:  45,
	}
	h := bill.Create(bill.Name, bill.Age)
	mux := http.NewServeMux()
	mux.Handle("/person/", h)
	log.Fatal(http.ListenAndServe(":8085", mux)) // Passing in mux
}
