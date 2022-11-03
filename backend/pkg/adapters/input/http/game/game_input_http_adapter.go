package gameInputHttpAdapter

import (
	"context"
	"errors"
	"time"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/game"
)

/* ADAPTERS TYPES */

type Game struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	OwnerUserID  string    `json:"owner-user-id"`
	CreationDate time.Time `json:"creation-date"`
	LastUpdate   time.Time `json:"last-update"`
}

type CreateGameRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	OwnerUserID string `json:"owner-user-id"`
}
type CreateGameResponse struct {
	Game
}
type CreateGameResponseWrapper struct {
	CreateGameResponse CreateGameResponse
}

type DescribeGameRequest struct {
	ID string `json:"id"`
}
type DescribeGameResponse struct {
	Game
}
type DescribeGameResponseWrapper struct {
	DescribeGameResponse DescribeGameResponse
}

type ListGameRequest struct{}

type ListGameResponse struct {
	Items []DescribeGameResponse `json:"items"`
}
type ListGameResponseWrapper struct {
	ListGameResponse ListGameResponse
}

type UpdateGameRequest struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	OwnerUserID string `json:"owner-user-id"`
}
type UpdateGameResponse struct {
	Game
}
type UpdateGameResponseWrapper struct {
	UpdateGameResponse UpdateGameResponse
}

type DeleteGameRequest struct {
	ID string `json:"id"`
}
type DeleteGameResponse struct {
	Game
}
type DeleteGameResponseWrapper struct {
	DeleteGameResponse DeleteGameResponse
}

type JoinGameRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user-id"`
}
type JoinGameResponse struct {
	Game
}
type JoinGameResponseWrapper struct {
	JoinGameResponse JoinGameResponse
}

/* INTERFACE ADAPTER */

type GameInputAdapter interface {
	AdaptCreateGame(ctx context.Context, data CreateGameRequest) (*CreateGameResponseWrapper, error)
	AdaptDescribeGame(ctx context.Context, data DescribeGameRequest) (*DescribeGameResponseWrapper, error)
	AdaptListGames(ctx context.Context, data ListGameRequest) (*ListGameResponseWrapper, error)
	AdaptUpdateGame(ctx context.Context, data UpdateGameRequest) (*UpdateGameResponseWrapper, error)
	AdaptDeleteGame(ctx context.Context, data DeleteGameRequest) (*DeleteGameResponseWrapper, error)
	AdaptJoinGame(ctx context.Context, data JoinGameRequest) (*JoinGameResponseWrapper, error)
}

/* DEFAULT USE CASE TYPES */

type DefaultHttpInputAdapter struct {
	gameUseCase game.GameUseCase
}

type DefaultHttpInputAdapterOptions struct {
	GameUseCase game.GameUseCase
}

/* CONSTRUCTOR */

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.GameUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		gameUseCase: options.GameUseCase,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ GameInputAdapter = (*DefaultHttpInputAdapter)(nil)

/* ADAPTERS IMPLEMENTATIONS */

func (adapter *DefaultHttpInputAdapter) AdaptCreateGame(ctx context.Context, data CreateGameRequest) (*CreateGameResponseWrapper, error) {

	// Map infra request to use case
	var inputData game.CreateGameInput
	inputData.ID = data.ID
	inputData.Name = data.Name
	inputData.OwnerUserID = data.OwnerUserID

	//Use Case
	outputData, err := adapter.gameUseCase.CreateGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response CreateGameResponse
	response.ID = outputData.ID
	response.Name = outputData.Name
	response.OwnerUserID = outputData.OwnerUserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := CreateGameResponseWrapper{}
	responseWrapper.CreateGameResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeGame(ctx context.Context, data DescribeGameRequest) (*DescribeGameResponseWrapper, error) {
	// Map infra request to use case
	var inputData game.DescribeGameInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.gameUseCase.DescribeGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DescribeGameResponse
	response.ID = outputData.ID
	response.Name = outputData.Name
	response.OwnerUserID = outputData.OwnerUserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeGameResponseWrapper{}
	responseWrapper.DescribeGameResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListGames(ctx context.Context, data ListGameRequest) (*ListGameResponseWrapper, error) {
	// Adapt payload input to use case input
	var inputData game.ListGameInput

	// Use case
	outputData, err := adapter.gameUseCase.ListGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Adapt use case output to payload output
	var response ListGameResponse
	response.Items = make([]DescribeGameResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeGameResponse
		responseItem.ID = item.ID
		responseItem.Name = item.Name
		responseItem.OwnerUserID = item.OwnerUserID
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListGameResponseWrapper{}
	responseWrapper.ListGameResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateGame(ctx context.Context, data UpdateGameRequest) (*UpdateGameResponseWrapper, error) {
	// Map infra request to use case
	var inputData game.UpdateGameInput
	inputData.ID = data.ID
	inputData.Name = data.Name
	inputData.OwnerUserID = data.OwnerUserID

	//Use Case
	outputData, err := adapter.gameUseCase.UpdateGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response UpdateGameResponse
	response.ID = outputData.ID
	response.Name = outputData.Name
	response.OwnerUserID = outputData.OwnerUserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := UpdateGameResponseWrapper{}
	responseWrapper.UpdateGameResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteGame(ctx context.Context, data DeleteGameRequest) (*DeleteGameResponseWrapper, error) {
	// Map infra request to use case
	var inputData game.RemoveGameInput
	inputData.ID = data.ID

	//Use Case
	outputData, err := adapter.gameUseCase.RemoveGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response DeleteGameResponse
	response.ID = outputData.ID
	response.Name = outputData.Name
	response.OwnerUserID = outputData.OwnerUserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteGameResponseWrapper{}
	responseWrapper.DeleteGameResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptJoinGame(ctx context.Context, data JoinGameRequest) (*JoinGameResponseWrapper, error) {

	// Map infra request to use case
	var inputData game.JoinGameInput
	inputData.ID = data.ID
	inputData.UserID = data.UserID

	//Use Case
	outputData, err := adapter.gameUseCase.JoinGame(ctx, inputData)
	if err != nil {
		return nil, err
	}

	// Map use case to infra response
	var response JoinGameResponse
	response.ID = outputData.ID
	response.Name = outputData.Name
	response.OwnerUserID = outputData.OwnerUserID
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := JoinGameResponseWrapper{}
	responseWrapper.JoinGameResponse = response

	return &responseWrapper, nil
}
