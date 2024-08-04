package contract

import "github.com/gin-gonic/gin"

type Save interface {
	Save(c *gin.Context)
}
