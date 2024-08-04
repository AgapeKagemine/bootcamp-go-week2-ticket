package routes

import (
	"gotik/internal/handler/user"

	"github.com/gin-gonic/gin"
)

func (r *Routes) User(rg *gin.RouterGroup, h user.UserHandler) {
	user := rg.Group("/user")

	user.GET("/", h.FindAll)
	user.POST("/register", h.Save)

	// TODO: If done with core...
	// h.DeleteById()
	// h.FindById()
	// h.Update()
}
