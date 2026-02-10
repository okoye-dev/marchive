package api

import "net/http"

func handleHealth(w http.ResponseWriter, r *http.Request) {
	Success(w, map[string]string{"status": "ok"})
}

func handleReady(w http.ResponseWriter, r *http.Request) {
	Success(w, map[string]string{"status": "ready"})
}
