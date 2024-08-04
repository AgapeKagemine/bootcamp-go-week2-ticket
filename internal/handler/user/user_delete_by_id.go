package user

import "github.com/gin-gonic/gin"

// DeleteById implements UserHanlder.
func (h *UserHandlerImpl) DeleteById(c *gin.Context) {
	h.uc.DeleteById(c.Request.Context(), 1)
}
