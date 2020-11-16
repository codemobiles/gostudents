package interceptor

import (
	"github.com/gin-gonic/gin"
)

func verifyIt(c *gin.Context) {
	c.Next()
}
