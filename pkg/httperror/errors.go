package httperror

import (
	"encoding/json"
	"net/http"
)

type DefaultError struct {
	Status int		`json:"status"`			
	Msg    string 	`json:"error_message"`
}

func (err DefaultError) WriteError(w http.ResponseWriter) {
	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}