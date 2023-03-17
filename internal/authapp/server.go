package authapp

import (
	"net/http"
)

type AuthorizationServer struct {
	server *http.Server
}

func NewAuthorizationServer() *AuthorizationServer {
	return &AuthorizationServer{}
}

func (s *AuthorizationServer) Start(handler http.Handler) error {
	s.server = &http.Server{
		Addr:    ":60",
		Handler: handler,
	}

	return s.server.ListenAndServe()
}
