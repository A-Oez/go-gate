package router

import (
	"go-gate/internal/server/handler/limiter"
	"go-gate/internal/server/handler/logging"
	"go-gate/internal/server/handler/proxy"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {  
	mux.Handle("/api/", limiter.RateLimiter(logging.InboundLogging(proxy.ReverseProxy())))
}