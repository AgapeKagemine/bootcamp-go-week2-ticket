package event

import (
	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
)

// Update implements EventHandler.
func (h *EventHandlerImpl) Update(c *gin.Context) {
	h.uc.Update(c.Request.Context(), &domain.Event{})
}
