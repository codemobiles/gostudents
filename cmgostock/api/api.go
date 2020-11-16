package api

import (
	"github.com/gin-gonic/gin"
)

// Setup -- Call this method to setup gin router
func Setup(router *gin.Engine) {
	SetupAuthenAPI(router)
	SetupProductAPI(router)
	SetupTransactionAPI(router)
}
