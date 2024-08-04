package contract

import "github.com/gin-gonic/gin"

type FindAll interface {
	FindAll(c *gin.Context)
}
