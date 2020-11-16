package api

import (
	"main/db"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
	var user model.User
	if e := c.ShouldBind(&user); e != nil {
		panic(e)
	}

	user.Password, _ = hashPassword(user.Password)
	if e := db.GetDB().Create(&user).Error; e != nil {
		panic(e)
	}

	c.String(http.StatusOK, "Register")
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
