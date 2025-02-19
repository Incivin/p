package handlers

import (
	"context"
	"gorm.io/gorm"
	"p/internal/taskService"
	"p/internal/web/tasks"
)

type Handler struct {
	Service *taskService.TaskService
}

func (h *Handler) DeleteTasksId(_ context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id
	err := h.Service.DeleteTask(taskID)
	if err != nil {
		return nil, err
	}
	return tasks.DeleteTasksId204Response{}, nil
}

func (h *Handler) PatchTasksId(_ context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	updateRequest := request.Body
	taskToUpdate := taskService.Task{
		Model: gorm.Model{
			ID: taskID,
		},
	}
	if updateRequest.Task != nil {
		taskToUpdate.Task = *updateRequest.Task
	}
	if updateRequest.IsDone != nil {
		taskToUpdate.IsDone = *updateRequest.IsDone
	}

	updateTask, err := h.Service.PatchTask(taskID, taskToUpdate)
	if err != nil {
		return nil, err
	}
	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updateTask.ID,
		IsDone: &updateTask.IsDone,
		Task:   &updateTask.Task,
	}
	return response, nil

}

func (h *Handler) GetTasks(_ context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	//Получение всех задач из сервиса
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}
	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			IsDone: &tsk.IsDone,
			Task:   &tsk.Task,
		}
		response = append(response, task)
	}
	return response, nil
}

func (h *Handler) PostTasks(_ context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body
	taskToCreate := taskService.Task{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}
	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		IsDone: &createdTask.IsDone,
		Task:   &createdTask.Task,
	}
	return response, nil
}

func NewHandler(service *taskService.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}
