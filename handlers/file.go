package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/events-app/events-api/internal/file"
	"github.com/events-app/events-api/internal/platform/web"
)

func UploadFile(uploadPath string, maxUploadSize int64) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			web.ErrorJSON(w, "file is too largre", http.StatusBadRequest)
			// renderError(w, "file is too largre", http.StatusBadRequest)
			return
		}

		// parse and validate file
		f, header, err := r.FormFile("file")
		if err != nil {
			web.ErrorJSON(w, "invalid file", http.StatusBadRequest)
			return
		}
		realName := header.Filename
		defer f.Close()
		file, err := file.Upload(f, uploadPath, realName)
		if err != nil {
			web.ErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		file.Path = fmt.Sprintf("https://%s/files/%s", r.Host, file.Path)
		// fl := file.New(fmt.Sprintf("https://%s/files/%s", r.Host, file.Name))
		// f := file.File{Path: fmt.Sprintf("https://%s/files/%s", r.Host, filename)}
		if err := json.NewEncoder(w).Encode(&file); err != nil {
			log.Printf("error: encoding response: %s", err)
		}
	})
}
func GetFiles(path string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files, err := file.ReadDir(path)
		if err != nil {
			web.ErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := json.NewEncoder(w).Encode(&files); err != nil {
			log.Printf("error: encoding response: %s", err)
		}
	})
}
