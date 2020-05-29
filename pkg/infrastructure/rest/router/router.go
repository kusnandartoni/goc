package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(cors.Default())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world")
	})

	return r
}
