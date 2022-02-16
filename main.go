package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type HealthCheck struct {
	HealthStatus string `json:"health_status"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var users = []User{
	{1, "Muinde", 24},
}

func Index(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "UJUZI-WEB-SERVICES")
}

func healthyHandler(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, HealthCheck{"API is healthy..."})
}

func getAllUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var user User

	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	users = append(users, user)
	c.IndentedJSON(http.StatusCreated, users)
}

func getUserById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	for _, item := range users {
		if item.Id == id {
			c.IndentedJSON(http.StatusOK, item)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found!"})
}

func Router() *gin.Engine {
	router := gin.Default()
	router.GET("/", Index)
	router.GET("/health", healthyHandler)
	router.GET("/users", getAllUsers)
	router.GET("/users/:id", getUserById)
	router.POST("/users", createUser)

	return router
}

func main() {
	router := Router()

	router.Run("localhost:8080")

}
