package routes

import (
	controller "golang-go-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/users", controller.GetUsers)
	incomingRoutes.GET("/users/:user_id", controller.GetUser)
	incomingRoutes.POST("/users/signup", controller.SignUp)
	incomingRoutes.POST("/users/login", controller.Login)
}
