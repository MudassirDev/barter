package web

import (
	"context"
	"net/http"

	"github.com/MudassirDev/barter/internal/auth"
)

func (c *apiConfig) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(AUTH_KEY)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "user not logged in", err)
			return
		}

		userId, err := auth.ValidateJWT(cookie.Value, c.JWTSecretKey)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "user not logged in", err)
			return
		}
		ctx := context.WithValue(r.Context(), AUTH_KEY, userId)

		request := r.WithContext(ctx)
		next.ServeHTTP(w, request)
	})
}
