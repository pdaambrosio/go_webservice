package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/pdaambrosio/go_webservice/models"
)

// The userController struct contains a regular expression pattern for matching user IDs.
// @property userIDPattern - The `userIDPattern` property is a regular expression pattern that is used
// to validate user IDs. It is a pointer to a `regexp.Regexp` object, which allows you to perform
// pattern matching and validation on user IDs.
type userController struct {
	userIDPattern *regexp.Regexp
}

// The `ServeHTTP` function is the main handler function for the `userController` struct. It implements
// the `http.Handler` interface, allowing the struct to handle HTTP requests.
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/users" {
		switch r.Method {
		case http.MethodGet:
			uc.getAll(w, r)
		case http.MethodPost:
			uc.post(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		matches := uc.userIDPattern.FindStringSubmatch(r.URL.Path)
		if len(matches) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		id, err := strconv.Atoi(matches[1])
		if err != nil {
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
			w.WriteHeader(http.StatusNotImplemented)
		}
	}
}

// The `getAll` function is a method of the `userController` struct. It is responsible for handling
// HTTP GET requests to retrieve all users.
func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	encodeResponseAsJSON(models.GetUsers(), w)
}

// The `get` function is a method of the `userController` struct. It is responsible for handling HTTP
// GET requests to retrieve a specific user by their ID.
func (uc userController) get(id int, w http.ResponseWriter, r *http.Request) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	encodeResponseAsJSON(u, w)
}

// The `post` function is a method of the `userController` struct. It is responsible for handling HTTP
// POST requests to create a new user.
func (uc *userController) post(w http.ResponseWriter, r *http.Request) {
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

// The `put` function is a method of the `userController` struct. It is responsible for handling HTTP
// PUT requests to update an existing user.
func (uc *userController) put(id int, w http.ResponseWriter, r *http.Request) {
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

// The `delete` function is a method of the `userController` struct. It is responsible for handling
// HTTP DELETE requests to remove an existing user.
func (uc *userController) delete(id int, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	err := models.RemoveUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

// The `parseRequest` function is a method of the `userController` struct. It is responsible for
// parsing the request body of an HTTP request and decoding it into a `models.User` object.
func (uc *userController) parseRequest(r *http.Request) (models.User, error) {
	dec := json.NewDecoder(r.Body)
	var u models.User
	err := dec.Decode(&u)
	if err != nil {
		return models.User{}, err
	}
	return u, nil
}

// The function `newUserController` returns a new instance of the `userController` struct.
func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
