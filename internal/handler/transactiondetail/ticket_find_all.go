package transactiondetail

import "net/http"

// FindAll implements TransactionDetailHandler.
func (h *TransactionDetailHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	h.uc.FindAll(r.Context())
}
