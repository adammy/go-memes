package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/adammy/memepen-services/pkg/pointer"
	"github.com/go-chi/chi/v5/middleware"
)

// SendJSON will set the appropriate headers and write the data argument to the response writer.
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set(ContentTypeHeader, ApplicationJson)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(data)
}

// SendErrorJSON will set the appropriate headers and write a structured Error to the response writer.
func SendErrorJSON(w http.ResponseWriter, statusCode int, error error) {
	w.Header().Set(ContentTypeHeader, ApplicationJson)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(Error{
		Error:     error.Error(),
		Timestamp: time.Now(),
	})
}

// SendErrorJSONWithRequest will set the appropriate headers and write a structured Error with a request ID to the response writer.
func SendErrorJSONWithRequest(w http.ResponseWriter, r *http.Request, statusCode int, error error) {
	if r == nil {
		SendErrorJSON(w, statusCode, error)
		return
	}

	w.Header().Set(ContentTypeHeader, ApplicationJson)
	w.WriteHeader(statusCode)

	_ = json.NewEncoder(w).Encode(Error{
		Error:     error.Error(),
		Timestamp: time.Now(),
		RequestID: pointer.GetStringP(middleware.GetReqID(r.Context())),
	})
}
