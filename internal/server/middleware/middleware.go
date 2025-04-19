package middleware

import (
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