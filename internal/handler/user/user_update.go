package user

import (
	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
)

// Update implements UserHandler.
func (h *UserHandlerImpl) Update(c *gin.Context) {
	h.uc.Update(c.Request.Context(), &domain.User{})
}
