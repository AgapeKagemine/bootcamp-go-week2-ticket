package event

import "github.com/gin-gonic/gin"

// DeleteById implements EventHandler.
func (h *EventHandlerImpl) DeleteById(c *gin.Context) {
	h.uc.DeleteById(c.Request.Context(), 1)
}
