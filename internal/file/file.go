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

type File struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// Upload handles uploading files like png, jpg, etc to server
func Upload(file io.Reader, path string, realName string) (*File, error) {
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.New("invalid file")
	}

	// check file type, detectcontenttype only needs the first 512 bytes
	filetype := http.DetectContentType(fileBytes)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		return nil, errors.New("invalid file type")
	}
	fileName := randToken(12)
	fileEndings, err := mime.ExtensionsByType(filetype)
	if err != nil {
		return nil, errors.New("cannot read file type")
	}
	newPath := filepath.Join(path, fileName+fileEndings[0])

	// write file
	newFile, err := os.Create(newPath)
	if err != nil {
		return nil, errors.New("cannot write file")
	}
	defer newFile.Close() // idempotent, okay to call twice
	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		return nil, errors.New("canot write file")
	}
	return &File{realName, fileName + fileEndings[0]}, nil
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

// ReadDir returns list of files in a given path
func ReadDir(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed opening directory: %s", err)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	return list, nil
}
