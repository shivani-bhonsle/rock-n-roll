package main

import (
	"rock-n-roll/controllers"
	"rock-n-roll/database"
	"sync"

	"github.com/gin-gonic/gin"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	router := gin.Default()

	database.ConnectDatabase()
	database.CreateStudentTable()

	router.POST("/student", controllers.InsertStudent)
	router.GET("/student", controllers.GetStudent)
	router.GET("/student/:id", controllers.GetStudentById)
	router.PUT("/student/:id", controllers.UpdateStudent)
	router.DELETE("student/:id", controllers.DeleteStudent)

	router.GET("photos", controllers.GetPhotos)

	wg.Add(2)
	go controllers.FetchPhotosRoutine(&wg, ch)
	go controllers.PrintPhotoId(&wg, ch)
	wg.Wait()

	router.GET("/ping", controllers.Ping)
	err := router.Run(":8080")
	if err != nil {
		panic("Error while starting the server!")
	}
}
