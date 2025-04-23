package limiter

import (
	"go-gate/pkg/httperror"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

var (
	ipLimiters = make(map[string]*rate.Limiter)
	mu        sync.Mutex
)

func RateLimiter(next http.Handler) http.Handler {
	go func() {
		for {
			time.Sleep(15 * time.Minute)

			mu.Lock()
			for ip := range ipLimiters {
				delete(ipLimiters, ip)
			}
			mu.Unlock()
		}
	}()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr

		mu.Lock()
		defer mu.Unlock()

		if _, exists := ipLimiters[ip]; !exists {
			ipLimiters[ip] = rate.NewLimiter(5, 10)
		}

		limiter := ipLimiters[ip]
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			httperror.DefaultError{
				Status: http.StatusTooManyRequests,
				Msg: httperror.ErrTooManyRequest.Error(),
			}.WriteError(w)
            return
		}
	})
}