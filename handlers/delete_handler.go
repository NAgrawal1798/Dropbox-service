package handlers

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileId"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := fileMetadata[fileID]; !exists {
		http.Error(w, fmt.Sprintf("File not found for fileID: %s", fileID), http.StatusNotFound)
		return
	}

	// Delete the file metadata
	delete(fileMetadata, fileID)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File with ID %s has been deleted", fileID)
}
