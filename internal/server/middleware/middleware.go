package middleware

import (
	"database/sql"
	adminauth "go-gate/internal/server/middleware/admin_auth"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/logging"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Handle() Middleware {
	return func(h http.Handler) http.Handler {
		return limiter.RateLimiter(logging.InboundLogging(h))
	}
}

func HandleAdmin(db *sql.DB) Middleware {
	return func(h http.Handler) http.Handler {
		return adminauth.AuthAdmin(db, limiter.RateLimiter(h))
	}
}
