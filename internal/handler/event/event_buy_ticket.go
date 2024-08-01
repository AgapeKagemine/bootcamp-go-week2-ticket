package event

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gotik/internal/domain"

	"github.com/rs/zerolog/log"
)

type BuyTicket interface {
	BuyTicket(w http.ResponseWriter, r *http.Request)
}

func (h *EventHandlerImpl) BuyTicket(w http.ResponseWriter, r *http.Request) {
	c := context.WithValue(r.Context(), domain.Start("start"), time.Now().Local())
	ct, cancel := context.WithDeadline(c, time.Now().Local().Add(time.Second*30))

	w.Header().Set("Content-Type", "application/json")

	response := &domain.ResponseBody{
		StatusCode: 0,
		Message:    "",
		Payload:    nil,
	}

	defer func() {
		w.WriteHeader(int(response.StatusCode))
		json.NewEncoder(w).Encode(response)
		r.Body.Close()
		cancel()
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ct.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("EVENT BUY TICKET - %s", http.StatusText(int(response.StatusCode))))
	}()

	if r.Method != http.MethodPost {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	// if len(r.Header.Values("Content-Type")) == 0 {
	// 	response.StatusCode = http.StatusBadRequest
	// 	response.Message = "Bad request - Content-type is not set"
	// 	return
	// }

	// if r.Header.Values("Content-Type")[0] != "application/json" {
	// 	response.StatusCode = http.StatusUnsupportedMediaType
	// 	response.Message = http.StatusText(http.StatusUnsupportedMediaType)
	// 	return
	// }

	response.StatusCode = http.StatusBadRequest

	if r.Body == nil {
		response.Message = "Bad request - Body is empty"
		return
	}

	var request domain.EventBuyTicket
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.Message = "Bad request - Bad JSON format"
		return
	}

	response.Payload = request

	if request.EventId == nil || *request.EventId < 1 {
		response.Message = "Buy Ticket: Invalid EventId Request"
		return
	}

	if request.UserId == nil || *request.UserId < 1 {
		response.Message = "Select Event: Invalid UserId Request"
		return
	}

	if request.Ticket == nil || len(*request.Ticket) == 0 {
		response.Message = "Select Event: Invalid Ticket Request"
		return
	}

	for i, t := range *request.Ticket {
		if t == nil || t.Quantity == nil || t.TicketId == nil {
			response.Message = fmt.Sprintf("Select Ticket: Invalid Ticket at index %d", i+1)
			return
		}

		if *t.TicketId < 1 {
			response.Message = fmt.Sprintf("Select Ticket: Invalid TicketId at index %d", i+1)
			return
		}

		if *t.Quantity < 1 {
			response.Message = fmt.Sprintf("Select Ticket: Invalid Quantity at index %d", i+1)
			return
		}
	}

	ctx := context.WithValue(ct, domain.Start("request"), &request)

	td, err := h.uc.BuyTicket(ctx)

	response.Payload = td

	if err != nil {
		response.Message = err.Error()
		return
	}

	response.StatusCode = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
}
