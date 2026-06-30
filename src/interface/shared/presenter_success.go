package shared


import (
	"net/http"
	"encoding/json"
)

type ResponseData interface{}

type SuccessResponse[T ResponseData] struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data T `json:"data"`
}

func RespondWithOk[T ResponseData](rw http.ResponseWriter, message string, data T) {
	response := newSuccessResponse(http.StatusOK, message, data)
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(response)
}

func RespondWithCreated[T ResponseData](rw http.ResponseWriter, message string, data T) {
	response := newSuccessResponse(http.StatusCreated, message, data)
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(response)
}

func newSuccessResponse[T ResponseData](status int, message string, data T) *SuccessResponse[T] {
	return &SuccessResponse[T]{
		Status: status,
		Message: message,
		Data: data,
	}
}