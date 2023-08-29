package main

import (
	"net/http"

	"github.com/pdaambrosio/go_webservice/controllers"
	// "github.com/pdaambrosio/go_webservice/models"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
