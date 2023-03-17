package botapp

import (
	"github.com/bonNope/pocketBot/internal/handler/bot"
	"log"
)

type Bot struct {
	h *bot.TelegramBotHandler
}

func NewBot(handler *bot.TelegramBotHandler) *Bot {
	return &Bot{h: handler}
}

func (b *Bot) Start() error {
	log.Printf("Authorized on account %s", b.h.GetBotName())

	updates, err := b.h.InitUpdatesChannel()
	if err != nil {
		return err
	}
	b.h.HandleUpdates(updates)
	return nil
}
