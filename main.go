package main

import (
	"fmt"
	"net/http"
	"time"

	"prueba.com/router"
)

func main() {

	// Initializations
	port := ":8080"
	readAndWriteTime := 10 * time.Second

	// Configuration

	server := http.Server{
		Addr:           port,
		ReadTimeout:    readAndWriteTime,
		WriteTimeout:   readAndWriteTime,
		MaxHeaderBytes: 1 << 20,
	}

	// Middlewares

	// Routing

	router.New()

	// Listening
	fmt.Printf("Server running on https://localhost%v", port)
	server.ListenAndServe()

}
