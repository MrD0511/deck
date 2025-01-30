package handlers

import (
	"encoding/json"
	"net/http"
)

func PodsHandler(w http.ResponseWriter, r *http.Request) {

	pods := []string{"pod1", "pod2", "pod3"}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pods)
}