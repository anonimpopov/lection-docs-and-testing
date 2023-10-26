package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) authHandlers() http.Handler {
	h := chi.NewMux()
	h.Route("/", func(r chi.Router) {
		h.Use(s.ValidateAuth())
		h.Post("/validate", s.Validate)
		h.Post("/info", s.Validate)
	})
	h.Post("/login", s.Login)
	return h
}

// Validate
// @ID Validate
// @tags auth
// @Summary Validate tokens
// @Description Validate tokens and refresh tokens if refresh token is valid
// @Security Auth
// @Param tokens body TokenPair{} false "user tokens"
// @Failure 400 {object} TokenPair "bad request"
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router /validate [post]
func (s *Server) Validate(w http.ResponseWriter, r *http.Request) {

}

// Login
// @ID Login
// @tags auth
// @Summary Generate auth tokens.
// @Description Validate credentials, return access and refresh tokens.
// @Param credentials body User{} false "user credentials"
// @Success 200
// @Failure 403 {string} string "forbidden"
// @Failure 500 {string} string "internal error"
// @Router /login [post]
func (s *Server) Login(w http.ResponseWriter, r *http.Request) {

}
