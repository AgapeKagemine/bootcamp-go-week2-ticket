package user

import (
	"gotik/internal/domain"
	"net/http"
)

// Update implements UserHandler.
func (h *UserHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	h.uc.Update(r.Context(), &domain.User{})
}
