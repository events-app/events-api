package file

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

// Upload handles uploading files like png, jpg, etc to server
func Upload(file io.Reader, path string) (string, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		// renderError(w, "INVALID_FILE", http.StatusBadRequest)
		return "", errors.New("invalid file")
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		// renderError(w, "INVALID_FILE_TYPE", http.StatusBadRequest)
		return "", errors.New("invalid file type")
	}
	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(filetype)
	if err != nil {
		// renderError(w, "CANT_READ_FILE_TYPE", http.StatusInternalServerError)
		return "", errors.New("cannot read file type")
	}
	newPath := filepath.Join(path, fileName+fileEndings[0])
	// fmt.Printf("FileType: %s, File: %s\n", fileType, newPath)

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		// renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return "", errors.New("cannot write file")
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		// renderError(w, "CANT_WRITE_FILE", http.StatusInternalServerError)
		return "", errors.New("canot write file")
	}
	return fileName + fileEndings[0], nil
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}
