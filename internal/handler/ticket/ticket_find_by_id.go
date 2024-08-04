package ticket

import "github.com/gin-gonic/gin"

// FindById implements TicketHandler.
func (h *TicketHandlerImpl) FindById(c *gin.Context) {
	h.uc.FindById(c.Request.Context(), 1)
}
