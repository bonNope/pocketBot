package bot

import (
	"github.com/bonNope/pocketBot/internal/config"
	"github.com/bonNope/pocketBot/internal/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBotHandler struct {
	bot         *tgbotapi.BotAPI
	services    *service.Service
	redirectURL string
	messages    config.Messages
}

func NewTelegramBotHandler(bot *tgbotapi.BotAPI, services *service.Service, redirectURL string, messages config.Messages) *TelegramBotHandler {
	return &TelegramBotHandler{bot: bot, services: services, redirectURL: redirectURL, messages: messages}
}

func (h *TelegramBotHandler) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := h.handleCommand(update.Message); err != nil {
					h.handleError(update.Message.Chat.ID, err)
				}
			} else {
				if err := h.handleMessage(update.Message); err != nil {
					h.handleError(update.Message.Chat.ID, err)
				}
			}
		}
	}
}

func (h *TelegramBotHandler) InitUpdatesChannel() (tgbotapi.UpdatesChannel, error) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return h.bot.GetUpdatesChan(u)
}

func (h *TelegramBotHandler) GetBotName() string {
	return h.bot.Self.UserName
}
