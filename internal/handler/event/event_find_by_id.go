package event

import "github.com/gin-gonic/gin"

// FindById implements EventHandler.
func (h *EventHandlerImpl) FindById(c *gin.Context) {
	h.uc.FindById(c.Request.Context(), 1)
}
