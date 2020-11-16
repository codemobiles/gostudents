package interceptor

import (
	"github.com/gin-gonic/gin"
)

func VerifyIt(c *gin.Context) {
	c.Next()
}
