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
	//proxy routing
	s.Mux.Handle("GET /api/", middleware.Handle()(proxy.ReverseProxy(s.Db)))

	//mapping
	s.Mux.Handle("POST /api/mapping", limiter.RateLimiter((mapping.AddRequest(s.Db))))
	s.Mux.Handle("GET /api/mapping", limiter.RateLimiter((mapping.GetRequestMappings(s.Db))))
	s.Mux.Handle("GET /api/mapping/{id}", limiter.RateLimiter((mapping.GetRequestMappingByID(s.Db))))
}