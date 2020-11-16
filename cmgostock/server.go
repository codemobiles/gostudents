package main

import (
	"main/api"
	_ "time"

	"github.com/gin-gonic/gin"
	_ "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()
	api.Setup(router)

	router.Run(":8081")
}
