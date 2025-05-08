package server

import (
	"go-gate/internal/server/middleware"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/proxy"
)

func (s *Server) registerRouter() {
	//proxy routing
	s.Mux.Handle("GET /api/", middleware.Handle()(proxy.ReverseProxy(s.Db)))

	//admin routes
	s.Mux.Handle("POST /admin/login", limiter.RateLimiter(s.AdminAuth.Login()))
	s.Mux.Handle("POST /admin/routes", limiter.RateLimiter((s.Routes.AddRoute())))
	s.Mux.Handle("PUT /admin/routes/{id}", limiter.RateLimiter((s.Routes.UpdateRoute())))
	s.Mux.Handle("GET /admin/routes", limiter.RateLimiter((s.Routes.GetAll())))
	s.Mux.Handle("GET /admin/routes/{id}", limiter.RateLimiter((s.Routes.GetRouteByID())))
	s.Mux.Handle("DELETE /admin/routes/{id}", limiter.RateLimiter(s.Routes.DeleteRouteByID()))
}