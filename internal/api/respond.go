package api

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func OK(w http.ResponseWriter) {
	JSON(w, http.StatusOK, "ok")
}

func Success(w http.ResponseWriter, msg interface{}) {
	JSON(w, http.StatusOK, msg)
}