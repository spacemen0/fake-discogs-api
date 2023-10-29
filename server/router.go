package server

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("api/v1")
	{
		v1.GET("/hello-world", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Hello, world!",
			})
		})
	}
	return router
}
