package api

import (
	"main/db"

	"github.com/gin-gonic/gin"
)

func SetupTransactionAPI(router *gin.Engine) {
	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction" /* interceptor.JwtVerify,*/, getTransaction)
		//transactionAPI.POST("/transaction", interceptor.JwtVerify, createTransaction)
	}
}

func getTransaction(c *gin.Context) {
	var result []map[string]interface{}
	db.GetDB().Debug().Raw("SELECT transactions.id, total, paid, change, payment_type, payment_detail, order_list, users.username as Staff, transactions.created_at FROM transactions join users on transactions.staff_id = users.id order by transactions.created_at DESC", nil).Scan(&result)
	c.JSON(200, result)
}
