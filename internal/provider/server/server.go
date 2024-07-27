package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	"gotik/internal/provider/routes"

	"github.com/rs/zerolog/log"
)

// Graceful Shutdown: https://medium.com/@dsilverdi/graceful-shutdown-in-go-a-polite-way-to-end-programs-6af16e025549
// Function to start the HTTP Server, with context
func StartHTTPServer(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Info().Msg("Hello")
		w.Write([]byte("Hello World"))
	})

	mux.Handle("/api/user/", routes.MuxUser())
	mux.Handle("/api/event/", routes.MuxEvent())
	mux.Handle("/api/ticket/", routes.MuxTicket())
	mux.Handle("/api/history/", routes.MuxTransactionDetail())

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
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

	log.Error().Msg("HTTP server stopped")
}
