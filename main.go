package main

import (
	"net/http"
	"fmt"

	"github.com/pdaambrosio/go_webservice/controllers"
)

// The main function registers controllers and starts a server on port 3000.
func main() {
	controllers.RegisterControllers()

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
