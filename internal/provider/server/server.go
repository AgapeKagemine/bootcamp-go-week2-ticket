package server

import (
	"context"
	"fmt"
	"gotik/internal/provider/routes"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type ServerConfig struct {
	Address string
	Port    uint
}

// Graceful Shutdown: https://medium.com/@dsilverdi/graceful-shutdown-in-go-a-polite-way-to-end-programs-6af16e025549
// Function to start the HTTP Server, with context
func StartHTTPServer() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	router := routes.NewRoutes()

	serverConfig := ServerConfig{
		Address: "127.0.0.1",
		Port:    8080,
	}

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", serverConfig.Address, serverConfig.Port),
		Handler: router.Server,
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

	// Listen for Interrupt signal (Ctrl+C or other termination signal)
	<-ctx.Done()

	// Restore Default Behaviour - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go
	stop()
	log.Info().Msg("Shutting down server...")

	// If the channel is closed, channel will return error
	// the context sent from main include a channel to communicate
	// ctx.Done (channel) is closed if the context is canceled
	// Checking if the context is canceled
	// Context with 5 seconds timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutting down the server, using the context to set the timeout within 5 seconds
	err := server.Shutdown(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Error shutting down server")
	}

	log.Error().Msg("HTTP server stopped")
}
