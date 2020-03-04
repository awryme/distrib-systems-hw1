package server

import (
	"go.uber.org/zap"
	"incrementer/log"
	"net/http"
)

type Middleware func(handler http.Handler) http.Handler

func middlewareUseLogger(logger *zap.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = log.ToReq(r, logger)
			next.ServeHTTP(w, r)
		})
	}
}

func middlewareSetApiVersion(version string) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = log.WithReq(r, zap.String("api-version", version))
			next.ServeHTTP(w, r)
		})
	}
}
