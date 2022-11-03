package game

import (
	"context"
	"errors"
	"github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/holdContract"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
)

/* USE CASE TYPES */

type Game struct {
	entities.Game
	CreationDate time.Time
	LastUpdate   time.Time
}

type CreateGameInput struct {
	ID          string `valid:"required,length(1|64)"`
	Name        string `valid:"required,length(1|64)"`
	OwnerUserID string `valid:"required,length(1|64)"`
}

type CreateGameOutput struct {
	Game
}

type DescribeGameInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeGameOutput struct {
	Game
}

type ListGameInput struct{}

type ListGameOutput struct {
	Items []Game
}

type UpdateGameInput struct {
	ID          string `valid:"required,length(1|64)"`
	Name        string `valid:"required,length(1|64)"`
	OwnerUserID string `valid:"required,length(1|64)"`
}

type UpdateGameOutput struct {
	Game
}

type RemoveGameInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveGameOutput struct {
	Game
}

type JoinGameInput struct {
	ID     string `valid:"required,length(1|64)"`
	UserID string `valid:"required,length(1|64)"`
}

type JoinGameOutput struct {
	Game
}

/* TYPES FOR STORAGE */

type AddInput struct {
	ID           string
	Name         string
	OwnerUserID  string
	CreationDate time.Time
	LastUpdate   time.Time
}

type AddOutput struct {
	Game
}

type GetInput struct {
	ID string
}

type GetOutput struct {
	Game
}

type AllInput struct{}

type AllOutput struct {
	Items []Game
}

type EditInput struct {
	ID          string
	Name        string
	OwnerUserID string
	LastUpdate  time.Time
}

type EditOutput struct {
	Game
}

type RemoveInput struct {
	ID string
}

type RemoveOutput struct {
	Game
}

/* COLLECTION TYPE */

type GameCollection []Game

/* INTERFACES */

type GameUseCase interface {
	CreateGame(ctx context.Context, input CreateGameInput) (*CreateGameOutput, error)
	DescribeGame(ctx context.Context, input DescribeGameInput) (*DescribeGameOutput, error)
	ListGame(ctx context.Context, input ListGameInput) (*ListGameOutput, error)
	UpdateGame(ctx context.Context, input UpdateGameInput) (*UpdateGameOutput, error)
	RemoveGame(ctx context.Context, input RemoveGameInput) (*RemoveGameOutput, error)
	JoinGame(ctx context.Context, input JoinGameInput) (*JoinGameOutput, error)
}

type GameStorage interface {
	Add(ctx context.Context, input AddInput) (*AddOutput, error)
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	All(ctx context.Context, input AllInput) (*AllOutput, error)
	Edit(ctx context.Context, input EditInput) (*EditOutput, error)
	Remove(ctx context.Context, input RemoveInput) (*RemoveOutput, error)
}

/* INSTANCE TYPES */

type DefaultGameUseCase struct {
	gameStorage  GameStorage
	userUseCase  user.UserUseCase
	betUseCase   bet.BetUseCase
	holdContract holdContract.HoldContractAdapter
}

type DefaultGameUseCaseOptions struct {
	GameStorage  GameStorage
	UserUseCase  user.UserUseCase
	BetUseCase   bet.BetUseCase
	HoldContract holdContract.HoldContractAdapter
}

/* CONSTRUCTOR TO CREATE NEW INSTANCES */

