package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wandersonsantos01/go-gin/controllers"
	"github.com/wandersonsantos01/go-gin/databases"
	"github.com/wandersonsantos01/go-gin/models"
)

var ID int

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func MockAnimal() {
	animal := models.Animal{Name: "Mock Animal", Age: 11, Nickname: "MockAn"}
	databases.DB.Create(&animal)
	ID = int(animal.ID)
}

func DeleteMockAnimal() {
	var animal models.Animal
	databases.DB.Delete(&animal, ID)
}

func TestCheckGreetingStatusCodeResponse(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/:name", controllers.Greeting)
	request, _ := http.NewRequest("GET", "/wan", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	// if response.Code != http.StatusOK {
	// 	t.Fatalf("Status code error - Expected: %d | Given: %d", http.StatusOK, response.Code)
	// }
	assert.Equal(t, http.StatusOK, response.Code)

	mockResponse := `{"message":"Hello wan"}`
	body, _ := ioutil.ReadAll(response.Body)
	assert.Equal(t, mockResponse, string(body))
}

func TestListingAllAnimalsHandler(t *testing.T) {
	databases.DbConnect()
	MockAnimal()
	defer DeleteMockAnimal()
	r := SetupTestRoutes()
	r.GET("/animals", controllers.ShowAllAnimals)
	request, _ := http.NewRequest("GET", "/animals", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}

func TestGetAnimalByNickname(t *testing.T) {
	databases.DbConnect()
	MockAnimal()
	defer DeleteMockAnimal()
	r := SetupTestRoutes()
	r.GET("/animals/nickname/:nickname", controllers.GetAnimalByNickname)
	request, _ := http.NewRequest("GET", "/animals/nickname/MockAn", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
