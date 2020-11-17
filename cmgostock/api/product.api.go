package api

import (
	"fmt"
	"main/db"
	"main/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupProductAPI(router *gin.Engine) {
	productAPI := router.Group("/api/v2")
	{
		// productAPI.GET("/product", interceptor.VerifyIt, getProduct)
		productAPI.GET("/product" /*interceptor.JwtVerify,*/, getProduct)
		productAPI.GET("/product/:id" /*interceptor.JwtVerify,*/, getProductByID)
		productAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)

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
		db.GetDB().Where("name like ?", keyword).Order("created_at DESC").Find(&products)
	} else {
		db.GetDB().Order("created_at DESC").Find(&products)
	}

	c.JSON(http.StatusOK, products)
}

func getProductByID(c *gin.Context) {
	var product model.Product
	db.GetDB().Where("id = ?", c.Param("id")).First(&product)
	c.JSON(200, product)
}

func createProduct(c *gin.Context) {
	product := model.Product{}
	product.Name = c.PostForm("name")
	product.Stock, _ = strconv.ParseInt(c.PostForm("stock"), 10, 64)
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.CreatedAt = time.Now()

}
