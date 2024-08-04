package transactiondetail

import (
	"gotik/internal/domain"

	"github.com/gin-gonic/gin"
)

// Save implements TransactionDetailHandler.
func (h *TransactionDetailHandlerImpl) Save(c *gin.Context) {
	h.uc.Save(c.Request.Context(), &domain.TransactionDetail{})
}
