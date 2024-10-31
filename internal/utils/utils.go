package utils

import (
	"net/http"
)

// UTILS
func ProcessError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
