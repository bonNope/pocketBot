package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/pkg/errors"
)

var (
	errInvalidURL   = errors.New("url is invalid")
	errUnauthorized = errors.New("user is not authorized")
	errUnableToSave = errors.New("unable to save")
)

func (h *TelegramBotHandler) handleError(ChatID int64, err error) {
	msg := tgbotapi.NewMessage(ChatID, h.messages.Default)
	switch err {
	case errInvalidURL:
		msg.Text = h.messages.InvalidURL
		h.bot.Send(msg)
	case errUnauthorized:
		msg.Text = h.messages.Unauthorized
		h.bot.Send(msg)
	case errUnableToSave:
		msg.Text = h.messages.UnableToSave
		h.bot.Send(msg)
	default:
		h.bot.Send(msg)
	}

}
