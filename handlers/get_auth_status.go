package handler

import "net/http"

func handleGetAuthStatus(w http.ResponseWriter, req *http.Request) {
	SetupCORSResponse(&w, req)
	w.WriteHeader(http.StatusOK)
}
