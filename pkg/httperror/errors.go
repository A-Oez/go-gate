package httperror

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type DefaultError struct {
	Status int		`json:"status"`			
	Msg    string 	`json:"error_message"`
	Error  error	`json:"-"` //is used for internal error logging
}

func (err DefaultError) WriteError(w http.ResponseWriter) {
	if err.Error != nil {
		log.Println(err.Error)
	} else {
		log.Println(err.Msg)
	}

	w.WriteHeader(err.Status)
	json.NewEncoder(w).Encode(err)
}

//used to store errors for middleware package
var (
	ErrTooManyRequest = errors.New("you have exceeded your request limit")
	ErrInternalServer = errors.New("something went wrong on our end, try again later")
	ErrRouteNotFound  = errors.New("no matching route for proxy forwarding")
)
