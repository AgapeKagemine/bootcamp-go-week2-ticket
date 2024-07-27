package ticket

import (
	"gotik/internal/domain"
	"net/http"
)

// Save implements TicketHandler.
func (h TicketHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.Ticket{})
}
