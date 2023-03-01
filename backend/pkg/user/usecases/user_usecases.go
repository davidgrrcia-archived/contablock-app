package user

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/google/uuid"
	"time"

	"github.com/asaskevich/govalidator"
)

type UserUseCase interface {
	DescribeUser(ctx context.Context, input DescribeUserInput) (*DescribeUserOutput, error)
	ListUsers(ctx context.Context, input ListUserInput) (*ListUserOutput, error)
	CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
	LoginUser(ctx context.Context, input LoginUserInput) (*LoginUserOutput, error)
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
	output.User.DefaultPasswordChanged = storedUser.DefaultPasswordChanged
	output.User.Network = storedUser.Network
	output.User.BusinessID = storedUser.BusinessID
	output.User.CreationDate = storedUser.CreationDate
	output.User.LastUpdate = storedUser.LastUpdate

	return &output, err
}

func (useCase *DefaultUserUseCase) ListUsers(ctx context.Context, input ListUserInput) (*ListUserOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var filters Filters
	storedUsers, err := useCase.userStorage.All(ctx, input.Project, filters)
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
		user.DefaultPasswordChanged = item.DefaultPasswordChanged
		user.Network = item.Network
		user.BusinessID = item.BusinessID
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
		return nil, errors.New("invalid input: " + err.Error())
	}

	//TODO if the project in the path is the admin, then allow to add

	storedUser, err := useCase.userStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedUser != nil {
		return nil, errors.New("user [" + storedUser.ID + "] already exists")
	}

	if input.ID == "" {
		input.ID = uuid.New().String()
	}

	var desiredUser UserStorageData
	desiredUser.Project = input.Project
	desiredUser.ID = input.ID
	desiredUser.Name = input.Name
	desiredUser.Email = input.Email
	desiredUser.Password = hashPassword(input.Password)
	desiredUser.Role = input.Role
	desiredUser.EthAddress = input.EthAddress
	desiredUser.DefaultPasswordChanged = false
	desiredUser.Network = input.Network
	desiredUser.BusinessID = input.BusinessID
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
	output.User.DefaultPasswordChanged = addedUser.DefaultPasswordChanged
	output.User.Network = addedUser.Network
	output.User.BusinessID = addedUser.BusinessID
	output.User.CreationDate = addedUser.CreationDate
	output.User.LastUpdate = addedUser.LastUpdate

	return &output, nil
}

func (useCase *DefaultUserUseCase) LoginUser(ctx context.Context, input LoginUserInput) (*LoginUserOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var filters Filters
	filters.Email = &input.Email
	hashedPassword := hashPassword(input.Password)
	filters.Password = &hashedPassword
	storedUsers, err := useCase.userStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	if len(storedUsers.Items) == 0 {
		return nil, errors.New("user not found")
	}

	if len(storedUsers.Items) > 1 {
		return nil, errors.New("more than one user found")
	}

	storedUser := storedUsers.Items[0]

	var output LoginUserOutput
	output.User.Project = storedUser.Project
	output.User.ID = storedUser.ID
	output.User.Name = storedUser.Name
	output.User.Email = storedUser.Email
	output.User.Password = storedUser.Password
	output.User.Role = storedUser.Role
	output.User.EthAddress = storedUser.EthAddress
	output.User.DefaultPasswordChanged = storedUser.DefaultPasswordChanged
	output.User.Network = storedUser.Network
	output.User.BusinessID = storedUser.BusinessID
	output.User.CreationDate = storedUser.CreationDate
	output.User.LastUpdate = storedUser.LastUpdate

	return &output, nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
