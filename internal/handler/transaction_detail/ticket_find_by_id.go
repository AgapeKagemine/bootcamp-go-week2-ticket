package transaction_detail

import "net/http"

// FindById implements TransactionDetailHandler.
func (h TransactionDetailHandlerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	h.uc.FindById(r.Context(), 1)
}
