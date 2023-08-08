package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/wandersonsantos01/go-gin/controllers"
)

func Setup() *gin.Engine {
	routes := gin.Default()
	return routes
}

func TestCheckGreetingStatusCodeResponse(t *testing.T) {
	r := Setup()
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
