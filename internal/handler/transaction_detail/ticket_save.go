package transaction_detail

import (
	"gotik/internal/domain"
	"net/http"
)

// Save implements TransactionDetailHandler.
func (h TransactionDetailHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.TransactionDetail{})
}
