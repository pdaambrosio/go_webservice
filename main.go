package main

import (
	"net/http"
	"fmt"

	"github.com/pdaambrosio/go_webservice/controllers"
	// "github.com/pdaambrosio/go_webservice/models"
)

func main() {
	// Register the user controller with the http package so that it can handle requests
	controllers.RegisterControllers()

	// Create a server that listens on port 3000 and passes all requests to the http package
	err := http.ListenAndServe(":3000", nil)
	// If there is an error, print it to the console and exit the program
	// It's possible to use the log package to log the error to a file instead of printing it to the console, example: log.Fatal(http.ListenAndServe(":3000", nil))
	if err != nil {
		fmt.Println(err.Error())
	}
}
