package routes

import (
	"golang/final/controller"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.GET("/", controller.Show)
	router.POST("/add", controller.Add)
	router.GET("/delete/:id", controller.Delete)
	router.GET("/complete/:id", controller.Complete)
	router.POST("/update/:id", controller.Update)
	return router
}
