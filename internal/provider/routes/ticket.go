package routes

import (
	"encoding/json"
	"net/http"

	"gotik/internal/domain"
	"gotik/internal/handler/ticket"

	"github.com/rs/zerolog/log"
)

func MuxTicket(h ticket.TicketHandler) http.Handler {
	mux := http.NewServeMux()

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
