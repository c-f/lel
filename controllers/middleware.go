package controllers

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
)

var (
	AuthTokenPrefix            = "Token "
	ContextUserKey  ContextKey = "user"
)

type ContextKey string

func UserFromContext(ctx context.Context) string {
	return ctx.Value(ContextUserKey).(string)
}

func devMiddleware(next http.Handler) http.Handler {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Development allow preflights
		if r.Method == "OPTIONS" {
			JSONResponse(w, "debug", http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// logging
func (s *APIHandler) requireAuth(next http.HandlerFunc) http.HandlerFunc {
	if s.mode == "single" {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), ContextUserKey, "local")
			next(w, r.WithContext(ctx))
		})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		raw := r.Header.Get("Authorization")
		if ok := strings.HasPrefix(raw, AuthTokenPrefix); !ok {
			JSONResponse(w, "unauthenticated", http.StatusUnauthorized)
			return
		}
		input := strings.TrimPrefix(raw, AuthTokenPrefix)
		sum := sha256.Sum256([]byte(input))
		if len(sum) != len(input) {
			JSONResponse(w, "wrong length", http.StatusUnauthorized)
			return
		}
		hash := fmt.Sprintf("%x", sum)
		data, valid := s.auth[hash]
		if valid {
			JSONResponse(w, "unauthenticated", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), ContextUserKey, data)
		next(w, r.WithContext(ctx))
	})
}
