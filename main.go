package main

import (
	"fmt"
	"net/http"

	"prueba.com/router"
)

func main() {

	// Initializations
	port := ":8080"

	// Configuration

	// Middlewares

	// Routing
	serverMux := http.NewServeMux()

	serverRouter := router.New(serverMux)

	// Listening
	fmt.Printf("Server running on http://localhost%v", port)
	http.ListenAndServe(port, serverRouter.Mux)

}
