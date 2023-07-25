package main

import "github.com/gin-gonic/gin"

func ShowAllAnimals(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "Thor",
	})
}

func main() {
	r := gin.Default()
	r.GET("/animals", ShowAllAnimals)
	r.Run()
}
