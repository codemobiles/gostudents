package api

import (
	"main/model"
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

	var user model.User
	if e := c.ShouldBind(&user); e != nil {
		// fmt.Println(e)
		panic(e)
	}

	c.JSON(http.StatusOK, user)

}

func register(c *gin.Context) {
	c.String(http.StatusOK, "Register")
}
