package routes

import (
	"net/http"

	userHandler "gotik/internal/handler/user"
	userRepository "gotik/internal/repository/user"
	userUsecase "gotik/internal/usecase/user"
)

func MuxUser() http.Handler {
	mux := http.NewServeMux()

	repo := userRepository.NewUserRepository()
	uc := userUsecase.NewUserUsecase(repo)
	h := userHandler.NewUserHandler(uc)

	mux.HandleFunc("/list", h.FindAll)
	mux.HandleFunc("/register", h.Save)

	return http.StripPrefix("/api/user", mux)
}
