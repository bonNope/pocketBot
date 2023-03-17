package auth

import (
	"github.com/bonNope/pocketBot/internal/service"
	"net/http"
	"strconv"
)

type AuthorizationHandler struct {
	services    *service.Service
	redirectURL string
}

func NewAuthorizationHandler(services *service.Service, redirectURL string) *AuthorizationHandler {
	return &AuthorizationHandler{services: services, redirectURL: redirectURL}
}

func (h *AuthorizationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := h.services.Get(chatID, service.RequestTokens)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResponse, err := h.services.PocketClient.Authorize(r.Context(), requestToken)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = h.services.Save(chatID, authResponse.AccessToken, service.AccessTokens)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Location", h.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
}
