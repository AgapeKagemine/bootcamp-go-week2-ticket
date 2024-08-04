package ticket

import (
	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
)

// Save implements TicketHandler.
func (h *TicketHandlerImpl) Save(c *gin.Context) {
	h.uc.Save(c.Request.Context(), &domain.Ticket{})
}
