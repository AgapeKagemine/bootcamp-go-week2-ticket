package routes

import (
	"net/http"

	tdHandler "gotik/internal/handler/transactiondetail"
	tdRepository "gotik/internal/repository/transactiondetail"
	tdUsecase "gotik/internal/usecase/transactiondetail"
)

func MuxTransactionDetail() http.Handler {
	mux := http.NewServeMux()

	repo := tdRepository.NewTransactionDetailRepository()
	uc := tdUsecase.NewTransactionDetailUsecase(repo)
	h := tdHandler.NewTransactionDetailHandler(uc)

	// 3. Melihat pesanan
	mux.HandleFunc("/", h.FindAll)

	// 5. Melihah keseluruhan pembeli tiket
	mux.HandleFunc("/buyer", h.FindAll)

	return http.StripPrefix("/api/history", mux)
}
