package routes

import (
	"github.com/gin-gonic/gin"
	"workshop/controllers"
	"workshop/middlewares"
)

func Router(r *gin.Engine) {
	r.GET("/", controllers.HelloWorld)
	r.GET("/item/template", controllers.ItemTemplate)
	r.POST("/item/display", controllers.ItemDisplay)
	r.GET("/item", controllers.GetItem)
	r.POST("/item", controllers.CreateItem)
	r.DELETE("/item", controllers.DeleteItem)
	r.PUT("/item", controllers.UpdateItem)
	r.POST("/items", controllers.CreateItems)

	r.Use(middlewares.NotFound)
}
