package ticket

import (
	"net/http"

	"gotik/internal/domain"
)

// Save implements TicketHandler.
func (h *TicketHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.Ticket{})
}
