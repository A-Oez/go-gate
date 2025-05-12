package middleware

import (
	"database/sql"
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/logging"
	adminauth "go-gate/internal/service/admin_auth/handler"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Handle(db *sql.DB) Middleware {
	return func(h http.Handler) http.Handler {
		return limiter.RateLimiter(logging.InboundLogging(db, h))
	}
}

func HandleAdmin(handler *adminauth.AdminAuthHandler) Middleware {
	return func(h http.Handler) http.Handler {
		return handler.AuthAdmin(limiter.RateLimiter(h))
	}
}