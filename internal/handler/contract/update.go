package contract

import "github.com/gin-gonic/gin"

type Update interface {
	Update(c *gin.Context)
}
