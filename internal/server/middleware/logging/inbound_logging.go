package logging

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"go-gate/pkg/httperror"
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

func InboundLogging(db *sql.DB, next http.Handler) http.Handler {
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

        err := write(db, logData)
        if err != nil {
            httperror.DefaultError{
				Status: http.StatusInternalServerError,
				Msg: err.Error(),
			}.WriteError(w)
            return
        }

        ctx := r.Context()
		ctx = context.WithValue(ctx, ContextKey("RequestID"), requestID)
        ctx = context.WithValue(ctx, ContextKey("RemoteAddr"), r.RemoteAddr)
		r = r.WithContext(ctx)

        next.ServeHTTP(w, r)
    })
}

func write(db *sql.DB, log logInbound) error {
    query := `
        INSERT INTO log_inbound (request_id, time, method, uri, remote, headers, body)
        VALUES ($1, $2, $3, $4, $5, $6, $7);
    `

    parsedTime, err := time.Parse(time.RFC3339, log.Time)
    if err != nil {
        return fmt.Errorf("invalid time format: %w", err)
    }

    headersJSON, err := json.Marshal(log.Headers)
    if err != nil {
        return fmt.Errorf("error serializing headers: %w", err)
    }

    var bodyBytes []byte
    if log.Body != nil {
        bodyBytes, err = io.ReadAll(log.Body)
        if err != nil {
            return fmt.Errorf("error reading body: %w", err)
        }
    }

    _, err = db.Exec(query,
        log.RequestID,
        parsedTime,
        log.Method,
        log.URI,
        log.Remote,
        headersJSON,
        string(bodyBytes),
    )

    return err
}