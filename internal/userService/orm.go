package userService

import (
	"gorm.io/gorm"
	"p/internal/taskService"
)

type User struct {
	gorm.Model
	Email    string             `gorm:"unique;not null" json:"email"`
	Password string             `gorm:"not null" json:"password"`
	Tasks    []taskService.Task `gorm:"foreignKey:UserID" json:"tasks"`
}
