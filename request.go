package httputil

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/schema"
)

// ParseRequest parses HTTP requests into a struct. It supports POST form values
// and JSON encoded data as request.
// The request type is determined by the Content-Type header.
func ParseRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	switch r.Header.Get("Content-Type") {
	case "application/x-www-form-urlencoded":
		r.ParseForm()
		err := schema.NewDecoder().Decode(data, r.PostForm)
		if err != nil {
			WriteHttpStatus(w, http.StatusBadRequest)
			return err
		}
	case "application/json":
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			WriteHttpStatus(w, http.StatusBadRequest)
			return err
		}
	default:
		w.Header().Set("Accept", "application/x-www-form-urlencoded, application/json")
		WriteHttpStatus(w, http.StatusUnsupportedMediaType)
		return fmt.Errorf("Unsupported Media Type: %s", r.Header.Get("Content-Type"))
	}

	return nil
}
