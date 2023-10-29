package server

import (
	"NewApp/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.POST("/create-record", controllers.CreateRecord)
		v1.POST("/get-records", controllers.GetAllRecords)
	}
	return router
}
