package server

import (
	"io"
	"net/http"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	_, err := io.WriteString(w, `{"alive": true}`)
	if err != nil {
		return
	}
}
