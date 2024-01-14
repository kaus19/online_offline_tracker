package api

import (
    "net/http"
    "encoding/json"
)

type StatusParams struct {
	Username string
}

type StatusResponse struct {
	Code int
	Status bool
}

type Error struct {
	Code int
	Message string
}

func writeError(w http.ResponseWriter, message string, code int){
	resp := Error{
		Code: code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error){
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w, "An Unexpected error occured", http.StatusInternalServerError)
	}
)