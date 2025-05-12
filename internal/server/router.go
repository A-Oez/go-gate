package server

import (
	"go-gate/internal/server/middleware"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/proxy"
)

func (s *Server) registerRouter() {
	//proxy routing
	s.Mux.Handle("GET /api/", middleware.Handle(s.Db)(proxy.ReverseProxy(s.Db)))

	//admin routes
	s.Mux.Handle("POST /admin/login", limiter.RateLimiter(s.AdminAuth.Login()))
	s.Mux.Handle("POST /admin/routes", middleware.HandleAdmin(s.AdminAuth)(s.Routes.AddRoute()))
	s.Mux.Handle("PUT /admin/routes/{id}", middleware.HandleAdmin(s.AdminAuth)(s.Routes.UpdateRoute()))
	s.Mux.Handle("GET /admin/routes", middleware.HandleAdmin(s.AdminAuth)(s.Routes.GetAll()))
	s.Mux.Handle("GET /admin/routes/{id}", middleware.HandleAdmin(s.AdminAuth)(s.Routes.GetRouteByID()))
	s.Mux.Handle("DELETE /admin/routes/{id}", middleware.HandleAdmin(s.AdminAuth)(s.Routes.DeleteRouteByID()))
}