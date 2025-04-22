package server

import (
	"go-gate/internal/server/middleware"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/proxy"
)

func (s *Server) registerRouter() {
	//proxy routing
	s.Mux.Handle("GET /api/", middleware.Handle()(proxy.ReverseProxy(s.Db)))

	//routes
	s.Mux.Handle("POST /api/routes", limiter.RateLimiter((s.Routes.AddRoute())))
	s.Mux.Handle("GET /api/routes", limiter.RateLimiter((s.Routes.GetAll())))
	s.Mux.Handle("GET /api/routes/{id}", limiter.RateLimiter((s.Routes.GetRouteByID())))
}