package event

import (
	"gotik/internal/domain"
	"net/http"
)

// Update implements EventHandler.
func (h *EventHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	h.uc.Update(r.Context(), &domain.Event{})
}
