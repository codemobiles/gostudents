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
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": e})
		return
	}

	var queryUser model.User
	if e := db.GetDB().First(&queryUser, "username = ?", user.Username).Error; e != nil {
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": "invalid username"})
		return
	}

	if checkPasswordHash(user.Password, queryUser.Password) {
		c.JSON(200, gin.H{"result": "ok", "token": "1234"})
	} else {
		c.JSON(http.StatusOK, gin.H{"result": "nok", "error": "invalid password"})
	}
}

func register(c *gin.Context) {
	var user model.User
	if e := c.ShouldBind(&user); e != nil {
		// panic(e)
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": e})
	}

	user.Password, _ = hashPassword(user.Password)
	if e := db.GetDB().Create(&user).Error; e != nil {
		// panic(e)
		// fmt.Println(e)
		c.JSON(http.StatusBadRequest, gin.H{"result": "nok", "error": e})
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok", "message": user})
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
