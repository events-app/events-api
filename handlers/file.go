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
		f, _, err := r.FormFile("file")
		if err != nil {
			web.ErrorJSON(w, "invalid file", http.StatusBadRequest)
			return
		}
		defer f.Close()
		filename, err := file.Upload(f, uploadPath)
		if err != nil {
			web.ErrorJSON(w, err.Error(), http.StatusBadRequest)
			return
		}
		fl := file.New(fmt.Sprintf("https://%s/files/%s", r.Host, filename))
		// f := file.File{Path: fmt.Sprintf("https://%s/files/%s", r.Host, filename)}
		if err := json.NewEncoder(w).Encode(&fl); err != nil {
			log.Printf("error: encoding response: %s", err)
		}
	})
}
