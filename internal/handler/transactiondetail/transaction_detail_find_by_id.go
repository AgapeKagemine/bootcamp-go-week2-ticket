package transactiondetail

import "github.com/gin-gonic/gin"

// FindById implements TransactionDetailHandler.
func (h *TransactionDetailHandlerImpl) FindById(c *gin.Context) {
	h.uc.FindById(c.Request.Context(), 1)
}
