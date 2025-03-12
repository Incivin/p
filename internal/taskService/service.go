package taskService

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}
func (s *TaskService) CreateTask(task Task, userID uint) (Task, error) {
	task.UserID = userID
	return s.repo.CreateTask(task)
}
func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()
}
func (s *TaskService) PatchTask(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}
func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTaskByID(id)
}
func (s *TaskService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
