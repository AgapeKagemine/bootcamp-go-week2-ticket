package event

import "net/http"

// FindAll implements EventHandler.
func (h *EventHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	h.uc.FindAll(r.Context())
}
