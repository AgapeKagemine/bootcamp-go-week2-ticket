package server

import (
	"context"
	"net/http"
	"sync"
	"time"

	eventHandler "gotik/internal/handler/event"
	"gotik/internal/provider/routes"
	eventRepository "gotik/internal/repository/event"
	eventUsecase "gotik/internal/usecase/event"

	ticketHandler "gotik/internal/handler/ticket"
	ticketRepository "gotik/internal/repository/ticket"
	ticketUsecase "gotik/internal/usecase/ticket"

	tdHandler "gotik/internal/handler/transactiondetail"
	tdRepository "gotik/internal/repository/transactiondetail"
	tdUsecase "gotik/internal/usecase/transactiondetail"

	userHandler "gotik/internal/handler/user"
	userRepository "gotik/internal/repository/user"
	userUsecase "gotik/internal/usecase/user"

	"github.com/rs/zerolog/log"
)

func autowired() (eventH eventHandler.EventHandler, ticketH ticketHandler.TicketHandler, tdH tdHandler.TransactionDetailHandler, userH userHandler.UserHandler) {
	ticketRepo := ticketRepository.NewTicketRepository()
	ticketUsecase := ticketUsecase.NewTicketUsecase(ticketRepo)
	ticketH = ticketHandler.NewTicketHandler(ticketUsecase)

	userRepo := userRepository.NewUserRepository()
	userUsecase := userUsecase.NewUserUsecase(userRepo)
	userH = userHandler.NewUserHandler(userUsecase)

	tdRepo := tdRepository.NewTransactionDetailRepository()
	tdUsecase := tdUsecase.NewTransactionDetailUsecase(tdRepo, userRepo)
	tdH = tdHandler.NewTransactionDetailHandler(tdUsecase)

	eventRepo := eventRepository.NewEventRepository()
	eventUsecase := eventUsecase.NewEventUsecase(eventRepo, userRepo, tdRepo)
	eventH = eventHandler.NewEventHandler(eventUsecase)

	return
}

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

	eventH, ticketH, tdH, userH := autowired()

	mux.Handle("/api/user/", routes.MuxUser(userH))
	mux.Handle("/api/event/", routes.MuxEvent(eventH))
	mux.Handle("/api/ticket/", routes.MuxTicket(ticketH))
	mux.Handle("/api/history/", routes.MuxTransactionDetail(tdH))

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
