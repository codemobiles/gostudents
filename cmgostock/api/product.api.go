package api

import (
	"main/interceptor"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", interceptor.VerifyIt, getProduct)
	}

}

func getProduct(c *gin.Context) {
	c.JSON(http.StatusOK, []int{1, 2, 3, 4})
}
