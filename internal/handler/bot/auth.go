package bot

import (
	"context"
	"fmt"
	"github.com/bonNope/pocketBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (h *TelegramBotHandler) initAuthorizationProcess(message *tgbotapi.Message) error {
	authLink, err := h.generateAuthorizationLink(message.Chat.ID)
	if err != nil {
		return err
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf(h.messages.Start, authLink))

	_, err = h.bot.Send(msg)
	return err
}

func (h *TelegramBotHandler) generateAuthorizationLink(chatID int64) (string, error) {
	redirectURL := h.generateRedirectURL(chatID)
	requestToken, err := h.services.PocketClient.GetRequestToken(context.Background(), redirectURL)
	if err != nil {
		return "", err
	}

	if err := h.services.Save(chatID, requestToken, service.RequestTokens); err != nil {
		return "", err
	}

	return h.services.PocketClient.GetAuthorizationURL(requestToken, redirectURL)
}

func (h *TelegramBotHandler) getAccessToken(chatID int64) (string, error) {
	return h.services.Get(chatID, service.AccessTokens)
}

func (h *TelegramBotHandler) generateRedirectURL(chatID int64) string {
	return fmt.Sprintf("%s?chat_id=%d", h.redirectURL, chatID)
}
