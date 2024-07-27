package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Graveful Shutdown: https://medium.com/@dsilverdi/graceful-shutdown-in-go-a-polite-way-to-end-programs-6af16e025549
// Function to start the HTTP Server, with context
func startHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	routes := http.NewServeMux()

	// Routes
	routes.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Hello")
		w.Write([]byte("Hello World"))
	})

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: routes,
	}

	// The actual starting point for the HTTP server
	go func() {
		log.Info().Msg("Trying to start server on port 8080...")

		err := server.ListenAndServe()

		log.Info().Msg("Server started on port 8080")

		if err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Error starting server")
		}
	}()

	// If the channel is closed, channel will return error
	// the context sent from main include a channel to communicate
	// ctx.Done (channel) is closed if the context is canceled
	// Checking if the context is canceled
	if err := ctx.Err(); err != nil {
		log.Error().Err(err).Msg("Shutting down server...")
		// Context with 5 seconds timeout
		shutdown_ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// Shutting down the server, using the context to set the timeout within 5 seconds
		err := server.Shutdown(shutdown_ctx)
		if err != nil {
			log.Error().Err(err).Msg("Error shutting down server")
		}
	}

	// Server stopped
	log.Error().Msg("HTTP server stopped")
}

func main() {
	// Zerolog Time Formating
	zerolog.TimeFieldFormat = time.RFC1123

	// Context with the ability to cancel the context
	ctx, cancel := context.WithCancel(context.Background())

	// Craeting a *sync.Waitgroup
	wg := &sync.WaitGroup{}

	// Wait groupt counter + 1 for the http server
	wg.Add(1)

	// starting the HTTP server via a goroutine
	go startHTTPServer(ctx, wg)

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
