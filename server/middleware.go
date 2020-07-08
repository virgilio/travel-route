package server

import (
	"context"
	"net/http"
)

// ContextKey key type
type ContextKey string

// ContextMiddleware implements adding ContextKeys to handlerFunc
func ContextMiddleware(key ContextKey, value string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), key, value)
		handler.ServeHTTP(w, r.WithContext(ctx))
	})
}
