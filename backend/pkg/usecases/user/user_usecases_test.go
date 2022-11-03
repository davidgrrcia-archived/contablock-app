package user_test

import (
	"context"
	"errors"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
	"testing"
	"time"
)

type mockedUserStorage struct{}

func (m *mockedUserStorage) Add(ctx context.Context, input user.AddInput) (*user.AddOutput, error) {
	var u = user.User{
		User: entities.User{
			ID:       input.ID,
			Email:    input.Email,
			Username: input.Username,
			Password: input.Password,
		},
		CreationDate: time.Now(),
		LastUpdate:   time.Now(),
	}

	return &user.AddOutput{
		User: u,
	}, nil
}

func (m *mockedUserStorage) Get(ctx context.Context, input user.GetInput) (*user.GetOutput, error) {
	/*var u = user.User{
		User: entities.User{
			ID: input.ID,
		},
		CreationDate: time.Now(),
		LastUpdate:   time.Now(),
	}

	return &user.GetOutput{
		User: u,
	}, nil*/

	return nil, errors.New("user already exists")
}

func (m *mockedUserStorage) All(ctx context.Context, input user.AllInput) (*user.AllOutput, error) {
	panic("implement me")
}

func (m *mockedUserStorage) Edit(ctx context.Context, input user.EditInput) (*user.EditOutput, error) {
	panic("implement me")
}

func (m *mockedUserStorage) Remove(ctx context.Context, input user.RemoveInput) (*user.RemoveOutput, error) {
	panic("implement me")
}

var _ user.UserStorage = (*mockedUserStorage)(nil)

func TestCreateUser(t *testing.T) {

	//before each test

	options := user.DefaultUserUseCaseOptions{
		UserStorage: &mockedUserStorage{},
	}

	useCase, _ := user.NewDefaultUserUseCase(options)

	//desired input
	input := user.CreateUserInput{
		ID:       "1",
		Email:    "test@test.com",
		Username: "david",
		Password: "123",
	}

	output, err := useCase.CreateUser(context.Background(), input)

	if err != nil {
		t.Fatalf("error in createUser method")
	}

	if output.ID != input.ID || output.Username != input.Username {
		t.Errorf("error creating user")
	}
}
