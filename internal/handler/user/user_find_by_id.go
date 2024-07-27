package user

import "net/http"

// FindById implements UserHandler.
func (h *UserHandlerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	h.uc.FindById(r.Context(), 1)
}
