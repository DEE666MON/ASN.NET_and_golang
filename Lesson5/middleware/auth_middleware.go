package middleware

import (
	"net/http"
	"strings"
	"user-api/auth"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Неудачная авторизация.", http.StatusUnauthorized)
			return
		}
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenStr == authHeader {
			http.Error(w, "Неправильный формат токена.", http.StatusUnauthorized)
			return
		}
		_, err := auth.ParseToken(tokenStr)
		if err != nil {
			http.Error(w, "Неправильный токен.", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}