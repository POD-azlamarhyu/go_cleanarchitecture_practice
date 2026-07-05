package shared

import (
	"net/http"
	"encoding/json"
)

type ErrorResponse struct {
	Status int `json:"status"`
	Message string `json:"message"`
}

func RespondWithError(rw http.ResponseWriter, status int, message string) {
	response := newErrorResponse(status, message)
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(response)
}

func RespondWithBadRequest(rw http.ResponseWriter, message string) {
	RespondWithError(rw, http.StatusBadRequest, message)
}

func RespondWithInternalServerError(rw http.ResponseWriter, message string) {
	RespondWithError(rw, http.StatusInternalServerError, message)
}

func RespondWithNotFound(rw http.ResponseWriter, message string) {
	RespondWithError(rw, http.StatusNotFound, message)
}

func RespondWithUnauthorized(rw http.ResponseWriter, message string) {
	RespondWithError(rw, http.StatusUnauthorized, message)
}

func RespondWithForbidden(rw http.ResponseWriter, message string) {
	RespondWithError(rw, http.StatusForbidden, message)
}

func newErrorResponse(status int, message string) *ErrorResponse {
	return &ErrorResponse{
		Status: status,
		Message: message,
	}
}