package user

import "net/http"

// FindAll implements UserHandler.
func (h *UserHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	h.uc.FindAll(r.Context())
}
