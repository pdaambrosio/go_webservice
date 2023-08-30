package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pdaambrosio/go_webservice/models"
)

// userController is a struct that contains a pointer to a regular expression object that will be used to match the URL path
type userController struct {
	userIDPattern *regexp.Regexp
}

// ServerHTTP is a method that implements the Handler interface and will be called automatically by the http package when a request is received
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			// If the request method is not supported, write a 405 status code to the http.ResponseWriter object
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		// If the matches slice is not empty, then the URL path matched the regular expression
		if len(matches) == 0 {
			// If the URL path did not match the regular expression, write a 404 status code to the http.ResponseWriter object
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// If the URL path matched the regular expression, then the first element in the matches slice will be the entire URL path
		// and the second element will be the first subexpression match, which is the user ID
		id, err := strconv.Atoi(matches[1])
		if err != nil {
			// If the user ID cannot be converted to an integer, write a 404 status code to the http.ResponseWriter object
			w.WriteHeader(http.StatusNotFound)
		}

		switch r.Method {
		case http.MethodGet:
			uc.get(id, w, r)
		case http.MethodPut:
			uc.put(id, w, r)
		case http.MethodDelete:
			uc.delete(id, w)
		default:
			// If the request method is not supported, write a 405 status code to the http.ResponseWriter object
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	// Write the response to the http.ResponseWriter object
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJSON(models.GetUsers(), w)
}

func (uc userController) get(id int, w http.ResponseWriter, r *http.Request) {
	// Write the response to the http.ResponseWriter object
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
	// Parse the request body and create a new user object
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	u, err = models.AddUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
	// Parse the request body and create a new user object
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not parse User object"))
		return
	}
	if id != u.ID {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID of submitted user must match ID in URL"))
		return
	}
	u, err = models.UpdateUser(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	encodeResponseAsJSON(u, w)
}

func (uc *userController) delete(id int, w http.ResponseWriter) {
	// Delete the user with the specified ID
	w.WriteHeader(http.StatusOK)
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

func newUserController() *userController {
	// Create a new user controller and initialize the userIDPattern field with a regular expression that will match the URL path
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
