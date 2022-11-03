package userWallet

import (
	"context"
	"errors"
	"github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/holdContract"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
)

/* USE CASE TYPES */

type UserWallet struct {
	entities.UserWallet
	CreationDate time.Time
	LastUpdate   time.Time
}

type CreateUserWalletInput struct {
	ID      string `valid:"required,length(1|64)"`
	Address string `valid:"required,length(1|64)"`
	UserID  string `valid:"required,length(1|64)"`
}

type CreateUserWalletOutput struct {
	UserWallet
}

type DescribeUserWalletInput struct {
	ID     string `valid:"required,length(1|64)"`
	UserID string `valid:"required,length(1|64)"`
}

type DescribeUserWalletOutput struct {
	UserWallet
}

type DescribeUserWalletsByUserIDInput struct {
	UserID string `valid:"required,length(1|64)"`
}

type DescribeUserWalletsByUserIDOutput struct {
	Items []UserWallet
}

type ListUserWalletInput struct{}

type ListUserWalletOutput struct {
	Items []UserWallet
}

type UpdateUserWalletInput struct {
	ID      string `valid:"required,length(1|64)"`
	Address string `valid:"required,length(1|64)"`
	UserID  string `valid:"required,length(1|64)"`
}

type UpdateUserWalletOutput struct {
	UserWallet
}

type RemoveUserWalletInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveUserWalletOutput struct {
	UserWallet
}

/* TYPES FOR STORAGE */

type AddInput struct {
	ID            string
	Address       string
	Balance       int64
	BalanceOnHold int64
	UserID        string
	CreationDate  time.Time
	LastUpdate    time.Time
}

type AddOutput struct {
	UserWallet
}

type GetInput struct {
	ID string
}

type GetOutput struct {
	UserWallet
}

type GetByUserIDInput struct {
	UserID string
}

type GetByUserIDOutput struct {
	Items []UserWallet
}

type AllInput struct{}

type AllOutput struct {
	Items []UserWallet
}

type EditInput struct {
	ID            string
	Address       string
	Balance       int64
	BalanceOnHold int64
	UserID        string
	LastUpdate    time.Time
}

type EditOutput struct {
	UserWallet
}

type RemoveInput struct {
	ID string
}

type RemoveOutput struct {
	UserWallet
}

/* COLLECTION TYPE */

type UserWalletCollection []UserWallet

/* INTERFACES */

type UserWalletUseCase interface {
	CreateUserWallet(ctx context.Context, input CreateUserWalletInput) (*CreateUserWalletOutput, error)
	DescribeUserWallet(ctx context.Context, input DescribeUserWalletInput) (*DescribeUserWalletOutput, error)
	DescribeUserWalletsByUserID(ctx context.Context, input DescribeUserWalletsByUserIDInput) (*DescribeUserWalletsByUserIDOutput, error)
	ListUserWallet(ctx context.Context, input ListUserWalletInput) (*ListUserWalletOutput, error)
	UpdateUserWallet(ctx context.Context, input UpdateUserWalletInput) (*UpdateUserWalletOutput, error)
	RemoveUserWallet(ctx context.Context, input RemoveUserWalletInput) (*RemoveUserWalletOutput, error)
}

type UserWalletStorage interface {
	Add(ctx context.Context, input AddInput) (*AddOutput, error)
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	GetByUserID(ctx context.Context, input GetByUserIDInput) (*GetByUserIDOutput, error)
	All(ctx context.Context, input AllInput) (*AllOutput, error)
	Edit(ctx context.Context, input EditInput) (*EditOutput, error)
	Remove(ctx context.Context, input RemoveInput) (*RemoveOutput, error)
}

/* INSTANCE TYPES */

type DefaultUserWalletUseCase struct {
	userWalletStorage UserWalletStorage
	holdContract      holdContract.HoldContractAdapter
}

