package user

import (
	"gotik/internal/domain"
	"net/http"
)

// Save implements UserHandler.
func (h *UserHandlerImpl) Save(w http.ResponseWriter, r *http.Request) {
	h.uc.Save(r.Context(), &domain.User{})
}
