package proxy

import (
	"encoding/json"
	"go-gate/internal/service/routes/entity"
	"log"
	"net/http"
	"time"
)

//https://ndersson.me/post/capturing_status_code_in_net_http/
type LoggingResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
    lrw.StatusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

type proxyLog struct {
	RequestID    string `json:"request_id"`
	ClientPath    string `json:"client_path"`
	ServiceURL   string `json:"service_url"`
	StatusCode   int    `json:"status_code"`
	Timestamp    string `json:"timestamp"`
}

func Log(statusCode int, requestID string, entity entity.Route) {
	entry := proxyLog{
		RequestID:  requestID,
		ClientPath: entity.PublicPath,
		ServiceURL: entity.ServiceHost + entity.ServicePath,
		StatusCode: statusCode,
		Timestamp:  time.Now().Format(time.RFC3339),
	}

	if logBytes, err := json.Marshal(entry); err == nil {
		log.Println(string(logBytes))
	}
}