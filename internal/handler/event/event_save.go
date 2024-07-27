package event

import (
	"net/http"

	"gotik/internal/domain"
)

// Save implements EventHandler.
func (h *EventHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.Event{})
}
