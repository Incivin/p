package handlers

import (
	"context"
	"gorm.io/gorm"
	"p/internal/userService"
	"p/internal/web/users"
)

type UserHandler struct {
	Service *userService.UserService
}

func NewUserHandler(service *userService.UserService) *UserHandler {
	return &UserHandler{service}
}

func (h *UserHandler) GetUsers(_ context.Context, _ users.GetUsersRequestObject) (users.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}
	response := users.GetUsers200JSONResponse{}
	for _, usr := range allUsers {
		userResponse := users.User{
			Email:    &usr.Email,
			Id:       &usr.ID,
			Password: &usr.Password,
		}
		response = append(response, userResponse)

	}
	return response, nil
}

func (h *UserHandler) PostUser(_ context.Context, request users.PostUserRequestObject) (users.PostUserResponseObject, error) {
	requestUser := request.Body
	CreateUser := userService.User{
		Email:    *requestUser.Email,
		Password: *requestUser.Password,
	}
	create, err := h.Service.CreateUser(CreateUser)
	if err != nil {
		return nil, err
	}
	response := users.PostUser201JSONResponse{
		Email:    &create.Email,
		Id:       &create.ID,
		Password: &create.Password,
	}
	return response, nil
}

func (h *UserHandler) DeleteUserByID(_ context.Context, request users.DeleteUserByIDRequestObject) (users.DeleteUserByIDResponseObject, error) {
	requestDelete := request.Id
	err := h.Service.DeleteUser(requestDelete)
	if err != nil {
		return nil, err
	}
	return users.DeleteUserByID204Response{}, nil
}

func (h *UserHandler) PatchUserByID(_ context.Context, request users.PatchUserByIDRequestObject) (users.PatchUserByIDResponseObject, error) {
	userid := request.Id
	updateRequest := request.Body
	userToUpdate := userService.User{
		Model: gorm.Model{
			ID: userid,
		},
	}
	if updateRequest.Email != nil {
		userToUpdate.Email = *updateRequest.Email
	}
	if updateRequest.Password != nil {
		userToUpdate.Password = *updateRequest.Password
	}
	updatedUser, err := h.Service.UpdateUser(userid, userToUpdate)
	if err != nil {
		return nil, err
	}

	response := users.PatchUserByID200JSONResponse{
		Email:    &updatedUser.Email,
		Id:       &updatedUser.ID,
		Password: &updatedUser.Password,
	}
	return response, nil
}
