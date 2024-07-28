package routes

import (
	"net/http"

	userHandler "gotik/internal/handler/user"
)

func MuxUser(h userHandler.UserHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/list", h.FindAll)
	mux.HandleFunc("/register", h.Save)

	return http.StripPrefix("/api/user", mux)
}
