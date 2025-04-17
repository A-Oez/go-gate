package router

import (
	"go-gate/backend/internal/server/handler/limiter"
	"go-gate/backend/internal/server/handler/logging"
	"go-gate/backend/internal/server/handler/proxy"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {  
	mux.Handle("/api/", limiter.RateLimiter(logging.InboundLogging(proxy.ReverseProxy())))
}