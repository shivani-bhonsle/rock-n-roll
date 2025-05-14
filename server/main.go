package main

import (
	"rock-n-roll/controllers"
	"rock-n-roll/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	database.ConnectDatabase()
	database.CreateStudentTable()

	router.POST("/student", controllers.InsertStudent)
	router.GET("/student", controllers.GetStudent)
	router.GET("/student/:id", controllers.GetStudentById)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("student/:id", controllers.DeleteStudent)

	router.GET("/ping", controllers.Ping)
	err := router.Run(":8080")
	if err != nil {
		panic("Error while starting the server!")
	}
}
