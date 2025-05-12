package proxy

import (
	"context"
	"database/sql"
	"errors"
	"go-gate/internal/server/middleware/logging"
	"go-gate/internal/service/routes"
	"go-gate/internal/service/routes/entity"
	"go-gate/internal/service/routes/repo"
	"go-gate/pkg/httperror"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

type ContextKey string

func ReverseProxy(db *sql.DB) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := getRequestID(r)
		remoteAddr := getRemoteAddr(r)
		if requestID == "" {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg: httperror.ErrInternalServer.Error(),
				Error: errors.New("couldnt create request ID, check inbound_logging middleware"),
			}.WriteError(w)
            return
		}

		service := routes.NewRoutesService(repo.NewRouteRepository(db))
		proxyRoute, err := service.GetRouteByClient(r.Method, trimSuffix(r.URL.Path))
		if err != nil {
			httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg: httperror.ErrRouteNotFound.Error(),
			}.WriteError(w)
            return
		}
		
 		ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKey("RouteID"), proxyRoute.ID)
		r = r.WithContext(ctx)		
		
		proxy := newReverseProxy(requestID, remoteAddr, proxyRoute)
        lrw := &LoggingResponseWriter{w, http.StatusOK}
		proxy.ServeHTTP(lrw, r)

		Log(lrw.StatusCode, requestID, proxyRoute)
    })
}

func getRequestID(r *http.Request) string {
	id, ok := r.Context().Value(logging.ContextKey("RequestID")).(string)
	if !ok {
		return ""
	}
	return id
}

func getRemoteAddr(r *http.Request) string {
	ip, ok := r.Context().Value(logging.ContextKey("RemoteAddr")).(string)
	if !ok {
		return ""
	}
	return ip
}

func newReverseProxy(requestID string, remoteAddr string, entity entity.Route) *httputil.ReverseProxy {
	return &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.Header.Del("X-Forwarded-For")
			req.Header.Del("X-Real-IP")

			req.URL.Scheme = entity.ServiceScheme
			req.URL.Host = entity.ServiceHost
			req.URL.Path = entity.ServicePath

			req.Header.Set("X-Client-IP", remoteAddr)
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