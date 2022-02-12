package telegram

import (
	"bot/mvp/unsplash"
	"context"
	"fmt"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/hashicorp/go-hclog"
)

type Service interface {
	GetUpdates(ctx context.Context, wg *sync.WaitGroup, token string)
}

type service struct {
	unsplash unsplash.Service
	log      hclog.Logger
}

func NewService(unsplash unsplash.Service, log hclog.Logger) Service {
	return &service{
		unsplash: unsplash,
		log:      log,
	}
}

func (s *service) GetUpdates(ctx context.Context, wg *sync.WaitGroup, token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	wg.Add(1)

	go func() {
		<-ctx.Done()
		fmt.Println("stopping getUpdates")
		wg.Done()
		return
	}()

	// bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "help":
			msg.Text = "I understand /sayhi, /status and /random."
		case "sayhi":
			msg.Text = "Hi :)"
		case "status":
			msg.Text = "I'm ok."
		case "random":
			// sending messsage with random photo from Unsplash API
			randomPhoto, _ := s.unsplash.GetRandomPhoto()

			file := tgbotapi.FileURL(randomPhoto.URL.Regular)
			file.NeedsUpload()

			photoMsg := tgbotapi.NewPhoto(update.Message.Chat.ID, nil)
			photoMsg.File = file
			photoMsg.Caption = randomPhoto.Description

			bot.Send(photoMsg)
			continue
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
