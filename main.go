package main

import (
	"todolist/config"
	"todolist/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	config.ConnectDatabase()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/project/all", controllers.GetAllProject)
		v1.POST("/project/:id", controllers.CreateProject)
		v1.GET("/label/all", controllers.GetAllLabel)
		v1.POST("/label/:id", controllers.CreateLabel)
		v1.GET("/label/user/:id", controllers.GetLabelByUserId)
		v1.GET("/task/all", controllers.GetAllTask)
		v1.POST("/task/:id", controllers.CreateTask)
		v1.GET("/task/user/:id", controllers.GetTasksByUserId)
	}

	r.Run()
}
