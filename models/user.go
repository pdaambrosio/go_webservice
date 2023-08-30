package models

import (
	"errors"
	"fmt"
)

// The User type represents a user with an ID, first name, and last name.
// @property {int} ID - An integer representing the user's ID.
// @property {string} FirstName - The FirstName property is a string that represents the first name of
// a user.
// @property {string} LastName - The LastName property is a string that represents the last name of a
// user.
type User struct {
	ID        int
	FirstName string
	LastName  string
}

// The `var` statement is used to declare variables. In this case, it is declaring two variables:
// `users` and `nextID`.
var (
	users  []*User // slice of pointers to User
	nextID = 1
)

// The GetUsers function returns a slice of User pointers.
func GetUsers() []*User {
	return users
}

// The AddUser function adds a new user to a collection of users, ensuring that the user's ID is not
// already set.
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("new user must not include ID or it must be set to zero")
	}

	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// The function `GetUserByID` takes an ID as input and returns the corresponding user if found, or an
// error if not found.
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil // dereference pointer to User
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// The function UpdateUser updates a user in a slice of users based on their ID.
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u // update slice element
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// The function `RemoveUserByID` removes a user from a slice of users based on their ID.
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...) // remove slice element
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
