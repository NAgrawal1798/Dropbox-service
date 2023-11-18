package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/gorilla/mux"
)

func ReadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fileID := vars["fileId"]

	filePath := uploadPath + fileID

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("File not found for fileID: %s", fileID), http.StatusNotFound)
		return
	}
	defer file.Close()

	fileExt := filepath.Ext(filePath)
	contentType := getContentType(fileExt)
	w.Header().Set("Content-Type", contentType)

	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Failed to read the file", http.StatusInternalServerError)
		return
	}
}

func getContentType(fileExt string) string {
	switch strings.ToLower(fileExt) {
	case ".pdf":
		return "application/pdf"
	case ".txt":
		return "text/plain"
	default:
		return "application/octet-stream"
	}
}
