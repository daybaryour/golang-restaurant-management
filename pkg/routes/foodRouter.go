package routes

import (
	controller "github.com/daybaryour/golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func FoodRoutes(routes *gin.Engine) {
	routes.GET("/foods", controller.GetFoods)
	routes.GET("/foods/:id", controller.GetFood)
	routes.POST("/foods", controller.CreateFood)
	routes.PATCH("/foods/:id", controller.UpdateFood)
	routes.DELETE("/foods/:id", controller.DeleteFood)
}
