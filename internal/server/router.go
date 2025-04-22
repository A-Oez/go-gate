package server

import (
	"go-gate/internal/server/middleware"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/proxy"
	routes "go-gate/internal/service/routes/handler"
)

func (s *Server) registerRouter() {
	//proxy routing
	s.Mux.Handle("GET /api/", middleware.Handle()(proxy.ReverseProxy(s.Db)))

	//routes
	s.Mux.Handle("POST /api/routes", limiter.RateLimiter((routes.AddRequest(s.Db))))
	s.Mux.Handle("GET /api/routes", limiter.RateLimiter((routes.GetRequestMappings(s.Db))))
	s.Mux.Handle("GET /api/routes/{id}", limiter.RateLimiter((routes.GetRequestMappingByID(s.Db))))
}