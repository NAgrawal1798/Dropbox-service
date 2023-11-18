package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"dropbox-service/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/files/upload", handlers.UploadFile).Methods("POST")
	r.HandleFunc("/files/{fileId}", handlers.ReadFile).Methods("GET")
	r.HandleFunc("/files/{fileId}", handlers.UpdateFile).Methods("PUT")
	r.HandleFunc("/files/{fileId}", handlers.DeleteFile).Methods("DELETE")
	r.HandleFunc("/files", handlers.ListFiles).Methods("GET")

	log.Println("Server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
