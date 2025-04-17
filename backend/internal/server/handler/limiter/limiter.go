package limiter

import (
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

type message struct{
	Status string
	Body string
}

//global rate limiter 
func RateLimiter(next http.Handler) http.Handler {
	//TODO add logic to implement rate limiting per client
	limiter := rate.NewLimiter(5, 10)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if limiter.Allow() {
			next.ServeHTTP(w, r)
		} else {
			msg := message{
                Status: "Request Failed",
                Body:   "Try again later.",
            }

            w.WriteHeader(http.StatusTooManyRequests)
            json.NewEncoder(w).Encode(&msg)
            return
		}
	})
}