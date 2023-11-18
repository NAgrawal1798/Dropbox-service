package handlers

import (
	"encoding/json"
	"net/http"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	var files []FileMetadata

	mu.RLock()
	for _, file := range fileMetadata {
		files = append(files, file)
	}
	mu.RUnlock()

	jsonResponse, err := json.Marshal(files)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
