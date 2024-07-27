package routes

import (
	"net/http"

	eventHandler "gotik/internal/handler/event"
	eventRepository "gotik/internal/repository/event"
	eventUsecase "gotik/internal/usecase/event"
)

func MuxEvent() http.Handler {
	mux := http.NewServeMux()

	repo := eventRepository.NewEventRepository()
	uc := eventUsecase.NewEventUsecase(repo)
	h := eventHandler.NewEventHandler(uc)

	// 1. Melihat daftar acara
	mux.HandleFunc("/", h.FindAll)

	// 4. Melihat Keseluruhan Stok Tiket
	mux.HandleFunc("/{id}", h.FindById)

	// Hardcoded
	mux.HandleFunc("/populate", h.Save)

	return http.StripPrefix("/api/event", mux)
}
