package api

import "net/http"

func handleFiles(w http.ResponseWriter, r *http.Request) {
	Success(w, "Files handled")
}
