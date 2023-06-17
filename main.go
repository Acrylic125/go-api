package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Dob  int64  `json:"date_of_birth"`
}

var users = []User{
	{Name: "John", Dob: 123},
	{Name: "Jane", Dob: 456},
	{Name: "Jack", Dob: 789},
}

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.Run("localhost:8080")
}