func NewDefaultGameUseCase(options DefaultGameUseCaseOptions) (*DefaultGameUseCase, error) {
	if options.GameStorage == nil {
		return nil, errors.New("cannot create an instance of GameWalletUseCase because provided GameWalletStorage is nil")
	}

	if options.UserUseCase == nil {
		return nil, errors.New("cannot create an instance of GameUseCase because provided UserUseCase is nil")
	}

	if options.BetUseCase == nil {
		return nil, errors.New("cannot create an instance of GameUseCase because provided BetUseCase is nil")
	}

	if options.HoldContract == nil {
		return nil, errors.New("cannot create an instance of GameUseCase because provided HoldContract is nil")
	}

	return &DefaultGameUseCase{
		gameStorage:  options.GameStorage,
		userUseCase:  options.UserUseCase,
		betUseCase:   options.BetUseCase,
		holdContract: options.HoldContract,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ GameUseCase = (*DefaultGameUseCase)(nil)

/* IMPLEMENTATIONS */

func (useCase *DefaultGameUseCase) CreateGame(ctx context.Context, input CreateGameInput) (*CreateGameOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//CRUD validation

	var checkGame GetInput
	checkGame.ID = input.ID
	_, err = useCase.gameStorage.Get(ctx, checkGame)
	if err == nil {
		return nil, err
	}

	//map desired game with input
	var desiredGame AddInput
	desiredGame.ID = input.ID
	desiredGame.Name = input.Name
	desiredGame.OwnerUserID = input.OwnerUserID
	desiredGame.CreationDate = time.Now()
	desiredGame.LastUpdate = desiredGame.CreationDate

	storedGame, err := useCase.gameStorage.Add(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	return &CreateGameOutput{
		Game: storedGame.Game,
	}, err
}

func (useCase *DefaultGameUseCase) DescribeGame(ctx context.Context, input DescribeGameInput) (*DescribeGameOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired game with input
	var desiredGame GetInput
	desiredGame.ID = input.ID

	storedGame, err := useCase.gameStorage.Get(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	return &DescribeGameOutput{
		Game: storedGame.Game,
	}, err
}

func (useCase *DefaultGameUseCase) ListGame(ctx context.Context, input ListGameInput) (*ListGameOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired game with input
	var desiredGame AllInput

	gameList, err := useCase.gameStorage.All(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	var items = gameList.Items

	return &ListGameOutput{
		Items: items,
	}, err
}

func (useCase *DefaultGameUseCase) UpdateGame(ctx context.Context, input UpdateGameInput) (*UpdateGameOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired game with input
	var desiredGame EditInput
	desiredGame.ID = input.ID
	desiredGame.Name = input.Name
	desiredGame.OwnerUserID = input.OwnerUserID
	desiredGame.LastUpdate = time.Now()

	storedGame, err := useCase.gameStorage.Edit(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	return &UpdateGameOutput{
		Game: storedGame.Game,
	}, err
}

func (useCase *DefaultGameUseCase) RemoveGame(ctx context.Context, input RemoveGameInput) (*RemoveGameOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired game with input
	var desiredGame RemoveInput
	desiredGame.ID = input.ID

	storedGame, err := useCase.gameStorage.Remove(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	return &RemoveGameOutput{
		Game: storedGame.Game,
	}, err
}

func (useCase *DefaultGameUseCase) JoinGame(ctx context.Context, input JoinGameInput) (*JoinGameOutput, error) {

	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired game with input
	var desiredGame DescribeGameInput
	desiredGame.ID = input.ID

	var desiredUser user.DescribeUserInput
	desiredUser.ID = input.UserID

	storedGame, err := useCase.DescribeGame(ctx, desiredGame)
	if err != nil {
		return nil, err
	}

	_ = storedGame

	storedUser, err := useCase.userUseCase.DescribeUser(ctx, desiredUser)
	if err != nil {
		return nil, err
	}

	_ = storedUser

	//todo get number of players
	useCase.betUseCase.ListBet()

	if n == 4 {
		return nil, errors.New("the game is fulfilled by 4 players")
	}

	//create hold
	var ownerWallet string
	var userWallet string
	var holdID int64
	//todo get this vars

	useCase.holdContract.HoldFrom(ctx, userWallet, ownerWallet, 5, holdID)

	//create bet
	var newBet bet.CreateBetInput
	newBet.ID = input.ID + input.UserID
	newBet.Amount = 5
	newBet.GameID = input.ID
	newBet.UserID = input.UserID

	_, err = useCase.betUseCase.CreateBet(ctx, newBet)
	if err != nil {
		return nil, err
	}

	//todo update game ++player
	var editGame EditInput

	updatedGame, err := useCase.gameStorage.Edit(ctx, editGame)
	if err != nil {
		return nil, err
	}

	if n == 3 {
		//todo select winner
		//todo execute holds
	}

	return &JoinGameOutput{
		Game: updatedGame.Game,
	}, err
}
