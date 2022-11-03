package user

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
)

/* USE CASE TYPES */

type User struct {
	entities.User
	CreationDate time.Time
	LastUpdate   time.Time
}

type CreateUserInput struct {
	ID       string `valid:"required,length(1|64)"`
	Email    string `valid:"required,length(1|64)"`
	Username string `valid:"required,length(1|64)"`
	Password string `valid:"required,length(1|64)"`
}

type CreateUserOutput struct {
	User
}

type DescribeUserInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeUserOutput struct {
	User
}

type ListUserInput struct{}

type ListUserOutput struct {
	Items []User
}

type UpdateUserInput struct {
	ID       string `valid:"required,length(1|64)"`
	Email    string `valid:"required,length(1|64)"`
	Username string `valid:"required,length(1|64)"`
	Password string `valid:"required,length(1|64)"`
}

type UpdateUserOutput struct {
	User
}

type RemoveUserInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveUserOutput struct {
	User
}

/* TYPES FOR STORAGE */

type AddInput struct {
	ID           string
	Email        string
	Username     string
	Password     string
	CreationDate time.Time
	LastUpdate   time.Time
}

type AddOutput struct {
	User
}

type GetInput struct {
	ID string
}

type GetOutput struct {
	User
}

type AllInput struct{}

type AllOutput struct {
	Items []User
}

type EditInput struct {
	ID         string
	Email      string
	Username   string
	Password   string
	LastUpdate time.Time
}

type EditOutput struct {
	User
}

type RemoveInput struct {
	ID string
}

type RemoveOutput struct {
	User
}

/* COLLECTION TYPE */

type UserCollection []User

/* INTERFACES */

type UserUseCase interface {
	CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error)
	DescribeUser(ctx context.Context, input DescribeUserInput) (*DescribeUserOutput, error)
	ListUser(ctx context.Context, input ListUserInput) (*ListUserOutput, error)
	UpdateUser(ctx context.Context, input UpdateUserInput) (*UpdateUserOutput, error)
	RemoveUser(ctx context.Context, input RemoveUserInput) (*RemoveUserOutput, error)
}

type UserStorage interface {
	Add(ctx context.Context, input AddInput) (*AddOutput, error)
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	All(ctx context.Context, input AllInput) (*AllOutput, error)
	Edit(ctx context.Context, input EditInput) (*EditOutput, error)
	Remove(ctx context.Context, input RemoveInput) (*RemoveOutput, error)
}

/* INSTANCE TYPES */

type DefaultUserUseCase struct {
	userStorage UserStorage
}

type DefaultUserUseCaseOptions struct {
	UserStorage UserStorage
}

/* CONSTRUCTOR TO CREATE NEW INSTANCES */

func NewDefaultUserUseCase(options DefaultUserUseCaseOptions) (*DefaultUserUseCase, error) {
	if options.UserStorage == nil {
		return nil, errors.New("cannot create an instance of UserWalletUseCase because provided UserWalletStorage is nil")
	}

	return &DefaultUserUseCase{
		userStorage: options.UserStorage,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ UserUseCase = (*DefaultUserUseCase)(nil)

/* IMPLEMENTATIONS */

func (useCase *DefaultUserUseCase) CreateUser(ctx context.Context, input CreateUserInput) (*CreateUserOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//CRUD validation

	var checkUser GetInput
	checkUser.ID = input.ID
	_, err = useCase.userStorage.Get(ctx, checkUser)
	if err == nil {
		return nil, errors.New("user already exists")
	}

	//map desired user with input
	var desiredUser AddInput
	desiredUser.ID = input.ID
	desiredUser.Username = input.Username
	desiredUser.Password, err = hashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	desiredUser.Email = input.Email
	desiredUser.CreationDate = time.Now()
	desiredUser.LastUpdate = desiredUser.CreationDate

	storedUser, err := useCase.userStorage.Add(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	return &CreateUserOutput{
		User: storedUser.User,
	}, nil
}

func (useCase *DefaultUserUseCase) DescribeUser(ctx context.Context, input DescribeUserInput) (*DescribeUserOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired user with input
	var desiredUser GetInput
	desiredUser.ID = input.ID

	storedUser, err := useCase.userStorage.Get(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	return &DescribeUserOutput{
		User: storedUser.User,
	}, err
}

func (useCase *DefaultUserUseCase) ListUser(ctx context.Context, input ListUserInput) (*ListUserOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired user with input
	var desiredUser AllInput

	userList, err := useCase.userStorage.All(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	var items = userList.Items

	return &ListUserOutput{
		Items: items,
	}, err
}

func (useCase *DefaultUserUseCase) UpdateUser(ctx context.Context, input UpdateUserInput) (*UpdateUserOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired user with input
	var desiredUser EditInput
	desiredUser.ID = input.ID
	desiredUser.Username = input.Username
	desiredUser.Password, err = hashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	desiredUser.Email = input.Email
	desiredUser.LastUpdate = time.Now()

	storedUser, err := useCase.userStorage.Edit(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	return &UpdateUserOutput{
		User: storedUser.User,
	}, err
}

func (useCase *DefaultUserUseCase) RemoveUser(ctx context.Context, input RemoveUserInput) (*RemoveUserOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired user with input
	var desiredUser RemoveInput
	desiredUser.ID = input.ID

	storedUser, err := useCase.userStorage.Remove(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	return &RemoveUserOutput{
		User: storedUser.User,
	}, err
}

func hashPassword(password string) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(passwordHash), nil
}

func (u User) passwordMatch(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))

	return err == nil
}
