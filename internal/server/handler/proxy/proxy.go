package proxy

import (
	"encoding/json"
	"go-gate/internal/server/handler/logging"
	"log"
	"net/http"
	"net/http/httputil"
	"time"
)

//https://ndersson.me/post/capturing_status_code_in_net_http/
type statusCapturingResponseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (lrw *statusCapturingResponseWriter) WriteHeader(code int) {
    lrw.statusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

type proxyLogEntry struct {
	RequestID    string `json:"request_id"`
	ClientPath    string `json:"client_path"`
	ServiceURL   string `json:"service_url"`
	StatusCode   int    `json:"status_code"`
	Timestamp    string `json:"timestamp"`
}


func ReverseProxy() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID, ok := r.Context().Value(logging.ContextKey("RequestID")).(string)
		if !ok {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		proxyRoute, err := FindRouteMapping(ClientRequest{
			Method: r.Method,
			Path: r.URL.Path,
		},"")
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusNotFound)
			return
		}
		
		//added nil because all values filled in director
		proxy := httputil.NewSingleHostReverseProxy(nil)
		proxy.Director = func(req *http.Request) {
			//delete to prevent ip spooling
			req.Header.Del("X-Forwarded-For")

			req.URL.Scheme = proxyRoute.ServiceScheme
			req.URL.Host = proxyRoute.ServiceHost
			req.URL.Path = proxyRoute.ServicePath
		}
        lrw := &statusCapturingResponseWriter{w, http.StatusOK}

		proxy.ServeHTTP(lrw, r)

		logData := proxyLogEntry{
			RequestID:  requestID,
			ClientPath:  proxyRoute.ClientPath,
			ServiceURL: proxyRoute.ServiceHost + proxyRoute.ServicePath,
			StatusCode: lrw.statusCode,
			Timestamp:  time.Now().Format(time.RFC3339),
        }

		logBytes, _ := json.Marshal(logData)
		log.Println(string(logBytes))
    })
}

