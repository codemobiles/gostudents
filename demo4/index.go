package main

import (
	_ "fmt"
	"net/http"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "I am Root")
	})

	server.Run(":85")
}
