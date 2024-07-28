package routes

import (
	"net/http"

	tdHandler "gotik/internal/handler/transactiondetail"
)

func MuxTransactionDetail(h tdHandler.TransactionDetailHandler) http.Handler {
	mux := http.NewServeMux()

	// 3. Melihat pesanan
	// 5. Melihah keseluruhan pembeli tiket
	mux.HandleFunc("/", h.FindAll)

	return http.StripPrefix("/api/history", mux)
}
