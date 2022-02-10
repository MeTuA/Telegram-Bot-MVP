package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/labstack/echo/v4"
)

func startService(ctx context.Context, wg *sync.WaitGroup, e *echo.Echo) {

	wg.Add(1)

	go func() {
		err := e.Start(":8080")
		if err != nil && err != http.ErrServerClosed {
			fmt.Println("error", err.Error())
		}

		wg.Done()
	}()

	go func() {

		<-ctx.Done()
		err := e.Shutdown(context.TODO())
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}()

}

func waitForExit(cancel func(), wg *sync.WaitGroup) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	cancel()
	wg.Wait()
	fmt.Println("exiting")
}
