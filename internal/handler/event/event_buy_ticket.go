package event

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type BuyTicket interface {
	BuyTicket(c *gin.Context)
}

func (h *EventHandlerImpl) BuyTicket(c *gin.Context) {
	ct := context.WithValue(c.Request.Context(), domain.Start("start"), time.Now().Local())
	ctx, cancel := context.WithDeadline(ct, time.Now().Local().Add(time.Second*30))

	c.Writer.Header().Set("Content-Type", "application/json")

	response := &domain.ResponseBody{
		StatusCode: 0,
		Message:    "",
		Payload:    nil,
	}

	defer func() {
		c.JSON(int(response.StatusCode), response)
		c.Request.Body.Close()
		cancel()
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ct.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("EVENT BUY TICKET - %s", http.StatusText(int(response.StatusCode))))
	}()

	if c.Request.Method != http.MethodPost {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	// if len(c.Request.Header.Values("Content-Type")) == 0 {
	// 	response.StatusCode = http.StatusBadRequest
	// 	response.Message = "Bad request - Content-type is not set"
	// 	return
	// }

	// if c.Request.Header.Values("Content-Type")[0] != "application/json" {
	// 	response.StatusCode = http.StatusUnsupportedMediaType
	// 	response.Message = http.StatusText(http.StatusUnsupportedMediaType)
	// 	return
	// }

	response.StatusCode = http.StatusBadRequest

	if c.Request.Body == nil {
		response.Message = "Bad request - Body is empty"
		return
	}

	var request domain.EventBuyTicket
	err := json.NewDecoder(c.Request.Body).Decode(&request)
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

	ctxR := context.WithValue(ctx, domain.Start("request"), &request)

	td, err := h.uc.BuyTicket(ctxR)

	response.Payload = td

	if err != nil {
		response.Message = err.Error()
		return
	}

	response.StatusCode = http.StatusOK
	response.Message = http.StatusText(http.StatusOK)
}
