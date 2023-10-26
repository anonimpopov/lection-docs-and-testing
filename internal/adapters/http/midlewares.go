package http

import (
	"context"
	"net/http"

	"ru.mts.teta.tests_and_docs/internal/domain/models"
)

type ctxKeyUser struct{}

const (
	errInvalidToken = "invalid token"
)

func (s *Server) ValidateAuth() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessToken, err := r.Cookie("access")
			if err != nil {
				http.Error(w, errInvalidToken, http.StatusUnauthorized)
				return
			}
			refreshToken, err := r.Cookie("refresh")
			if err != nil {
				http.Error(w, errInvalidToken, http.StatusUnauthorized)
				return
			}

			userId, err := s.auth.Validate(r.Context(), models.TokenPair{
				AuthToken: accessToken.Value,
				RefreshToken: refreshToken.Value,
			})
			if err != nil {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
			ctx := r.Context()

			ctx = context.WithValue(ctx, ctxKeyUser{}, userId)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
