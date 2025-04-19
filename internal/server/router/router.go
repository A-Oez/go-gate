package router

import (
	"go-gate/internal/server/handler/proxy"
	"go-gate/internal/server/middleware"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {  
	mux.Handle("/api/", middleware.Handle()(proxy.ReverseProxy()))
}