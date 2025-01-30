package server

import (
	// "fmt"
	"log"
	"net/http"
)

func StartServer() error {
	mux := http.NewServeMux()

	RegisterRoutes(mux)

	server := &http.Server{
		Addr: ":8080",
		Handler: mux,
	}

	log.Println("Server is Listening at localhost:8080.")
	return server.ListenAndServe()
}
