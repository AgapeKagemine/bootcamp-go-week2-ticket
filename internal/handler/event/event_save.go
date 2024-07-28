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

// Save implements EventHandler.
func (h *EventHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	// c := context.WithValue(r.Context(), domain.Start("method"), r.Method)
	ct := context.WithValue(r.Context(), domain.Start("start"), time.Now().Local())
	ctx, cancel := context.WithDeadline(ct, time.Now().Local().Add(time.Second*30))

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
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ctx.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("EVENT SAVE - %s", http.StatusText(int(response.StatusCode))))
	}()

	if r.Method != http.MethodGet {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	err := h.uc.Save(ctx, &domain.Event{})

	response.StatusCode = http.StatusOK

	if err != nil {
		response.Message = err.Error()
		return
	}

	response.Message = http.StatusText(http.StatusOK)
}
