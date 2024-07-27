package user

import "net/http"

// DeleteById implements UserHanlder.
func (h *UserHandlerImpl) DeleteById(w http.ResponseWriter, r *http.Request) {
	h.uc.DeleteById(r.Context(), 1)
}
