package bot

import (
	"context"
	"github.com/bonNope/pocketBot/pkg/pocket"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"net/url"
)

const (
	commandStart = "start"
)

func (h *TelegramBotHandler) handleMessage(message *tgbotapi.Message) error {
	accessToken, err := h.getAccessToken(message.Chat.ID)
	if err != nil {
		return errUnauthorized
	}

	_, err = url.ParseRequestURI(message.Text)
	if err != nil {
		return errInvalidURL
	}

	if err := h.services.PocketClient.Add(context.Background(), pocket.AddInput{
		URL:         message.Text,
		AccessToken: accessToken,
	}); err != nil {
		return errUnableToSave
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, "Ссылка успешно сохранена!")
	_, err = h.bot.Send(msg)
	return err
}

func (h *TelegramBotHandler) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case commandStart:
		return h.handleStartCommand(message)
	default:
		return h.handleUnknownCommand(message)
	}
}

func (h *TelegramBotHandler) handleStartCommand(message *tgbotapi.Message) error {
	_, err := h.getAccessToken(message.Chat.ID)
	if err != nil {
		return h.initAuthorizationProcess(message)
	}

	msg := tgbotapi.NewMessage(message.Chat.ID, h.messages.AlreadyAuthorized)
	_, err = h.bot.Send(msg)
	return err
}

func (h *TelegramBotHandler) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, h.messages.UnknownCommand)

	_, err := h.bot.Send(msg)
	return err
}
