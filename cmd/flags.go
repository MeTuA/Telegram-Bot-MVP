package main

import (
	"github.com/urfave/cli/v2"
)

var (
	token            string
	unsplashClientID string

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "token",
			EnvVars:     []string{"TELEGRAM_TOKEN"},
			Destination: &token,
			Required:    true,
		},
		&cli.StringFlag{
			Name:        "unsplashClientID",
			EnvVars:     []string{"UNSPLASH_CLIENTID"},
			Destination: &unsplashClientID,
			Required:    true,
		},
	}
)
