package proxy

import (
	"database/sql"
	entity "go-gate/internal/db/entity/mapping"
	repo "go-gate/internal/db/repo/mapping"
	"go-gate/internal/server/middleware/logging"
	service "go-gate/internal/service/mapping"
	"go-gate/pkg/proxy/proxylog"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func ReverseProxy(db *sql.DB) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := getRequestID(r)
		if requestID == "" {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		service := service.NewMappingService(repo.NewMappingRepository(db))
		proxyRoute, err := service.GetRequestByClient(r.Method, trimSuffix(r.URL.Path))
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusNotFound)
			return
		}
		
		proxy := newReverseProxy(requestID, proxyRoute)
        lrw := &proxylog.LoggingResponseWriter{w, http.StatusOK}
		proxy.ServeHTTP(lrw, r)

		proxylog.Log(lrw.StatusCode, requestID, proxyRoute)
    })
}

func getRequestID(r *http.Request) string {
	id, ok := r.Context().Value(logging.ContextKey("RequestID")).(string)
	if !ok {
		return ""
	}
	return id
}

func newReverseProxy(requestID string, entity entity.ProxyMapping) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Del("X-Forwarded-For")
			req.Header.Del("X-Real-IP")

			req.URL.Scheme = entity.ServiceScheme
			req.URL.Host = entity.ServiceHost
			req.URL.Path = entity.ServicePath

			req.Header.Set("X-Forwarded-Proto", "https")
			req.Header.Set("X-Forwarded-Host", req.Host)
			req.Header.Set("X-Request-ID", requestID)
			req.Header.Set("X-Request-Timestamp", time.Now().Format(time.RFC3339))
		},
	}
}

func trimSuffix(path string) string {
	if strings.HasSuffix(path, "/") {
		return path[:len(path)-1]
	}
	return path
}