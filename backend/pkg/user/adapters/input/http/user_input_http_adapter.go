package userInputHttpAdapter

import (
	"context"
	"errors"

	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
)

type UserInputAdapter interface {
	AdaptDescribeUser(ctx context.Context, data DescribeUserRequest) (*DescribeUserResponseWrapper, error)
	AdaptListUsers(ctx context.Context, data ListUserRequest) (*ListUserResponseWrapper, error)
	AdaptCreateUser(ctx context.Context, data CreateUserRequest) (*CreateUserResponseWrapper, error)
	AdaptLoginUser(ctx context.Context, data LoginUserRequest) (*LoginUserResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.UserUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		userUseCase: options.UserUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	userUseCase user.UserUseCase
}

type DefaultHttpInputAdapterOptions struct {
	UserUseCase user.UserUseCase
}

var _ UserInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptDescribeUser(ctx context.Context, data DescribeUserRequest) (*DescribeUserResponseWrapper, error) {
	var inputData user.DescribeUserInput
	inputData.Project = data.Project
	inputData.ID = data.ID

	outputData, err := adapter.userUseCase.DescribeUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeUserResponse
	response.Project = outputData.User.Project
	response.ID = outputData.User.ID
	response.ContractID = outputData.User.ContractID
	response.Name = outputData.User.Name
	response.Email = outputData.User.Email
	response.Role = outputData.User.Role
	response.EthAddress = outputData.User.EthAddress
	response.DefaultPasswordChanged = outputData.User.DefaultPasswordChanged
	response.Network = outputData.User.Network
	response.BusinessID = outputData.User.BusinessID
	response.CreationDate = outputData.User.CreationDate.Unix()
	response.LastUpdate = outputData.User.LastUpdate.Unix()

	responseWrapper := DescribeUserResponseWrapper{}
	responseWrapper.DescribeUserResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListUsers(ctx context.Context, data ListUserRequest) (*ListUserResponseWrapper, error) {
	var inputData user.ListUserInput
	inputData.Project = data.Project

	outputData, err := adapter.userUseCase.ListUsers(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListUserResponse
	response.Items = make([]User, 0)

	for _, item := range outputData.Items {
		var responseItem User
		responseItem.Project = item.Project
		responseItem.ID = item.ID
		responseItem.ContractID = item.ContractID
		responseItem.Name = item.Name
		responseItem.Email = item.Email
		responseItem.Role = item.Role
		responseItem.EthAddress = item.EthAddress
		responseItem.DefaultPasswordChanged = item.DefaultPasswordChanged
		responseItem.Network = item.Network
		responseItem.BusinessID = item.BusinessID
		responseItem.CreationDate = item.CreationDate.Unix()
		responseItem.LastUpdate = item.LastUpdate.Unix()

		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListUserResponseWrapper{}
	responseWrapper.ListUserResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptCreateUser(ctx context.Context, data CreateUserRequest) (*CreateUserResponseWrapper, error) {
	var inputData user.CreateUserInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.ContractID = data.ContractID
	inputData.Name = data.Name
	inputData.Email = data.Email
	inputData.Password = data.Password
	inputData.Role = data.Role
	inputData.EthAddress = data.EthAddress
	inputData.PrivateKey = data.PrivateKey
	inputData.Network = data.Network
	inputData.BusinessID = data.BusinessID

	outputData, err := adapter.userUseCase.CreateUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateUserResponse
	response.Project = outputData.User.Project
	response.ID = outputData.User.ID
	response.ContractID = outputData.User.ContractID
	response.Name = outputData.User.Name
	response.Email = outputData.User.Email
	response.Role = outputData.User.Role
	response.EthAddress = outputData.User.EthAddress
	response.DefaultPasswordChanged = outputData.User.DefaultPasswordChanged
	response.Network = outputData.User.Network
	response.BusinessID = outputData.User.BusinessID
	response.CreationDate = outputData.User.CreationDate.Unix()
	response.LastUpdate = outputData.User.LastUpdate.Unix()

	responseWrapper := CreateUserResponseWrapper{}
	responseWrapper.CreateUserResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptLoginUser(ctx context.Context, data LoginUserRequest) (*LoginUserResponseWrapper, error) {
	var inputData user.LoginUserInput
	inputData.Project = data.Project
	inputData.Email = data.Email
	inputData.Password = data.Password

	outputData, err := adapter.userUseCase.LoginUser(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response LoginUserResponse
	response.Project = outputData.User.Project
	response.ID = outputData.User.ID
	response.ContractID = outputData.User.ContractID
	response.Name = outputData.User.Name
	response.Email = outputData.User.Email
	response.Role = outputData.User.Role
	response.EthAddress = outputData.User.EthAddress
	response.DefaultPasswordChanged = outputData.User.DefaultPasswordChanged
	response.Network = outputData.User.Network
	response.BusinessID = outputData.User.BusinessID
	response.CreationDate = outputData.User.CreationDate.Unix()
	response.LastUpdate = outputData.User.LastUpdate.Unix()
	// TODO: response.Token = outputData.Token

	responseWrapper := LoginUserResponseWrapper{}
	responseWrapper.LoginUserResponse = response

	return &responseWrapper, nil
}
