package userOutputDbAdapter

import (
	"context"
	"errors"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
)

/* AUX VAR */

var _ user.UserStorage = (*InMemoryUserStorage)(nil)

/* INSTANCE TYPES */

type InMemoryUserStorageOptions struct {
	UserDatabase user.UserCollection
}

type InMemoryUserStorage struct {
	userDatabase user.UserCollection
}

/* INSTANCE CONSTRUCTOR */

func NewInMemoryUserStorage(opts InMemoryUserStorageOptions) (*InMemoryUserStorage, error) {
	if opts.UserDatabase == nil {
		return nil, errors.New("cannot create an instance of InMemoryUserStorage because provided UserDatabase is nil")
	}

	return &InMemoryUserStorage{
		userDatabase: opts.UserDatabase,
	}, nil
}

func (adapter *InMemoryUserStorage) Add(ctx context.Context, input user.AddInput) (*user.AddOutput, error) {

	//map use case to storage
	var u user.User
	u.ID = input.ID
	u.Email = input.Email
	u.Username = input.Username
	u.Password = input.Password
	u.CreationDate = input.CreationDate
	u.LastUpdate = input.LastUpdate

	adapter.userDatabase = append(adapter.userDatabase, u)
	return &user.AddOutput{
		User: u,
	}, nil
}

func (adapter *InMemoryUserStorage) Get(ctx context.Context, input user.GetInput) (*user.GetOutput, error) {

	for _, u := range adapter.userDatabase {
		if u.ID != input.ID {
			continue
		}

		return &user.GetOutput{
			User: u,
		}, nil
	}

	return nil, errors.New("user not found")
}

func (adapter *InMemoryUserStorage) All(ctx context.Context, input user.AllInput) (*user.AllOutput, error) {
	return &user.AllOutput{
		Items: adapter.userDatabase,
	}, nil
}

func (adapter *InMemoryUserStorage) Edit(ctx context.Context, input user.EditInput) (*user.EditOutput, error) {

	for _, u := range adapter.userDatabase {
		if u.ID == input.ID {
			//map use case to storage
			u.Username = input.Username
			u.Email = input.Email
			u.Password = input.Password
			u.LastUpdate = input.LastUpdate

			return &user.EditOutput{
				User: u,
			}, nil
		}
	}

	return nil, errors.New("user not found")
}

func (adapter *InMemoryUserStorage) Remove(ctx context.Context, input user.RemoveInput) (*user.RemoveOutput, error) {
	for i, u := range adapter.userDatabase {
		if u.ID != input.ID {
			continue
		}

		adapter.userDatabase = append(adapter.userDatabase[:i], adapter.userDatabase[i+1:]...)

		return &user.RemoveOutput{
			User: u,
		}, nil
	}

	return nil, errors.New("user not found")
}
