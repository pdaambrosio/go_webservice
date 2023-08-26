package main

import (
	"github.com/pdaambrosio/go_webservice/models"
	"fmt"
)

func main() {
	u := models.User{
		ID:			2,
		FirstName:	"John",
		LastName:	"Snow",
	}
	fmt.Println(u)
}