type DefaultUserWalletUseCaseOptions struct {
	UserWalletStorage UserWalletStorage
	HoldContract      holdContract.HoldContractAdapter
}

/* CONSTRUCTOR TO CREATE NEW INSTANCES */

func NewDefaultUserWalletUseCase(options DefaultUserWalletUseCaseOptions) (*DefaultUserWalletUseCase, error) {
	if options.UserWalletStorage == nil {
		return nil, errors.New("cannot create an instance of UserWalletWalletUseCase because provided UserWalletWalletStorage is nil")
	}

	if options.HoldContract == nil {
		return nil, errors.New("cannot create an instance of UserWalletWalletUseCase because provided HoldContract is nil")
	}

	return &DefaultUserWalletUseCase{
		userWalletStorage: options.UserWalletStorage,
		holdContract:      options.HoldContract,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ UserWalletUseCase = (*DefaultUserWalletUseCase)(nil)

/* IMPLEMENTATIONS */

func (useCase *DefaultUserWalletUseCase) CreateUserWallet(ctx context.Context, input CreateUserWalletInput) (*CreateUserWalletOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var checkUserWallet GetInput
	checkUserWallet.ID = input.ID
	_, err = useCase.userWalletStorage.Get(ctx, checkUserWallet)
	if err == nil {
		return nil, err //todo nil fails
	}

	//map desired userWallet with input
	var desiredUserWallet AddInput
	desiredUserWallet.ID = input.ID
	desiredUserWallet.Address = input.Address
	desiredUserWallet.UserID = input.UserID

	//blockchain transaction
	err = useCase.holdContract.Transfer(ctx, desiredUserWallet.Address, 100)
	if err != nil {
		return nil, err
	}

	desiredUserWallet.Balance, err = useCase.holdContract.BalanceOf(ctx, desiredUserWallet.Address)
	desiredUserWallet.BalanceOnHold, err = useCase.holdContract.BalanceOnHoldOf(ctx, desiredUserWallet.Address)

	desiredUserWallet.CreationDate = time.Now()
	desiredUserWallet.LastUpdate = desiredUserWallet.CreationDate

	storedUserWallet, err := useCase.userWalletStorage.Add(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	return &CreateUserWalletOutput{
		UserWallet: storedUserWallet.UserWallet,
	}, err
}

func (useCase *DefaultUserWalletUseCase) DescribeUserWallet(ctx context.Context, input DescribeUserWalletInput) (*DescribeUserWalletOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired userWallet with input
	var desiredUserWallet GetInput
	desiredUserWallet.ID = input.ID

	storedUserWallet, err := useCase.userWalletStorage.Get(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	if useCase.isBalanceChanged(ctx, storedUserWallet.UserWallet) == true {
		storedUserWalletUpdated, err := useCase.userWalletStorage.Edit(ctx, useCase.updateBalance(ctx, storedUserWallet.UserWallet))
		if err != nil {
			return nil, err
		}

		return &DescribeUserWalletOutput{
			UserWallet: storedUserWalletUpdated.UserWallet,
		}, err

	}

	return &DescribeUserWalletOutput{
		UserWallet: storedUserWallet.UserWallet,
	}, err
}

func (useCase *DefaultUserWalletUseCase) DescribeUserWalletsByUserID(ctx context.Context, input DescribeUserWalletsByUserIDInput) (*DescribeUserWalletsByUserIDOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired userWallet with input
	var desiredUserWallet GetByUserIDInput
	desiredUserWallet.UserID = input.UserID

	userWalletList, err := useCase.userWalletStorage.GetByUserID(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	var items = userWalletList.Items
	for _, w := range items {
		if useCase.isBalanceChanged(ctx, w) == true {
			_, err := useCase.userWalletStorage.Edit(ctx, useCase.updateBalance(ctx, w))
			if err != nil {
				return nil, err
			}
		}
	}

	return &DescribeUserWalletsByUserIDOutput{
		Items: userWalletList.Items,
	}, err
}

func (useCase *DefaultUserWalletUseCase) ListUserWallet(ctx context.Context, input ListUserWalletInput) (*ListUserWalletOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired userWallet with input
	var desiredUserWallet AllInput

	userWalletList, err := useCase.userWalletStorage.All(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	//update blockchain balance if necessary
	var items = userWalletList.Items
	for _, w := range items {
		if useCase.isBalanceChanged(ctx, w) == true {
			_, err := useCase.userWalletStorage.Edit(ctx, useCase.updateBalance(ctx, w))
			if err != nil {
				return nil, err
			}
		}
	}

	return &ListUserWalletOutput{
		Items: items,
	}, err
}

func (useCase *DefaultUserWalletUseCase) UpdateUserWallet(ctx context.Context, input UpdateUserWalletInput) (*UpdateUserWalletOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired userWallet with input
	var desiredUserWallet EditInput
	desiredUserWallet.ID = input.ID
	desiredUserWallet.Address = input.Address
	desiredUserWallet.UserID = input.UserID

	desiredUserWallet.Balance, err = useCase.holdContract.BalanceOf(ctx, desiredUserWallet.Address)
	if err != nil {
		return nil, err
	}
	desiredUserWallet.Balance, err = useCase.holdContract.BalanceOnHoldOf(ctx, desiredUserWallet.Address)
	if err != nil {
		return nil, err
	}

	desiredUserWallet.LastUpdate = time.Now()

	storedUserWallet, err := useCase.userWalletStorage.Edit(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	return &UpdateUserWalletOutput{
		UserWallet: storedUserWallet.UserWallet,
	}, err
}

func (useCase *DefaultUserWalletUseCase) RemoveUserWallet(ctx context.Context, input RemoveUserWalletInput) (*RemoveUserWalletOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired userWallet with input
	var desiredUserWallet RemoveInput
	desiredUserWallet.ID = input.ID

	storedUserWallet, err := useCase.userWalletStorage.Remove(ctx, desiredUserWallet)
	if err != nil {
		return nil, err
	}

	//update blockchain balance if necessary
	if useCase.isBalanceChanged(ctx, storedUserWallet.UserWallet) == true {
		storedUserWalletUpdated, err := useCase.userWalletStorage.Edit(ctx, useCase.updateBalance(ctx, storedUserWallet.UserWallet))
		if err != nil {
			return nil, err
		}

		return &RemoveUserWalletOutput{
			UserWallet: storedUserWalletUpdated.UserWallet,
		}, err
	}

	return &RemoveUserWalletOutput{
		UserWallet: storedUserWallet.UserWallet,
	}, err
}

func (useCase *DefaultUserWalletUseCase) isBalanceChanged(ctx context.Context, wallet UserWallet) bool {

	amount, _ := useCase.holdContract.BalanceOf(ctx, wallet.Address)
	if wallet.Balance != amount {
		return true
	}

	amountOnHold, _ := useCase.holdContract.BalanceOnHoldOf(ctx, wallet.Address)
	if wallet.BalanceOnHold != amountOnHold {
		return true
	}

	return false
}

func (useCase *DefaultUserWalletUseCase) updateBalance(ctx context.Context, wallet UserWallet) EditInput {

	var u EditInput
	u.ID = wallet.ID
	u.Address = wallet.Address
	u.Balance, _ = useCase.holdContract.BalanceOf(ctx, wallet.Address)
	u.BalanceOnHold, _ = useCase.holdContract.BalanceOnHoldOf(ctx, wallet.Address)
	u.UserID = wallet.UserID
	u.LastUpdate = time.Now()

	return EditInput{
		ID:            u.ID,
		Address:       u.Address,
		Balance:       u.Balance,
		BalanceOnHold: u.BalanceOnHold,
		UserID:        u.UserID,
		LastUpdate:    u.LastUpdate,
	}
}
