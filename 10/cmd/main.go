package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Dahs81/go-meetup-feb/10/cmd/server"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	s := server.New(r)
	s.Router.HandleFunc("/person/", personHandler)

	log.Fatal(s.Server.ListenAndServe())
}

func personHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Person - Custom Server")
}
