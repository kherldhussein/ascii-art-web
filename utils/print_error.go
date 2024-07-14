package webAscii

import (
	"net/http"
)

func SendError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
