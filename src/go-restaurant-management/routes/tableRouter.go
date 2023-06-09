package routes

import (
	"github.com/gin-gonic.gin"
	controller "go-golang-restaurant-management/controllers"
)

func TableRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/tables", controller.GetTables())
	incomingRoutes.GET("/tables/:table_id", controller.GetTable())
	incomingRoutes.POST("/tables", controller.CreateTable())
	incomingRoutes.PATCH("/tables/:tables_id", controller.UpdateTable())
}