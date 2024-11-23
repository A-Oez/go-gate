package router

import (
	"go-gate/internal/server/handler/logging"
	"go-gate/internal/server/handler/proxy"
	"net/http"
)

func RegisterRouter(mux *http.ServeMux) {  
	mux.Handle("/api/", logging.InboundLogging(proxy.ReverseProxy()))
}