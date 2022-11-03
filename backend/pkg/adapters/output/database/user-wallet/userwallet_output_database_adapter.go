package userWalletOutputDbAdapter

import (
	"context"
	"errors"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user-wallet"
)

/* AUX VAR */

var _ userWallet.UserWalletStorage = (*InMemoryUserWalletStorage)(nil)

/* INSTANCE TYPES */

type InMemoryUserWalletStorageOptions struct {
	UserWalletDatabase userWallet.UserWalletCollection
}

type InMemoryUserWalletStorage struct {
	userWalletDatabase userWallet.UserWalletCollection
}

/* INSTANCE CONSTRUCTOR */

func NewInMemoryUserWalletStorage(opts InMemoryUserWalletStorageOptions) (*InMemoryUserWalletStorage, error) {
	if opts.UserWalletDatabase == nil {
		return nil, errors.New("cannot create an instance of InMemoryUserWalletStorage because provided UserWalletDatabase is nil")
	}

	return &InMemoryUserWalletStorage{
		userWalletDatabase: opts.UserWalletDatabase,
	}, nil
}

func (adapter *InMemoryUserWalletStorage) Add(ctx context.Context, input userWallet.AddInput) (*userWallet.AddOutput, error) {

	//map use case to storage
	var u userWallet.UserWallet
	u.ID = input.ID
	u.Address = input.Address
	u.Balance = input.Balance
	u.BalanceOnHold = input.BalanceOnHold
	u.UserID = input.UserID
	u.CreationDate = input.CreationDate
	u.LastUpdate = input.LastUpdate

	adapter.userWalletDatabase = append(adapter.userWalletDatabase, u)
	return &userWallet.AddOutput{
		UserWallet: u,
	}, nil
}

func (adapter *InMemoryUserWalletStorage) Get(ctx context.Context, input userWallet.GetInput) (*userWallet.GetOutput, error) {

	for _, u := range adapter.userWalletDatabase {
		if u.ID != input.ID && u.UserID != input.ID {
			continue
		}

		return &userWallet.GetOutput{
			UserWallet: u,
		}, nil
	}

	return nil, errors.New("userWallet not found")
}

func (adapter *InMemoryUserWalletStorage) GetByUserID(ctx context.Context, input userWallet.GetByUserIDInput) (*userWallet.GetByUserIDOutput, error) {

	var items userWallet.UserWalletCollection

	for _, u := range adapter.userWalletDatabase {
		if u.UserID == input.UserID {
			items = append(items, u)
		}
	}

	return &userWallet.GetByUserIDOutput{
		Items: items,
	}, nil
}

func (adapter *InMemoryUserWalletStorage) All(ctx context.Context, input userWallet.AllInput) (*userWallet.AllOutput, error) {
	return &userWallet.AllOutput{
		Items: adapter.userWalletDatabase,
	}, nil
}

func (adapter *InMemoryUserWalletStorage) Edit(ctx context.Context, input userWallet.EditInput) (*userWallet.EditOutput, error) {

	for _, u := range adapter.userWalletDatabase {
		if u.ID == input.ID {
			//map use case to storage
			u.Address = input.Address
			u.Balance = input.Balance
			u.BalanceOnHold = input.BalanceOnHold
			u.UserID = input.UserID
			u.LastUpdate = input.LastUpdate

			return &userWallet.EditOutput{
				UserWallet: u,
			}, nil
		}
	}

	return nil, errors.New("userWallet not found")
}

func (adapter *InMemoryUserWalletStorage) Remove(ctx context.Context, input userWallet.RemoveInput) (*userWallet.RemoveOutput, error) {
	for i, u := range adapter.userWalletDatabase {
		if u.ID != input.ID {
			continue
		}

		adapter.userWalletDatabase = append(adapter.userWalletDatabase[:i], adapter.userWalletDatabase[i+1:]...)

		return &userWallet.RemoveOutput{
			UserWallet: u,
		}, nil
	}

	return nil, errors.New("userWallet not found")
}
