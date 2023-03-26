package routes

import (
	"github.com/daybaryour/golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(routes *gin.Engine) {
	routes.GET("/users", controllers.GetUsers)
	routes.GET("/users/:id", controllers.GetUser)
	routes.POST("/users", controllers.CreateUser)
	routes.POST("/users/login", controllers.Login)
	routes.POST("/users/signup", controllers.Logout)
	routes.PUT("/users/:id", controllers.UpdateUser)
	routes.DELETE("/users/:id", controllers.DeleteUser)
}
