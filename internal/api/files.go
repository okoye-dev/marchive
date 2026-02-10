package api

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

const maxUploadSize = 32 << 20

func (h *Handlers) handleFilesList(w http.ResponseWriter, r *http.Request) {
	prefix := r.URL.Query().Get("prefix")
	bucket := h.bucketFromReq(r)
	keys, err := h.Files.ListFiles(r.Context(), bucket, prefix)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	Success(w, map[string]interface{}{
		"bucket": bucket,
		"prefix": prefix,
		"keys":   keys,
	})
}

func (h *Handlers) handleFilesUpload(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		Error(w, http.StatusBadRequest, "invalid multipart form")
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		Error(w, http.StatusBadRequest, "file is required")
		return
	}
	defer file.Close()

	key := r.FormValue("key")
	if key == "" {
		key = header.Filename
	}
	if key == "" {
		Error(w, http.StatusBadRequest, "key is required")
		return
	}

	bucket := h.bucketFromReq(r)
	if err := h.Files.UploadFile(r.Context(), bucket, key, file); err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	Success(w, map[string]interface{}{
		"bucket":   bucket,
		"key":      key,
		"fileName": header.Filename,
		"size":     header.Size,
	})
}

func (h *Handlers) handleFilesDownload(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		Error(w, http.StatusBadRequest, "key is required")
		return
	}
	bucket := h.bucketFromReq(r)
	rc, err := h.Files.DownloadFile(r.Context(), bucket, key)
	if err != nil {
		if os.IsNotExist(err) {
			Error(w, http.StatusNotFound, "file not found")
			return
		}
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rc.Close()

	head := make([]byte, 512)
	n, _ := rc.Read(head)
	contentType := http.DetectContentType(head[:n])
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filepath.Base(key)+"\"")
	w.WriteHeader(http.StatusOK)

	reader := io.MultiReader(bytes.NewReader(head[:n]), rc)
	_, _ = io.Copy(w, reader)
}

func (h *Handlers) handleFilesDelete(w http.ResponseWriter, r *http.Request) {
	key := chi.URLParam(r, "key")
	if key == "" {
		Error(w, http.StatusBadRequest, "key is required")
		return
	}
	bucket := h.bucketFromReq(r)
	if err := h.Files.DeleteFile(r.Context(), bucket, key); err != nil {
		if os.IsNotExist(err) {
			Error(w, http.StatusNotFound, "file not found")
			return
		}
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}
	Success(w, map[string]string{"status": "deleted"})
}

func (h *Handlers) bucketFromReq(r *http.Request) string {
	bucket := r.URL.Query().Get("bucket")
	if bucket == "" {
		bucket = h.DefaultBucket
	}
	return bucket
}
