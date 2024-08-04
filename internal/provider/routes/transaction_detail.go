package routes

import (
	"gotik/internal/handler/transactiondetail"

	"github.com/gin-gonic/gin"
)

func (r *Routes) TransactionDetail(rg *gin.RouterGroup, h transactiondetail.TransactionDetailHandler) {
	td := rg.Group("/history")

	// 3. Melihat pesanan
	// 5. Melihah keseluruhan pembeli tiket
	td.GET("/", h.FindAll)
	td.GET("/all", h.FindAll)
}
