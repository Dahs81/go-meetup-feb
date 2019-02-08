package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Dahs81/go-meetup-feb/11/pkg/person"
)

// PersonHander -
func PersonHander() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p person.Person

		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&p); err != nil {
			// respondWithError(w, http.StatusBadRequest, "Invalid request payload")
			fmt.Printf("Error: %v\n", err)
			return
		}
		defer r.Body.Close()

		// Not really necessary, but using to show that you can call a function from the package.
		newPerson := p.NewPerson(p.Name, p.Age)

		respondWithJSON(w, http.StatusOK, &newPerson)
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
