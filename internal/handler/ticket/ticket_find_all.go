package ticket

import "net/http"

// FindAll implements TicketHandler.
func (h *TicketHandlerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	h.uc.FindAll(r.Context())
}
