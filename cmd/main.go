package main

import (
	"bot/mvp/telegram"
	"bot/mvp/unsplash"
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/hashicorp/go-hclog"

	"github.com/urfave/cli/v2"
)

func main() {

	log := hclog.New(&hclog.LoggerOptions{
		Name: "telegram bot",
	})

	app := &cli.App{
		Name:  "PhotoRand Telegram Bot",
		Usage: "go run .",
		Flags: flags,
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Starting application")

		ctx, cancel := context.WithCancel(context.TODO())
		wg := &sync.WaitGroup{}

		unsplashCredentials := unsplash.Credentials{
			ClientID: unsplashClientID,
		}

		unsplashService := unsplash.NewService(unsplashCredentials, log.Named("unsplashService"))

		tgService := telegram.NewService(unsplashService, log.Named("telegramService"))
		go func() {
			tgService.GetUpdates(ctx, wg, token)
		}()

		waitForExit(cancel, wg)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		os.Exit(1)
	}
}
