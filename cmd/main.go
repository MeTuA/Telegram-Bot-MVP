package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/labstack/echo/v4"

	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Name:  "PhotoRand Telegram Bot",
		Usage: "go run .",
	}

	app.Action = func(c *cli.Context) error {
		fmt.Println("Starting application")

		ctx, cancel := context.WithCancel(context.TODO())
		wg := &sync.WaitGroup{}

		e := echo.New()
		e.GET("/", func(c echo.Context) error { return c.JSON(200, "hi") })

		startService(ctx, wg, e)
		waitForExit(cancel, wg)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(1)
	}
}
