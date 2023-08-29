package controllers

import (
	"net/http"
	"regexp"
)

// userController is a struct that contains a pointer to a regular expression object that will be used to match the URL path
type userController struct {
	userIDPattern *regexp.Regexp
}

// ServerHTTP is a method that implements the Handler interface and will be called automatically by the http package when a request is received
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Write the response to the http.ResponseWriter object
	w.Write([]byte("Hello from the User Controller!"))
}

// newUserController is a function that creates a new user controller and returns a pointer to it so that it can be used to handle requests
func newUserController() *userController {
	// Create a new user controller and assign it to a variable and return it as a pointer to the userController struct
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
