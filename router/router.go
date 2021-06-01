package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zosma180/go-poc/controller"
)

func Create() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	v1.GET("/user", controller.GetUserList)
	v1.POST("/user", controller.CreateUser)

	return router
}
