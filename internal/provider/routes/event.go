package routes

import (
	"gotik/internal/handler/event"

	"github.com/gin-gonic/gin"
)

func (r *Routes) Event(rg *gin.RouterGroup, h event.EventHandler) {
	event := rg.Group("/event")

	// 1. Melihat daftar acara
	// 4. Melihat Keseluruhan Stok Tiket
	event.GET("/", h.FindAll)

	// 2. Memesan Tiket
	event.POST("/buy", h.BuyTicket)

	// Hardcoded
	event.GET("/populate", h.Save)
}
