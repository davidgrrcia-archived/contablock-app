package user

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

type UserUseCase interface {
	DescribeUser(ctx context.Context, input DescribeUserInput) (*DescribeUserOutput, error)
	ListUsers(ctx context.Context, input ListUserInput) (*ListUserOutput, error)
	CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
}

func NewDefaultUserUseCase(options DefaultUserUseCaseOptions) (*DefaultUserUseCase, error) {
	if options.UserStorage == nil {
		return nil, errors.New("cannot create an instance of UserUseCase because provided UserStorage is nil")
	}

	return &DefaultUserUseCase{
		userStorage: options.UserStorage,
	}, nil
}

type DefaultUserUseCase struct {
	userStorage UserStorage
}

type DefaultUserUseCaseOptions struct {
	UserStorage UserStorage
}

var _ UserUseCase = (*DefaultUserUseCase)(nil)

func (useCase *DefaultUserUseCase) DescribeUser(ctx context.Context, input DescribeUserInput) (*DescribeUserOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedUser, err := useCase.userStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeUserOutput
	output.User.Project = storedUser.Project
	output.User.ID = storedUser.ID
	output.User.Name = storedUser.Name
	output.User.Email = storedUser.Email
	output.User.Password = storedUser.Password
	output.User.Role = storedUser.Role
	output.User.EthAddress = storedUser.EthAddress
	output.User.CreationDate = storedUser.CreationDate
	output.User.LastUpdate = storedUser.LastUpdate

	return &output, err
}

func (useCase *DefaultUserUseCase) ListUsers(ctx context.Context, input ListUserInput) (*ListUserOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedUsers, err := useCase.userStorage.All(ctx, input.Project)
	if err != nil {
		return nil, err
	}

	var users []User

	for _, item := range storedUsers.Items {
		var user User
		user.Project = item.Project
		user.ID = item.ID
		user.Name = item.Name
		user.Email = item.Email
		user.Password = item.Password
		user.Role = item.Role
		user.EthAddress = item.EthAddress
		user.CreationDate = item.CreationDate
		user.LastUpdate = item.LastUpdate
		users = append(users, user)
	}

	return &ListUserOutput{
		Items: users,
	}, err
}

func (useCase *DefaultUserUseCase) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//TODO if the project in the path is the admin, then allow to add

	storedUser, err := useCase.userStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedUser != nil {
		return nil, errors.New("user [" + storedUser.ID + "] already exists")
	}

	var desiredUser UserStorageData
	desiredUser.Project = input.Project
	desiredUser.ID = input.ID
	desiredUser.Name = input.Name
	desiredUser.Email = input.Email
	desiredUser.Password = input.Password
	desiredUser.Role = input.Role
	desiredUser.EthAddress = input.EthAddress
	desiredUser.CreationDate = time.Now()
	desiredUser.LastUpdate = desiredUser.CreationDate

	addedUser, err := useCase.userStorage.Add(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	var output CreateUserOutput
	output.User.Project = addedUser.Project
	output.User.ID = addedUser.ID
	output.User.Name = addedUser.Name
	output.User.Email = addedUser.Email
	output.User.Password = addedUser.Password
	output.User.Role = addedUser.Role
	output.User.EthAddress = addedUser.EthAddress
	output.User.CreationDate = addedUser.CreationDate
	output.User.LastUpdate = addedUser.LastUpdate

	return &output, nil
}
