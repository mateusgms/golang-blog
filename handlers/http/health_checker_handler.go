package http

import (
	"net/http"
)

// HealthHandler is a handler for health check
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}