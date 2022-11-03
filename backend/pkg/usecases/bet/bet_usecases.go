package bet

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
	"time"
)

/* USE CASE TYPES */

type Bet struct {
	entities.Bet
	CreationDate time.Time
	LastUpdate   time.Time
}

type CreateBetInput struct {
	ID     string `valid:"required,length(1|64)"`
	Amount int64  `valid:"required,length(1|64)"`
	GameID string `valid:"required,length(1|64)"`
	UserID string `valid:"required,length(1|64)"`
}

type CreateBetOutput struct {
	Bet
}

type DescribeBetInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeBetOutput struct {
	Bet
}

type DescribeBetByUserIDInput struct {
	UserID string `valid:"required,length(1|64)"`
	GameID string `valid:"required,length(1|64)"`
}

type DescribeBetByUserIDOutput struct {
	Bet
}

type ListBetInput struct{}

type ListBetOutput struct {
	Items []Bet
}

type UpdateBetInput struct {
	ID     string `valid:"required,length(1|64)"`
	Amount int64  `valid:"required,length(1|64)"`
	GameID string `valid:"required,length(1|64)"`
	UserID string `valid:"required,length(1|64)"`
}

type UpdateBetOutput struct {
	Bet
}

type RemoveBetInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveBetOutput struct {
	Bet
}

/* TYPES FOR STORAGE */

type AddInput struct {
	ID           string
	Amount       int64
	GameID       string
	UserID       string
	CreationDate time.Time
	LastUpdate   time.Time
}

type AddOutput struct {
	Bet
}

type GetInput struct {
	ID string
}

type GetOutput struct {
	Bet
}

type GetByUserIDInput struct {
	UserID string
	GameID string
}

type GetByUserIDOutput struct {
	Bet
}

type AllInput struct{}

type AllOutput struct {
	Items []Bet
}

type EditInput struct {
	ID         string
	Amount     int64
	GameID     string
	UserID     string
	LastUpdate time.Time
}

type EditOutput struct {
	Bet
}

type RemoveInput struct {
	ID string
}

type RemoveOutput struct {
	Bet
}

/* COLLECTION TYPE */

type BetCollection []Bet

/* INTERFACES */

type BetUseCase interface {
	CreateBet(ctx context.Context, input CreateBetInput) (*CreateBetOutput, error)
	DescribeBet(ctx context.Context, input DescribeBetInput) (*DescribeBetOutput, error)
	DescribeBetByUserID(ctx context.Context, input DescribeBetByUserIDInput) (*DescribeBetByUserIDOutput, error)
	ListBet(ctx context.Context, input ListBetInput) (*ListBetOutput, error)
	UpdateBet(ctx context.Context, input UpdateBetInput) (*UpdateBetOutput, error)
	RemoveBet(ctx context.Context, input RemoveBetInput) (*RemoveBetOutput, error)
}

type BetStorage interface {
	Add(ctx context.Context, input AddInput) (*AddOutput, error)
	Get(ctx context.Context, input GetInput) (*GetOutput, error)
	GetByUserID(ctx context.Context, input GetByUserIDInput) (*GetByUserIDOutput, error)
	All(ctx context.Context, input AllInput) (*AllOutput, error)
	Edit(ctx context.Context, input EditInput) (*EditOutput, error)
	Remove(ctx context.Context, input RemoveInput) (*RemoveOutput, error)
}

/* INSTANCE TYPES */

type DefaultBetUseCase struct {
	betStorage BetStorage
}

type DefaultBetUseCaseOptions struct {
	BetStorage BetStorage
}

/* CONSTRUCTOR TO CREATE NEW INSTANCES */

func NewDefaultBetUseCase(options DefaultBetUseCaseOptions) (*DefaultBetUseCase, error) {
	if options.BetStorage == nil {
		return nil, errors.New("cannot create an instance of BetWalletUseCase because provided BetWalletStorage is nil")
	}

	return &DefaultBetUseCase{
		betStorage: options.BetStorage,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ BetUseCase = (*DefaultBetUseCase)(nil)

/* IMPLEMENTATIONS */

func (useCase *DefaultBetUseCase) CreateBet(ctx context.Context, input CreateBetInput) (*CreateBetOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//CRUD validation

	var checkBet GetInput
	checkBet.ID = input.ID
	_, err = useCase.betStorage.Get(ctx, checkBet)
	if err == nil {
		return nil, err
	}

	//map desired bet with input
	var desiredBet AddInput
	desiredBet.ID = input.ID
	desiredBet.Amount = input.Amount
	desiredBet.GameID = input.GameID
	desiredBet.UserID = input.UserID
	desiredBet.CreationDate = time.Now()
	desiredBet.LastUpdate = desiredBet.CreationDate

	storedBet, err := useCase.betStorage.Add(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	return &CreateBetOutput{
		Bet: storedBet.Bet,
	}, err
}

func (useCase *DefaultBetUseCase) DescribeBet(ctx context.Context, input DescribeBetInput) (*DescribeBetOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired bet with input
	var desiredBet GetInput
	desiredBet.ID = input.ID

	storedBet, err := useCase.betStorage.Get(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	return &DescribeBetOutput{
		Bet: storedBet.Bet,
	}, err
}

func (useCase *DefaultBetUseCase) DescribeBetByUserID(ctx context.Context, input DescribeBetByUserIDInput) (*DescribeBetByUserIDOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired bet with input
	var desiredBet GetByUserIDInput
	desiredBet.UserID = input.UserID
	desiredBet.GameID = input.GameID

	storedBet, err := useCase.betStorage.GetByUserID(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	return &DescribeBetByUserIDOutput{
		Bet: storedBet.Bet,
	}, err
}

func (useCase *DefaultBetUseCase) ListBet(ctx context.Context, input ListBetInput) (*ListBetOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired bet with input
	var desiredBet AllInput

	betList, err := useCase.betStorage.All(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	var items = betList.Items

	return &ListBetOutput{
		Items: items,
	}, err
}

func (useCase *DefaultBetUseCase) UpdateBet(ctx context.Context, input UpdateBetInput) (*UpdateBetOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired bet with input
	var desiredBet EditInput
	desiredBet.ID = input.ID
	desiredBet.Amount = input.Amount
	desiredBet.GameID = input.GameID
	desiredBet.UserID = input.UserID
	desiredBet.LastUpdate = time.Now()

	storedBet, err := useCase.betStorage.Edit(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	return &UpdateBetOutput{
		Bet: storedBet.Bet,
	}, err
}

func (useCase *DefaultBetUseCase) RemoveBet(ctx context.Context, input RemoveBetInput) (*RemoveBetOutput, error) {
	//Syntactic validation
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	//map desired bet with input
	var desiredBet RemoveInput
	desiredBet.ID = input.ID

	storedBet, err := useCase.betStorage.Remove(ctx, desiredBet)
	if err != nil {
		return nil, err
	}

	return &RemoveBetOutput{
		Bet: storedBet.Bet,
	}, err
}
