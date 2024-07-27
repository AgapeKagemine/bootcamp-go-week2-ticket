package event

import (
	"net/http"

	"gotik/internal/domain"
)

// Update implements EventHandler.
func (h *EventHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	h.uc.Update(r.Context(), &domain.Event{})
}
