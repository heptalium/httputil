package httputil

import (
	"fmt"
	"net/http"
)

// WriteHttpStatus replies to the HTTP request with the status code and
// a message consisting of the status code and the corresponding message.
func WriteHttpStatus(w http.ResponseWriter, status int) {
	message := fmt.Sprintf("%d %s", status, http.StatusText(status))
	http.Error(w, message, status)
}
