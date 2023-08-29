package controllers

import "net/http"

// RegisterControllers is a function that registers the user controller with the http package so that it can be used to handle requests
func RegisterControllers() {
	// Create a new user controller instance and assign it to a variable
	uc := newUserController()

	// Register the user controller
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
