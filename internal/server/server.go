package server

import (
	"database/sql"
	"go-gate/internal/server/handler/mapping"
	"go-gate/internal/server/handler/proxy"
	"go-gate/internal/server/middleware"
	"go-gate/internal/server/middleware/limiter"
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

func (s *Server) RegisterRouter() {  
	s.Mux.Handle("/api/", middleware.Handle()(proxy.ReverseProxy(s.Db)))
	s.Mux.Handle("/api/mapping", limiter.RateLimiter((mapping.GetRequestMappings(s.Db))))
}