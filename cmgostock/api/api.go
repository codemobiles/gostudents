package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Setup -- Call this method to setup gin router
func Setup(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hey")
	})
}
