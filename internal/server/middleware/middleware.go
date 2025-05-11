package middleware

import (
	"go-gate/internal/server/middleware/limiter"
	"go-gate/internal/server/middleware/logging"
	adminauth "go-gate/internal/service/admin_auth/handler"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

func Handle() Middleware {
	return func(h http.Handler) http.Handler {
		return limiter.RateLimiter(logging.InboundLogging(h))
	}
}

func HandleAdmin(handler *adminauth.AdminAuthHandler) Middleware {
	return func(h http.Handler) http.Handler {
		return handler.AuthAdmin(limiter.RateLimiter(h))
	}
}