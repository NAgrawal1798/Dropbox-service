package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type FileMetadata struct {
	FileID    string    `json:"fileID"`
	FileName  string    `json:"fileName"`
	CreatedAt time.Time `json:"createdAt"`
	FileType  string    `json:"fileType"`
	FileSize  int64     `json:"fileSize"`
}

var (
	mu           sync.RWMutex
	fileMetadata = make(map[string]FileMetadata)
	uploadPath   = "./uploads/"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20) // 10 MB maximum file size
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileID := generateFileID() // Generate a unique file ID
	fileName := handler.Filename
	fileType := filepath.Ext(handler.Filename)
	fileSize := handler.Size
	createdAt := time.Now()

	// Create the upload directory if it doesn't exist
	err = os.MkdirAll(uploadPath, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new file in the uploads directory
	f, err := os.OpenFile(uploadPath+fileID+fileType, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Copy the file to the destination path
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Store file metadata in memory
	meta := FileMetadata{
		FileID:    fileID,
		FileName:  fileName,
		CreatedAt: createdAt,
		FileType:  fileType,
		FileSize:  fileSize,
	}

	mu.Lock()
	fileMetadata[fileID] = meta
	mu.Unlock()

	jsonResponse, err := json.Marshal(meta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func generateFileID() string {
	mu.Lock()
	defer mu.Unlock()
	return "file-" + time.Now().Format("20060102150405")
}
