package server

import (
	"NewApp/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	v1 := router.Group("api/v1")
	{
		v1.GET("get-record/:id", controllers.GetRecordByID)
		v1.POST("/create-record", controllers.CreateRecord)
		v1.POST("/get-records", controllers.GetAllRecords)
		v1.POST("/get-records-by-seller-id/:id", controllers.GetRecordsBySellerID)
		v1.PUT("/update-record/:id", controllers.UpdateRecord)
		v1.DELETE("delete-record/:id", controllers.DeleteRecord)
	}
	{
		v1.GET("get-user/:id", controllers.GetUserByID)
		v1.GET("get-user-by-username/:username", controllers.GetUserByUsername)
		v1.POST("/create-user", controllers.CreateUser)
		v1.PUT("/update-user/:id", controllers.UpdateUser)
		v1.DELETE("delete-user/:id", controllers.DeleteUser)
	}
	return router
}
