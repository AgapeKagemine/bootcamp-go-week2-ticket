package ticket

import "github.com/gin-gonic/gin"

// FindAll implements TicketHandler.
func (h *TicketHandlerImpl) FindAll(c *gin.Context) {
	h.uc.FindAll(c.Request.Context())
}
