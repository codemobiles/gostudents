package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

// Setup -- Call this method to setup gin router
func Setup(router *gin.Engine) {

	db.SetupDB()
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}
