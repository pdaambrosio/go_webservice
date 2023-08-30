package controllers

import (
	"encoding/json"
	"net/http"
)

// The function "RegisterControllers" registers a user controller to handle HTTP requests for the
// "/users" and "/users/" routes.
func RegisterControllers() {
	uc := newUserController()
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

// The function encodes the given data as JSON and writes it to the HTTP response.
func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
