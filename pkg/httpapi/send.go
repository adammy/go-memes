package httpapi

import (
	"encoding/json"
	"net/http"
	"time"
)

// SendJSON will set the appropriate headers and write to the response writer for some data.
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set(ContentTypeHeader, ApplicationJson)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}

	return nil
}

// SendErrorJSON will set the appropriate headers and a structured Error to the response writer.
func SendErrorJSON(w http.ResponseWriter, statusCode int, error error) error {
	w.Header().Set(ContentTypeHeader, ApplicationJson)
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(Error{
		Error:     error.Error(),
		Timestamp: time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
