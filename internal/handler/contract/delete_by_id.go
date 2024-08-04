package contract

import "github.com/gin-gonic/gin"

type DeleteById interface {
	DeleteById(c *gin.Context)
}
