package user

import "github.com/gin-gonic/gin"

// FindById implements UserHandler.
func (h *UserHandlerImpl) FindById(c *gin.Context) {
	h.uc.FindById(c.Request.Context(), 1)
}
