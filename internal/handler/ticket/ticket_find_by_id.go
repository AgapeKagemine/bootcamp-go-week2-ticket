package ticket

import "net/http"

// FindById implements TicketHandler.
func (h TicketHandlerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	h.uc.FindById(r.Context(), 1)
}
