package main

import (
	"github.com/urfave/cli/v2"
)

var (
	port             string
	token            string
	unsplashClientID string

	flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			EnvVars:     []string{"PORT"},
			Destination: &port,
			Required:    true,
		},
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
