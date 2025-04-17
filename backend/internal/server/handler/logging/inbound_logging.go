package logging

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ContextKey string

type logInbound struct {
    RequestID string    `json:"request_id"`
    Time      string    `json:"time"`
    Method    string    `json:"method"`
    URI       string    `json:"uri"`
    Remote    string    `json:"remote"`
    Headers   http.Header `json:"headers"`
    Body      io.Reader `json:"body"`
}

func InboundLogging(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := uuid.New().String()
        
        logData := logInbound{
            RequestID: requestID,
            Headers:   r.Header,
            URI:       r.RequestURI,
            Method:    r.Method,
            Remote:    r.RemoteAddr,
            Body:      r.Body,
            Time:      time.Now().Format(time.RFC3339),
        }
    
        logBytes, _ := json.Marshal(logData)
        log.Println(string(logBytes))

        ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKey("RequestID"), requestID)
		r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}