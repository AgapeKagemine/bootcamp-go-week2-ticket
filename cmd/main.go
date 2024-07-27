package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"gotik/internal/provider/server"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	// Zerolog Time Formating > "Mon, 02 Jan 2006 15:04:05 MST"
	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:          os.Stdout,
		TimeFormat:   time.RFC1123,
		TimeLocation: time.Local,
	})

	// Context with the ability to cancel the context
	ctx, cancel := context.WithCancel(context.Background())

	// Craeting a *sync.Waitgroup
	wg := &sync.WaitGroup{}

	// Wait groupt counter + 1 for the http server
	wg.Add(1)

	// starting the HTTP server
	server.StartHTTPServer(ctx, wg)

	// Channel to capture OS signal
	signal_channel := make(chan os.Signal, 1)

	// Sending the signal from OS to the channel
	signal.Notify(signal_channel, syscall.SIGINT, syscall.SIGTERM)

	// Receiving os signal to terminate the program
	// SIGINT  = Signal Interupt
	// SIGTERM = Signal Terminate
	<-signal_channel

	log.Info().Msg("Gracefully shutting down http server...")

	// Canceling the context
	cancel()

	// Wait till the goroutine for the HTTP server is stopped
	wg.Wait()

	log.Log().Msg("Shutdown complete")
}
