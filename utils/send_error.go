package webAscii

import (
	"net/http"
)

// Printing HTTP related errors to the user
func SendError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
