package userWalletInputHttpAdapter

import (
	"context"
	"errors"
	"time"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user-wallet"
)

/* ADAPTERS TYPES */

type UserWallet struct {
	ID            string    `json:"id"`
	Address       string    `json:"address"`
	Balance       int64     `json:"balance"`
	BalanceOnHold int64     `json:"balance-on-hold"`
	UserID        string    `json:"user-id"`
	CreationDate  time.Time `json:"creation-date"`
	LastUpdate    time.Time `json:"last-update"`
}

type CreateUserWalletRequest struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	UserID  string `json:"user-id"`
}
type CreateUserWalletResponse struct {
	UserWallet
}
type CreateUserWalletResponseWrapper struct {
	CreateUserWalletResponse CreateUserWalletResponse
}

type DescribeUserWalletRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user-id"`
}
type DescribeUserWalletResponse struct {
	UserWallet
}
type DescribeUserWalletResponseWrapper struct {
	DescribeUserWalletResponse DescribeUserWalletResponse
}

type DescribeUserWalletsByUserIDRequest struct {
	UserID string `json:"user-id"`
}
type DescribeUserWalletsByUserIDResponse struct {
	Items []DescribeUserWalletResponse `json:"items"`
}
type DescribeUserWalletsByUserIDResponseWrapper struct {
	DescribeUserWalletsByIDResponse DescribeUserWalletsByUserIDResponse
}

type ListUserWalletRequest struct{}

type ListUserWalletResponse struct {
	Items []DescribeUserWalletResponse `json:"items"`
}
type ListUserWalletResponseWrapper struct {
	ListUserWalletResponse ListUserWalletResponse
}

type UpdateUserWalletRequest struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	UserID  string `json:"user-id"`
}
type UpdateUserWalletResponse struct {
	UserWallet
}
type UpdateUserWalletResponseWrapper struct {
	UpdateUserWalletResponse UpdateUserWalletResponse
}

type DeleteUserWalletRequest struct {
	ID string `json:"id"`
}
type DeleteUserWalletResponse struct {
	UserWallet
}
type DeleteUserWalletResponseWrapper struct {
	DeleteUserWalletResponse DeleteUserWalletResponse
}

/* INTERFACE ADAPTER */

type UserWalletInputAdapter interface {
	AdaptCreateUserWallet(ctx context.Context, data CreateUserWalletRequest) (*CreateUserWalletResponseWrapper, error)
	AdaptDescribeUserWallet(ctx context.Context, data DescribeUserWalletRequest) (*DescribeUserWalletResponseWrapper, error)
	AdaptDescribeUserWalletsByUserID(ctx context.Context, data DescribeUserWalletsByUserIDRequest) (*DescribeUserWalletsByUserIDResponseWrapper, error)
	AdaptListUserWallets(ctx context.Context, data ListUserWalletRequest) (*ListUserWalletResponseWrapper, error)
	AdaptUpdateUserWallet(ctx context.Context, data UpdateUserWalletRequest) (*UpdateUserWalletResponseWrapper, error)
	AdaptDeleteUserWallet(ctx context.Context, data DeleteUserWalletRequest) (*DeleteUserWalletResponseWrapper, error)
}

/* DEFAULT USE CASE TYPES */

type DefaultHttpInputAdapter struct {
	userWalletUseCase userWallet.UserWalletUseCase
}

type DefaultHttpInputAdapterOptions struct {
	UserWalletUseCase userWallet.UserWalletUseCase
}

/* CONSTRUCTOR */

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.UserWalletUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		userWalletUseCase: options.UserWalletUseCase,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ UserWalletInputAdapter = (*DefaultHttpInputAdapter)(nil)

/* ADAPTERS IMPLEMENTATIONS */

