package middleware

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		wrapper := &WrapperWriter{
			ResponseWriter: w,
			StatusCode:     http.StatusOK,
		}
		next.ServeHTTP(wrapper, r)
		log.WithFields(log.Fields{
    "Status code": wrapper.StatusCode,
    "Method":   r.Method,
		"Path":     r.URL.Path,
		"Duration": time.Since(start),
  }).Info("Request completed")
	})
}