package api

import (
	"fmt"
	"main/db"
	"main/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// productAPI.GET("/product", interceptor.VerifyIt, getProduct)
		productAPI.GET("/product" /*interceptor.JwtVerify,*/, getProduct)
	}

}

/*
// http://localhost:8081/api/v2/product
func getProduct(c *gin.Context) {
	var products []model.Product
	db.GetDB().Find(&products)
	c.JSON(http.StatusOK, products)
}
*/

func getProduct(c *gin.Context) {

	var products []model.Product
	if keyword := c.Query("keyword"); keyword != "" {
		keyword = fmt.Sprintf("%%%s%%", keyword)
		db.GetDB().Where("name like ?", keyword).Find(&products)
	} else {
		db.GetDB().Find(&products)
	}

	c.JSON(http.StatusOK, products)
}
