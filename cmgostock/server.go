package main

import (
	_ "time"

	_ "github.com/gin-gonic/gin"
	_ "github.com/itsjamie/gin-cors"
)

func main() {
	router := gin.Default()

	router.Run(":8081")
}
