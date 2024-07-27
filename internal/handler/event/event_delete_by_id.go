package event

import (
	"net/http"
)

// DeleteById implements EventHandler.
func (h *EventHandlerImpl) DeleteById(w http.ResponseWriter, r *http.Request) {
	h.uc.DeleteById(r.Context(), 1)
}
