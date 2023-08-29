package main

import (
	"net/http"
	"fmt"

	"github.com/pdaambrosio/go_webservice/controllers"
	// "github.com/pdaambrosio/go_webservice/models"
)

func main() {
	controllers.RegisterControllers()

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
