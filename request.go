package httputil

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"mime"
	"net/http"

	"github.com/gorilla/schema"
)

// ParseRequest parses HTTP requests into a struct. It supports POST form
// values, JSON encoded data and XML encoded data as request.
// The request type is determined by the Content-Type header.
func ParseRequest(w http.ResponseWriter, r *http.Request, data interface{}) error {
	mediaType, _, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if err != nil {
		WriteHttpStatus(w, http.StatusBadRequest)
		return err
	}

	switch mediaType {
	case "application/x-www-form-urlencoded":
		r.ParseForm()
		err := schema.NewDecoder().Decode(data, r.PostForm)
		if err != nil {
			WriteHttpStatus(w, http.StatusBadRequest)
			return err
		}
	case "multipart/form-data":
		r.ParseMultipartForm(64 * 1024 * 1024)
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
	case "application/xml", "text/xml":
		err := xml.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			WriteHttpStatus(w, http.StatusBadRequest)
			return err
		}
	default:
		w.Header().Set("Accept", "application/x-www-form-urlencoded, multipart/form-data, application/json, application/xml, text/xml")
		WriteHttpStatus(w, http.StatusUnsupportedMediaType)
		return fmt.Errorf("Unsupported Media Type: %s", mediaType)
	}

	return nil
}
