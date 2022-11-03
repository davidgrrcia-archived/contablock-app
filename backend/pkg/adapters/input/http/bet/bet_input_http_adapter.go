package betInputHttpAdapter

import (
	"context"
	"errors"
	"time"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
)

/* ADAPTERS TYPES */

type Bet struct {
	ID           string    `json:"id"`
	Amount       int64     `json:"amount"`
	GameID       string    `json:"game-id"`
	UserID       string    `json:"user-id"`
	CreationDate time.Time `json:"creation-date"`
	LastUpdate   time.Time `json:"last-update"`
}

type CreateBetRequest struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
	GameID string `json:"game-id"`
	UserID string `json:"user-id"`
}
type CreateBetResponse struct {
	Bet
}
type CreateBetResponseWrapper struct {
	CreateBetResponse CreateBetResponse
}

type DescribeBetRequest struct {
	ID string `json:"id"`
}
type DescribeBetResponse struct {
	Bet
}
type DescribeBetResponseWrapper struct {
	DescribeBetResponse DescribeBetResponse
}

type DescribeBetByUserIDRequest struct {
	UserID string `json:"user-id"`
	GameID string `json:"game-id"`
}
type DescribeBetByUserIDResponse struct {
	Bet
}
type DescribeBetByUserIDResponseWrapper struct {
	DescribeBetByIDResponse DescribeBetByUserIDResponse
}

type ListBetRequest struct{}

type ListBetResponse struct {
	Items []DescribeBetResponse `json:"items"`
}
type ListBetResponseWrapper struct {
	ListBetResponse ListBetResponse
}

type UpdateBetRequest struct {
	ID     string `json:"id"`
	Amount int64  `json:"amount"`
	GameID string `json:"game-id"`
	UserID string `json:"user-id"`
}
type UpdateBetResponse struct {
	Bet
}
type UpdateBetResponseWrapper struct {
	UpdateBetResponse UpdateBetResponse
}

type DeleteBetRequest struct {
	ID string `json:"id"`
}
type DeleteBetResponse struct {
	Bet
}
type DeleteBetResponseWrapper struct {
	DeleteBetResponse DeleteBetResponse
}

/* INTERFACE ADAPTER */

type BetInputAdapter interface {
	AdaptCreateBet(ctx context.Context, data CreateBetRequest) (*CreateBetResponseWrapper, error)
	AdaptDescribeBet(ctx context.Context, data DescribeBetRequest) (*DescribeBetResponseWrapper, error)
	AdaptDescribeBetByUserID(ctx context.Context, data DescribeBetByUserIDRequest) (*DescribeBetByUserIDResponseWrapper, error)
	AdaptListBets(ctx context.Context, data ListBetRequest) (*ListBetResponseWrapper, error)
	AdaptUpdateBet(ctx context.Context, data UpdateBetRequest) (*UpdateBetResponseWrapper, error)
	AdaptDeleteBet(ctx context.Context, data DeleteBetRequest) (*DeleteBetResponseWrapper, error)
}

/* DEFAULT USE CASE TYPES */

type DefaultHttpInputAdapter struct {
	betUseCase bet.BetUseCase
}

type DefaultHttpInputAdapterOptions struct {
	BetUseCase bet.BetUseCase
}

/* CONSTRUCTOR */

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.BetUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		betUseCase: options.BetUseCase,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ BetInputAdapter = (*DefaultHttpInputAdapter)(nil)

/* ADAPTERS IMPLEMENTATIONS */

func (adapter *DefaultHttpInputAdapter) AdaptCreateBet(ctx context.Context, data CreateBetRequest) (*CreateBetResponseWrapper, error) {

	// Map infra request to use case
	var inputData bet.CreateBetInput
	inputData.ID = data.ID
	inputData.Amount = data.Amount
	inputData.GameID = data.GameID
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.betUseCase.CreateBet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response CreateBetResponse
	response.ID = outputData.ID
	response.Amount = outputData.Amount
	response.GameID = outputData.GameID
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := CreateBetResponseWrapper{}
	responseWrapper.CreateBetResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeBet(ctx context.Context, data DescribeBetRequest) (*DescribeBetResponseWrapper, error) {
	// Map infra request to use case
	var inputData bet.DescribeBetInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.betUseCase.DescribeBet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DescribeBetResponse
	response.ID = outputData.ID
	response.Amount = outputData.Amount
	response.GameID = outputData.GameID
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeBetResponseWrapper{}
	responseWrapper.DescribeBetResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeBetByUserID(ctx context.Context, data DescribeBetByUserIDRequest) (*DescribeBetByUserIDResponseWrapper, error) {
	// Map infra request to use case
	var inputData bet.DescribeBetByUserIDInput
	inputData.UserID = data.UserID
	inputData.GameID = data.GameID

	//Use Case
	outputData, err := adapter.betUseCase.DescribeBetByUserID(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DescribeBetByUserIDResponse
	response.ID = outputData.ID
	response.Amount = outputData.Amount
	response.GameID = outputData.GameID
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeBetByUserIDResponseWrapper{}
	responseWrapper.DescribeBetByIDResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListBets(ctx context.Context, data ListBetRequest) (*ListBetResponseWrapper, error) {
	// Adapt payload input to use case input
	var inputData bet.ListBetInput

	// Use case
	outputData, err := adapter.betUseCase.ListBet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Adapt use case output to payload output
	var response ListBetResponse
	response.Items = make([]DescribeBetResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeBetResponse
		responseItem.ID = item.ID
		responseItem.Amount = item.Amount
		responseItem.GameID = item.GameID
		responseItem.UserID = item.UserID
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListBetResponseWrapper{}
	responseWrapper.ListBetResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateBet(ctx context.Context, data UpdateBetRequest) (*UpdateBetResponseWrapper, error) {
	// Map infra request to use case
	var inputData bet.UpdateBetInput
	inputData.ID = data.ID
	inputData.Amount = data.Amount
	inputData.GameID = data.GameID
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.betUseCase.UpdateBet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response UpdateBetResponse
	response.ID = outputData.ID
	response.Amount = outputData.Amount
	response.GameID = outputData.GameID
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := UpdateBetResponseWrapper{}
	responseWrapper.UpdateBetResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteBet(ctx context.Context, data DeleteBetRequest) (*DeleteBetResponseWrapper, error) {
	// Map infra request to use case
	var inputData bet.RemoveBetInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.betUseCase.RemoveBet(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DeleteBetResponse
	response.ID = outputData.ID
	response.Amount = outputData.Amount
	response.GameID = outputData.GameID
	response.UserID = outputData.UserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteBetResponseWrapper{}
	responseWrapper.DeleteBetResponse = response

	return &responseWrapper, nil
}
