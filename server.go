package main

import (
	"fmt"
	"gin_project/router"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	fmt.Println("hasbeen started ...")
	router.Routers(server)
	server.Run("localhost:2200")
}
