package controllers

import (
	"net/http"
	"todolist/config"
	"todolist/models"

	"github.com/gin-gonic/gin"
)

func GetAllProject(c *gin.Context) {
	var project []models.Project

	if err := config.DB.Find(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, project)
}

func CreateProject(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	var project models.Project

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&project)

	if err := config.DB.Omit("User").Create(&project).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, project)
}

func GetAllLabel(c *gin.Context) {
	var label []models.Label

	if err := config.DB.Find(&label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, label)
}

func GetLabelByUserId(c *gin.Context) {
	id := c.Params.ByName("id")

	var user models.User

	if err := config.DB.Preload("Label").First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateLabel(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	var label models.Label

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&label)

	if err := config.DB.Omit("User").Create(&label).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, label)
}

func GetAllTask(c *gin.Context) {
	var task []models.Task

	if err := config.DB.Find(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, task)
}

func GetTasksByUserId(c *gin.Context) {
	id := c.Params.ByName("id")

	var user models.User

	if err := config.DB.Preload("Project").Preload("Label").Preload("Task").First(&user, id).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateTask(c *gin.Context) {
	id := c.Params.ByName("id")
	var user models.User
	var task models.Task

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.ShouldBindJSON(&task)

	if err := config.DB.Omit("User", "Status", "Label").Create(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record can't create!"})
		return
	}

	c.JSON(http.StatusOK, task)
}
