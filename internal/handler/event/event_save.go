package event

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

// Save implements EventHandler.
func (h *EventHandlerImpl) Save(c *gin.Context) {
	// c := context.WithValue(c.Request.Context(), domain.Start("method"), c.Request.Method)
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
		log.Info().Uint("httpStatus", response.StatusCode).Str("statusDesc", response.Message).Str("processTime", time.Now().Local().Sub(ctx.Value(domain.Start("start")).(time.Time)).String()).Msg(fmt.Sprintf("EVENT SAVE - %s", http.StatusText(int(response.StatusCode))))
	}()

	if c.Request.Method != http.MethodGet {
		response.StatusCode = http.StatusMethodNotAllowed
		response.Message = http.StatusText(http.StatusMethodNotAllowed)
		return
	}

	event, err := h.uc.Save(ctx, &domain.Event{})
	response.Payload = event

	response.StatusCode = http.StatusOK

	if err != nil {
		response.Message = err.Error()
		return
	}

	response.Message = http.StatusText(http.StatusOK)
}
