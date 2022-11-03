package userInputHttpAdapter

import (
	"context"
	"errors"
	"time"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
)

/* ADAPTERS TYPES */

type User struct {
	ID           string    `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	CreationDate time.Time `json:"creation-date"`
	LastUpdate   time.Time `json:"last-update"`
}

type CreateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type CreateUserResponse struct {
	User
}
type CreateUserResponseWrapper struct {
	CreateUserResponse CreateUserResponse
}

type DescribeUserRequest struct {
	ID string `json:"id"`
}
type DescribeUserResponse struct {
	User
}
type DescribeUserResponseWrapper struct {
	DescribeUserResponse DescribeUserResponse
}

type ListUserRequest struct{}

type ListUserResponse struct {
	Items []DescribeUserResponse `json:"items"`
}
type ListUserResponseWrapper struct {
	ListUserResponse ListUserResponse
}

type UpdateUserRequest struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type UpdateUserResponse struct {
	User
}
type UpdateUserResponseWrapper struct {
	UpdateUserResponse UpdateUserResponse
}

type DeleteUserRequest struct {
	ID string `json:"id"`
}
type DeleteUserResponse struct {
	User
}
type DeleteUserResponseWrapper struct {
	DeleteUserResponse DeleteUserResponse
}

/* INTERFACE ADAPTER */

type UserInputAdapter interface {
	AdaptCreateUser(ctx context.Context, data CreateUserRequest) (*CreateUserResponseWrapper, error)
	AdaptDescribeUser(ctx context.Context, data DescribeUserRequest) (*DescribeUserResponseWrapper, error)
	AdaptListUsers(ctx context.Context, data ListUserRequest) (*ListUserResponseWrapper, error)
	AdaptUpdateUser(ctx context.Context, data UpdateUserRequest) (*UpdateUserResponseWrapper, error)
	AdaptDeleteUser(ctx context.Context, data DeleteUserRequest) (*DeleteUserResponseWrapper, error)
}

/* DEFAULT USE CASE TYPES */

type DefaultHttpInputAdapter struct {
	userUseCase user.UserUseCase
}

type DefaultHttpInputAdapterOptions struct {
	UserUseCase user.UserUseCase
}

/* CONSTRUCTOR */

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.UserUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		userUseCase: options.UserUseCase,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ UserInputAdapter = (*DefaultHttpInputAdapter)(nil)

/* ADAPTERS IMPLEMENTATIONS */

func (adapter *DefaultHttpInputAdapter) AdaptCreateUser(ctx context.Context, data CreateUserRequest) (*CreateUserResponseWrapper, error) {

	// Map infra request to use case
	var inputData user.CreateUserInput
	inputData.ID = data.ID
	inputData.Username = data.Username
	inputData.Password = data.Password
	inputData.Email = data.Email

	//Use Case
	outputData, err := adapter.userUseCase.CreateUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response CreateUserResponse
	response.ID = outputData.ID
	response.Username = outputData.Username
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := CreateUserResponseWrapper{}
	responseWrapper.CreateUserResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeUser(ctx context.Context, data DescribeUserRequest) (*DescribeUserResponseWrapper, error) {
	// Map infra request to use case
	var inputData user.DescribeUserInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.userUseCase.DescribeUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DescribeUserResponse
	response.ID = outputData.ID
	response.Username = outputData.Username
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeUserResponseWrapper{}
	responseWrapper.DescribeUserResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListUsers(ctx context.Context, data ListUserRequest) (*ListUserResponseWrapper, error) {
	// Adapt payload input to use case input
	var inputData user.ListUserInput

	// Use case
	outputData, err := adapter.userUseCase.ListUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Adapt use case output to payload output
	var response ListUserResponse
	response.Items = make([]DescribeUserResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeUserResponse
		responseItem.ID = item.ID
		responseItem.Username = item.Username
		responseItem.Password = item.Password
		responseItem.Email = item.Email
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListUserResponseWrapper{}
	responseWrapper.ListUserResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateUser(ctx context.Context, data UpdateUserRequest) (*UpdateUserResponseWrapper, error) {
	// Map infra request to use case
	var inputData user.UpdateUserInput
	inputData.ID = data.ID
	inputData.Username = data.Username
	inputData.Password = data.Password
	inputData.Email = data.Email

	//Use Case
	outputData, err := adapter.userUseCase.UpdateUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response UpdateUserResponse
	response.ID = outputData.ID
	response.Username = outputData.Username
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := UpdateUserResponseWrapper{}
	responseWrapper.UpdateUserResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteUser(ctx context.Context, data DeleteUserRequest) (*DeleteUserResponseWrapper, error) {
	// Map infra request to use case
	var inputData user.RemoveUserInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.userUseCase.RemoveUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DeleteUserResponse
	response.ID = outputData.ID
	response.Username = outputData.Username
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteUserResponseWrapper{}
	responseWrapper.DeleteUserResponse = response

	return &responseWrapper, nil
}
