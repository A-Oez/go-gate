package proxy

import (
	"encoding/json"
	"go-gate/internal/server/middleware/logging"
	"go-gate/pkg/proxy/proxylog"
	"go-gate/pkg/proxy/routing"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"time"
)

func ReverseProxy() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := getRequestID(r)
		if requestID == "" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		proxyRoute, err := findProxyRoute(r)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusNotFound)
			return
		}
		
		proxy := newReverseProxy(requestID, proxyRoute)
        lrw := &proxylog.LoggingResponseWriter{w, http.StatusOK}
		proxy.ServeHTTP(lrw, r)

		logProxyRequest(lrw.StatusCode, requestID, proxyRoute)
    })
}

func getRequestID(r *http.Request) string {
	id, ok := r.Context().Value(logging.ContextKey("RequestID")).(string)
	if !ok {
		return ""
	}
	return id
}

func findProxyRoute(r *http.Request) (*routing.ProxyRoute, error) {
	path := "../internal/server/config/proxy_mapping.json"
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return routing.FindRouteMapping(routing.ClientRequest{
		Method: r.Method,
		Path:   routing.TrimSuffix(r.URL.Path),
	}, file)
}

func newReverseProxy(requestID string, route *routing.ProxyRoute) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Del("X-Forwarded-For")
			req.Header.Del("X-Real-IP")

			req.URL.Scheme = route.ServiceScheme
			req.URL.Host = route.ServiceHost
			req.URL.Path = route.ServicePath

			req.Header.Set("X-Forwarded-Proto", "https")
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Request-ID", requestID)
			req.Header.Set("X-Request-Timestamp", time.Now().Format(time.RFC3339))
		},
	}
}

func logProxyRequest(statusCode int, requestID string, route *routing.ProxyRoute) {
	entry := proxylog.ProxyLogEntry{
		RequestID:  requestID,
		ClientPath: route.ClientPath,
		ServiceURL: route.ServiceHost + route.ServicePath,
		StatusCode: statusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	if logBytes, err := json.Marshal(entry); err == nil {
		log.Println(string(logBytes))
	}
}