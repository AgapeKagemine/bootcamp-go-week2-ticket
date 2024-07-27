package user

import (
	"net/http"

	"gotik/internal/domain"
)

// Save implements UserHandler.
func (h *UserHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.User{})
}
