package event

import "net/http"

// FindById implements EventHandler.
func (h *EventHandlerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	h.uc.FindById(r.Context(), 1)
}
