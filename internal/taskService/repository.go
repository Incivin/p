package taskService

import (
	"fmt"
	"gorm.io/gorm"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	UpdateTaskByID(id uint, task Task) (Task, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]Task, error)
}
type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}

}
func (r *taskRepository) GetTasksByUserID(userID uint) ([]Task, error) {
	var tasks []Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *taskRepository) CreateTask(task Task) (Task, error) {
	if task.UserID == 0 {
		return Task{}, fmt.Errorf("userID is required to create a task")
	}
	result := r.db.Create(&task)
	if result.Error != nil {
		return Task{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}
func (r *taskRepository) UpdateTaskByID(id uint, task Task) (Task, error) {
	var tasks Task
	if err := r.db.First(&tasks, id).Error; err != nil {
		return Task{}, err
	}
	if err := r.db.Model(&tasks).Updates(task).Error; err != nil {
		return Task{}, err
	}
	return tasks, nil

}
func (r *taskRepository) DeleteTaskByID(id uint) error {
	if err := r.db.Delete(&Task{}, id).Error; err != nil {
		return err
	}
	return nil
}
