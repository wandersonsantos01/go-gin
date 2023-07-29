package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wandersonsantos01/go-gin/databases"
	"github.com/wandersonsantos01/go-gin/models"
)

func ShowAllAnimals(c *gin.Context) {
	var animals []models.Animal
	databases.DB.Find(&animals)
	c.JSON(200, animals)
}

func Greeting(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "Hello " + name,
	})
}

func InsertAnimal(c *gin.Context) {
	var animal models.Animal
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	databases.DB.Create(&animal)
	c.JSON(http.StatusOK, animal)
}

func GetAnimalById(c *gin.Context) {
	id := c.Params.ByName("id")
	var animal models.Animal
	databases.DB.First(&animal, id)
	if animal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Animal not found",
		})
		return
	}

	c.JSON(http.StatusOK, animal)
}

func DeleteAnimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var animal models.Animal
	databases.DB.Delete(&animal, id)
	c.JSON(http.StatusOK, gin.H{
		"message": "Animal ID " + id + " deleted",
	})
}

func UpdateAnimal(c *gin.Context) {
	id := c.Params.ByName("id")
	var animal models.Animal

	databases.DB.First(&animal, id)
	if err := c.ShouldBindJSON(&animal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	databases.DB.Model(&animal).UpdateColumns(animal)
	c.JSON(http.StatusOK, animal)
}

func GetAnimalByNickname(c *gin.Context) {
	nickname := c.Param("nickname")
	var animal models.Animal
	databases.DB.Where(&models.Animal{Nickname: nickname}).First(&animal)
	if animal.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Animal not found",
		})
		return
	}
	c.JSON(http.StatusOK, animal)
}
