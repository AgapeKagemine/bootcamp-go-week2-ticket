package routes

import (
	"net/http"

	"gotik/internal/domain"
	"gotik/internal/handler/ticket"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func (r *Routes) Ticket(rg *gin.RouterGroup, h ticket.TicketHandler) {
	ticket := rg.Group("/ticket")

	// Just No
	ticket.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.WriteHeader(http.StatusNotFound)

		response := domain.ResponseBody{
			StatusCode: http.StatusNotFound,
			Message:    "Not Found",
			Payload:    nil,
		}

		log.Error().Msg("TICKET BEING HIT")

		c.JSON(http.StatusNotFound, response)
	})
}
