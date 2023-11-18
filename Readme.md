# Dropbox Service
The Dropbox Service project is a Go-based application that provides functionality for file uploads and management.


## Dependencies
Gorilla Mux: This project uses the Gorilla Mux router for handling HTTP requests.

# Clone the Repository:

```bash
git clone https://github.com/NAgrawal/dropbox-service.git
```

# Install Go:

If you haven't installed Go, follow the steps below:

Mac/Linux: You can use Homebrew: brew install go

Install Dependencies: 
```bash
cd dropbox-service
go mod tidy
```

Start the server:
```bash
go run main.go
```

Features
Endpoints


###Upload File
Endpoint: POST /files/upload
Description: Uploads a file onto the platform.
Input: File binary data, file name, metadata (if any).
Output: A unique file identifier.
Metadata Saved: File name, createdAt timestamp, size, file type.


### Read File
Endpoint: GET /files/{fileId}
Description: Retrieves a specific file based on a unique identifier.
Input: Unique file identifier.
Output: File binary data.


### Update File
Endpoint: PUT /files/{fileId}
Description: Updates an existing file or its metadata.
Input: New file binary data or new metadata.
Output: Updated metadata or a success message.

###Delete File
Endpoint: DELETE /files/{fileId}
Description: Deletes a specific file based on a unique identifier.
Input: Unique file identifier.
Output: A success or failure message.


###List Files
Endpoint: GET /files
Description: Lists all available files and their metadata.
Input: None.
Output: A list of file metadata objects, including file IDs, names, createdAt timestamps, etc.



Technologies Used
Backend Language & Framework: Go with Gorilla Mux
Database: In-memory storage for file metadata
Storage: Local storage (File system) for simplicity


Usage

Uploading a File:
```bash
curl -X POST -F "file=@/path/to/your/file" http://localhost:8000/files/upload
```
Reading a File:
```bash
curl http://localhost:8000/files/{fileId}
```
Updating a File:
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"FileName": "Updated_File_Name"}' http://localhost:8000/files/{fileId}
```
Deleting a File:
```bash
curl -X DELETE http://localhost:8000/files/{fileId}
```

Listing Files:
```bash
curl http://localhost:8000/files
```
Notes
The server runs on port 8000 by default. Adjust the port if needed.
Replace {fileId} with the actual unique file identifier.
