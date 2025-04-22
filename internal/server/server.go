package server

import (
	"database/sql"
	"net/http"
)

type Server struct {
	Db  *sql.DB
	Mux *http.ServeMux
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		Db:  db,
		Mux: http.NewServeMux(),
	}
}

func (s *Server) Start(port string) {
	s.registerRouter()
	http.ListenAndServe(port, s.Mux)
}