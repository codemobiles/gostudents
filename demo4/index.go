package main

import (
	_ "fmt"
	_ "net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.Run(":85")
}
