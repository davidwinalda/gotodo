package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Fullname string    `json:"fullname"`
	Email    string    `json:"email"`
	Task     []Task    `json:"tasks" gorm:"many2many:user_tasks;"`
	Project  []Project `json:"projects"`
	Label    []Label   `json:"labels"`
}

type UserTask struct {
	UserID int `gorm:"primaryKey"`
	TaskID int `gorm:"primaryKey"`
}

type Task struct {
	gorm.Model
	Title     string    `json:"title"`
	Schedule  time.Time `json:"schedule"`
	ProjectID int
	Label     []Label `gorm:"many2many:task_labels;"`
	StatusID  int     `gorm:"default:1"`
}

type TaskLabel struct {
	TaskID  int `gorm:"primaryKey"`
	LabelID int `gorm:"primaryKey"`
}

type Project struct {
	gorm.Model
	ProjectTitle string `json:"projectTitle"`
	UserID       int
}

type Label struct {
	gorm.Model
	LabelTitle string `json:"labelTitle"`
	UserID     int
}

type Status struct {
	gorm.Model
	StatusTitle string `json:"statusTitle"`
}
