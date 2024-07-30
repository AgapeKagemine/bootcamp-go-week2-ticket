package routes

import (
	"net/http"

	"gotik/internal/handler/event"
)

func MuxEvent(h event.EventHandler) http.Handler {
	mux := http.NewServeMux()

	// 1. Melihat daftar acara
	// 4. Melihat Keseluruhan Stok Tiket
	mux.HandleFunc("/", h.FindAll)

	// 2. Memesan Tiket
	mux.HandleFunc("/buy", h.BuyTicket)

	// Hardcoded
	mux.HandleFunc("/populate", h.Save)

	return http.StripPrefix("/api/event", mux)
}
