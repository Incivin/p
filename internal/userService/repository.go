package userService

import (
	"gorm.io/gorm"
	"p/internal/taskService"
)

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(id uint, user User) (User, error)
	DeleteUser(id uint) error
	GetTasksForUser(userID uint) ([]taskService.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	var tasks []taskService.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetAllUsers() ([]User, error) {
	var user []User
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) UpdateUser(id uint, updatedUser User) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return User{}, err
	}
	if err := r.db.Model(&user).Updates(updatedUser).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		return err
	}
	return nil
}
