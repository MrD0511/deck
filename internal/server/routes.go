package server

import (
	"net/http"
	"github.com/MrD0511/deck/internal/server/handlers"
)

func RegisterRoutes(mux *http.ServeMux) {
 	mux.HandleFunc("/api/pods",handlers.PodsHandler)
}

