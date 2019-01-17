package handlers

import (
	"fmt"
	"net/http"

	"github.com/events-app/events-api/internal/file"
	"github.com/events-app/events-api/internal/platform/web"
)

func UploadFile(uploadPath string, maxUploadSize int64) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// validate file size
		r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
		if err := r.ParseMultipartForm(maxUploadSize); err != nil {
			web.RespondWithError(w, http.StatusBadRequest, "file is too largre")
			return
		}

		// parse and validate file
		f, header, err := r.FormFile("file")
		if err != nil {
			web.RespondWithError(w, http.StatusBadRequest, "invalid file")
			return
		}
		realName := header.Filename
		defer f.Close()
		file, err := file.Upload(f, uploadPath, realName)
		if err != nil {
			web.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		file.Path = fmt.Sprintf("%s://%s/files/%s", GetProtocol(r), r.Host, file.Path)
		web.RespondWithJSON(w, http.StatusOK, file)
	})
}
func GetFiles(path string) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		files, err := file.ReadDir(path)
		if err != nil {
			web.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		web.RespondWithJSON(w, http.StatusOK, files)
	})
}
