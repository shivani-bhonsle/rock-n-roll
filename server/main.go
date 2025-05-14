package main

import (
	"rock-n-roll/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/student", controllers.InsertStudent)
	router.GET("/student", controllers.GetStudent)
	router.Run(":8080")
}
