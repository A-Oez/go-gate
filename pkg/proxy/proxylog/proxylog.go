package proxylog

import "net/http"

//https://ndersson.me/post/capturing_status_code_in_net_http/
type LoggingResponseWriter struct {
    http.ResponseWriter
    StatusCode int
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
    lrw.StatusCode = code
    lrw.ResponseWriter.WriteHeader(code)
}

type ProxyLogEntry struct {
	RequestID    string `json:"request_id"`
	ClientPath    string `json:"client_path"`
	ServiceURL   string `json:"service_url"`
	StatusCode   int    `json:"status_code"`
	Timestamp    string `json:"timestamp"`
}