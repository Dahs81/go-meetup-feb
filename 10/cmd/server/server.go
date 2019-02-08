package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// MyServer is the default server for the analytics service
type MyServer struct {
	*http.Server
	Router *mux.Router
}

// New creates a new default losant server - This could be totally wrong
func New(r *mux.Router) *MyServer {

	s := &http.Server{
		Addr:           ":8085",
		Handler:        r, // This could be a mux!
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	return &MyServer{
		Server: s,
		Router: r,
	}
}
