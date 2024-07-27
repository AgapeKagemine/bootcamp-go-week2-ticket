package event

import (
	"gotik/internal/domain"
	"net/http"
)

// Save implements EventHandler.
func (h *EventHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.Event{})
}
