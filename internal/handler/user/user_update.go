package user

import (
	"net/http"

	"gotik/internal/domain"
)

// Update implements UserHandler.
func (h *UserHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {
	h.uc.Update(r.Context(), &domain.User{})
}
