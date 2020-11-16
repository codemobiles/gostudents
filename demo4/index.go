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

	server.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		// http://localhost:85/login?username=lek&password=1234
		c.JSON(http.StatusOK, gin.H{"username": username, "password": password})
	})

	server.Run(":85")
}
