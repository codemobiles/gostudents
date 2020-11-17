package api

import (
	"main/db"
	"main/interceptor"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// productAPI.GET("/product", interceptor.VerifyIt, getProduct)
		productAPI.GET("/product", interceptor.JwtVerify, getProduct)
	}

}

func getProduct(c *gin.Context) {
	var products []model.Product
	db.GetDB().Find(&products)
	c.JSON(http.StatusOK, products)
}
