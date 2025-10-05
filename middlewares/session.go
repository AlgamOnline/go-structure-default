package middlewares

import (
	"context"
	"golang-default/services"
	"net/http"
	"strings"
)

func SessionMiddleware(sessionService *services.SessionService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header", http.StatusUnauthorized)
				return
			}

			token := parts[1]
			valid, userID, err := sessionService.ValidateToken(token)
			if err != nil || !valid {
				http.Error(w, "Session expired or invalid", http.StatusUnauthorized)
				return
			}

			// Simpan userID di context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "userID", userID)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}
