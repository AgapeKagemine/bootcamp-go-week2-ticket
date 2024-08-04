package contract

import "github.com/gin-gonic/gin"

type FindById interface {
	FindById(c *gin.Context)
}
