package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileId"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := fileMetadata[fileID]; !exists {
		http.Error(w, fmt.Sprintf("File not found for fileID: %s", fileID), http.StatusNotFound)
		return
	}

	var updatedMeta FileMetadata
	err := json.NewDecoder(r.Body).Decode(&updatedMeta)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Update file name
	if updatedMeta.FileName != "" {
		fileMetadata[fileID] = FileMetadata{
			FileID:    fileMetadata[fileID].FileID,
			FileName:  updatedMeta.FileName,
			CreatedAt: fileMetadata[fileID].CreatedAt,
			FileType:  fileMetadata[fileID].FileType,
			FileSize:  fileMetadata[fileID].FileSize,
		}
	}

	jsonResponse, err := json.Marshal(fileMetadata[fileID])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

