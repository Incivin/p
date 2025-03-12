package userService

import "p/internal/taskService"

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}
func (s *UserService) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}
func (s *UserService) UpdateUser(id uint, user User) (User, error) {
	return s.repo.UpdateUser(id, user)
}
func (s *UserService) DeleteUser(id uint) error {
	return s.repo.DeleteUser(id)
}

func (s *UserService) GetTasksForUser(userID uint) ([]taskService.Task, error) {
	return s.repo.GetTasksForUser(userID)
}
