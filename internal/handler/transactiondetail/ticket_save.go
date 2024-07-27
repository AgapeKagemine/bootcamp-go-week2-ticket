package transactiondetail

import (
	"net/http"

	"gotik/internal/domain"
)

// Save implements TransactionDetailHandler.
func (h *TransactionDetailHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.TransactionDetail{})
}
