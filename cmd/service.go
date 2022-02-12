package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func waitForExit(cancel func(), wg *sync.WaitGroup) {
	fmt.Println("starting waitForExit")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	cancel()
	wg.Wait()
	fmt.Println("Stopping application")
}
