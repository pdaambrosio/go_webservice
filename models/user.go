package models

type User struct {
	ID			int
	FirstName	string
	LastName	string
}

var (
	users []*User // slice of pointers to User
	nextID = 1
)
