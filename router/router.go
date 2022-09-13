package router

import (
	"gin_project/controller"

	"github.com/gin-gonic/gin"
)

func Routers(c *gin.Engine) {
	c.GET("/users", controller.GetUsers)
	c.POST("/users", controller.PostUser)
	c.GET("/users/:id", controller.GetByID)
	c.DELETE("/users/:id", controller.DeleteByid)
	c.PUT("/users", controller.UpdateUser)
}