func (adapter *DefaultHttpInputAdapter) AdaptCreateUserWallet(ctx context.Context, data CreateUserWalletRequest) (*CreateUserWalletResponseWrapper, error) {

	// Map infra request to use case
	var inputData userWallet.CreateUserWalletInput
	inputData.ID = data.ID
	inputData.Address = data.Address
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.userWalletUseCase.CreateUserWallet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response CreateUserWalletResponse
	response.ID = outputData.ID
	response.Address = outputData.Address
	response.Balance = outputData.Balance
	response.BalanceOnHold = outputData.BalanceOnHold
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := CreateUserWalletResponseWrapper{}
	responseWrapper.CreateUserWalletResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeUserWallet(ctx context.Context, data DescribeUserWalletRequest) (*DescribeUserWalletResponseWrapper, error) {
	// Map infra request to use case
	var inputData userWallet.DescribeUserWalletInput
	inputData.ID = data.ID
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.userWalletUseCase.DescribeUserWallet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DescribeUserWalletResponse
	response.ID = outputData.ID
	response.Address = outputData.Address
	response.Balance = outputData.Balance
	response.BalanceOnHold = outputData.BalanceOnHold
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeUserWalletResponseWrapper{}
	responseWrapper.DescribeUserWalletResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeUserWalletsByUserID(ctx context.Context, data DescribeUserWalletsByUserIDRequest) (*DescribeUserWalletsByUserIDResponseWrapper, error) {
	// Map infra request to use case
	var inputData userWallet.DescribeUserWalletsByUserIDInput
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.userWalletUseCase.DescribeUserWalletsByUserID(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Adapt use case output to payload output
	var response DescribeUserWalletsByUserIDResponse
	response.Items = make([]DescribeUserWalletResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeUserWalletResponse
		responseItem.ID = item.ID
		responseItem.Address = item.Address
		responseItem.Balance = item.Balance
		responseItem.BalanceOnHold = item.BalanceOnHold
		responseItem.UserID = item.UserID
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := DescribeUserWalletsByUserIDResponseWrapper{}
	responseWrapper.DescribeUserWalletsByIDResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListUserWallets(ctx context.Context, data ListUserWalletRequest) (*ListUserWalletResponseWrapper, error) {
	// Adapt payload input to use case input
	var inputData userWallet.ListUserWalletInput

	// Use case
	outputData, err := adapter.userWalletUseCase.ListUserWallet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Adapt use case output to payload output
	var response ListUserWalletResponse
	response.Items = make([]DescribeUserWalletResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeUserWalletResponse
		responseItem.ID = item.ID
		responseItem.Address = item.Address
		responseItem.Balance = item.Balance
		responseItem.BalanceOnHold = item.BalanceOnHold
		responseItem.UserID = item.UserID
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListUserWalletResponseWrapper{}
	responseWrapper.ListUserWalletResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateUserWallet(ctx context.Context, data UpdateUserWalletRequest) (*UpdateUserWalletResponseWrapper, error) {
	// Map infra request to use case
	var inputData userWallet.UpdateUserWalletInput
	inputData.ID = data.ID
	inputData.Address = data.Address
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.userWalletUseCase.UpdateUserWallet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response UpdateUserWalletResponse
	response.ID = outputData.ID
	response.Address = outputData.Address
	response.Balance = outputData.Balance
	response.BalanceOnHold = outputData.BalanceOnHold
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := UpdateUserWalletResponseWrapper{}
	responseWrapper.UpdateUserWalletResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteUserWallet(ctx context.Context, data DeleteUserWalletRequest) (*DeleteUserWalletResponseWrapper, error) {
	// Map infra request to use case
	var inputData userWallet.RemoveUserWalletInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.userWalletUseCase.RemoveUserWallet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DeleteUserWalletResponse
	response.ID = outputData.ID
	response.Address = outputData.Address
	response.Balance = outputData.Balance
	response.BalanceOnHold = outputData.BalanceOnHold
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteUserWalletResponseWrapper{}
	responseWrapper.DeleteUserWalletResponse = response

	return &responseWrapper, nil
}
