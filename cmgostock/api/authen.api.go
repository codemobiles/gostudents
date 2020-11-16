package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetupAuthenAPI - Call this method to setup authen api
func SetupAuthenAPI(router *gin.Engine) {
	authenAPI := router.Group("/api/v2")
	authenAPI.POST("/login", login)
	authenAPI.POST("/register", register)
}

func login(c *gin.Context) {
	c.String(http.StatusOK, "Login")
}

func register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}
