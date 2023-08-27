package controllers

import "net/http"

type usersController struct {}

func (uc usersController) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
