package routes

import (
	"encoding/json"
	"net/http"

	"gotik/internal/domain"
	ticketHandler "gotik/internal/handler/ticket"
	ticketRepository "gotik/internal/repository/ticket"
	ticketUsecase "gotik/internal/usecase/ticket"

	"github.com/rs/zerolog/log"
)

func MuxTicket() http.Handler {
	mux := http.NewServeMux()

	repo := ticketRepository.NewTicketRepository()
	uc := ticketUsecase.NewTicketUsecase(repo)
	_ = ticketHandler.NewTicketHandler(uc)

	// List Ticket
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)

		response := domain.ResponseBody{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found",
			Payload:    nil,
		}

		log.Error().Msg("TICKET BEING HIT")

		json.NewEncoder(w).Encode(response)
	})

	return http.StripPrefix("/api/ticket", mux)
}
