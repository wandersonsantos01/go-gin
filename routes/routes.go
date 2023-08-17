package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wandersonsantos01/go-gin/controllers"
)

func HandleRequests() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/animals", controllers.ShowAllAnimals)
	r.GET("/animals/:id", controllers.GetAnimalById)
	r.GET("/:name", controllers.Greeting)
	r.POST("/animals", controllers.InsertAnimal)
	r.DELETE("/animals/:id", controllers.DeleteAnimal)
	r.PUT("/animals/:id", controllers.UpdateAnimal)
	r.GET("/animals/nickname/:nickname", controllers.GetAnimalByNickname)
	r.GET("/index", controllers.ShowIndexPage)
	
	r.NoRoute(controllers.RouteNotFound)

	r.Run()
}
