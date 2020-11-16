package api

import "github.com/gin-gonic/gin"


func SetupProductAPI(router *gin.Engine) {
	router.Group("/api/v2", getProduct)

}

func getProduct(c *gin.Context) {

}
