package server

import (
	"fake-discogs-api/controllers"
	"fake-discogs-api/middlewares"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())
	v1 := router.Group("api/v1")
	{
		v1.GET("get-user-by-username/:username", controllers.GetUserByUsername)
		v1.GET("get-users-by-username/:username", controllers.GetUsersByUsername)
		v1.POST("/user-register", controllers.CreateUser)
		v1.POST("user-login/:option", controllers.UserLogin)
		v1.GET("get-record/:id", controllers.GetRecordByID)
		v1.POST("/get-records", controllers.GetAllRecords)
		v1.POST("/get-records-by-seller-name/:name", controllers.GetRecordsBySellerName)
		v1.POST("search-records", controllers.SearchRecordsWithPagination)
	}
	v1.Use(middlewares.AuthMiddleware())
	{
		v1.PUT("/update-record/:id", controllers.UpdateRecord)
		v1.DELETE("delete-record/:id", controllers.DeleteRecord)
		v1.POST("/create-record", controllers.CreateRecord)
		v1.POST("/create-image/:id", controllers.CreateImage)
		v1.GET("get-user", controllers.GetUserByID)
		v1.PUT("/update-user", controllers.UpdateUser)
		v1.DELETE("delete-user", controllers.DeleteUser)
	}
	return router
}
