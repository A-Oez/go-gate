package httperror

import (
	"encoding/json"
	"net/http"
)

type DefaultError struct {
	Status int
	Err    error
}

func (de DefaultError) WriteError(w http.ResponseWriter) {
	w.WriteHeader(de.Status)
	json.NewEncoder(w).Encode(de.Err)
}