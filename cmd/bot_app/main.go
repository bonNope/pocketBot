package main

import (
	"github.com/boltdb/bolt"
	"github.com/bonNope/pocketBot/internal/authapp"
	"github.com/bonNope/pocketBot/internal/botapp"
	"github.com/bonNope/pocketBot/internal/config"
	"github.com/bonNope/pocketBot/internal/handler/auth"
	"github.com/bonNope/pocketBot/internal/handler/bot"
	"github.com/bonNope/pocketBot/internal/repository"
	"github.com/bonNope/pocketBot/internal/service"
	"github.com/bonNope/pocketBot/pkg/pocket"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	telegramBotAPI, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}

	telegramBotAPI.Debug = true

	pocketClient, err := pocket.NewClient(cfg.PocketConsumerKey)
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewRepository(db)
	services := service.NewService(repo, pocketClient)

	botHandlers := bot.NewTelegramBotHandler(telegramBotAPI, services, cfg.AuthServerUrl, cfg.Messages)
	authHandler := auth.NewAuthorizationHandler(services, cfg.TelegramBotURL)

	authorizationServer := authapp.NewAuthorizationServer()
	telegramBot := botapp.NewBot(botHandlers)

	go func() {
		if err := telegramBot.Start(); err != nil {
			log.Fatal(err)
		}
	}()

	if err := authorizationServer.Start(authHandler); err != nil {
		log.Fatal(err)
	}

}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(service.AccessTokens))
		if err != nil {
			return err
		}

		_, err = tx.CreateBucketIfNotExists([]byte(service.RequestTokens))
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return db, nil
}
