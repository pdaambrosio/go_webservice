package controllers

import (
	"encoding/json"
	"net/http"
)

// RegisterControllers is a function that registers the user controller with the http package so that it can be used to handle requests
func RegisterControllers() {
	// Create a new user controller instance and assign it to a variable
	uc := newUserController()

	// Register the user controller
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w http.ResponseWriter) {
	// Set the content type header on the http.ResponseWriter object
	w.Header().Set("Content-Type", "application/json")
	// Use the json package to encode the data parameter and write it to the http.ResponseWriter object
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
