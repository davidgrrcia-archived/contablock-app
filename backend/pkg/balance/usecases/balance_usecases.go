package balance

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type BalanceUseCase interface {
	CreateBalance(ctx context.Context, input CreateBalanceInput) (*CreateBalanceOutput, error)
	ListBalances(ctx context.Context, input ListBalanceInput) (*ListBalanceOutput, error)
}

func NewDefaultBalanceUseCase(options DefaultBalanceUseCaseOptions) (*DefaultBalanceUseCase, error) {
	if options.BalanceStorage == nil {
		return nil, errors.New("cannot create an instance of BalanceUseCase because provided BalanceStorage is nil")
	}

	return &DefaultBalanceUseCase{
		balanceStorage: options.BalanceStorage,
	}, nil
}

type DefaultBalanceUseCase struct {
	balanceStorage BalanceStorage
}

type DefaultBalanceUseCaseOptions struct {
	BalanceStorage BalanceStorage
}

var _ BalanceUseCase = (*DefaultBalanceUseCase)(nil)

func (useCase *DefaultBalanceUseCase) CreateBalance(ctx context.Context, input CreateBalanceInput) (*CreateBalanceOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	id := uuid.New().String()

	var desiredBalance BalanceByYearStorageData
	if input.ID == nil {
		input.ID = &id
	}
	desiredBalance.ID = *input.ID
	desiredBalance.BalanceID = input.BalanceID
	desiredBalance.Amount = input.Amount
	desiredBalance.Year = input.Year
	storedBalance, err := useCase.balanceStorage.Add(ctx, desiredBalance)
	if err != nil {
		return nil, err
	}

	var output CreateBalanceOutput
	output.ID = storedBalance.ID
	output.BalanceID = storedBalance.BalanceID
	output.Amount = storedBalance.Amount
	output.Year = storedBalance.Year

	return &output, nil
}

func (useCase *DefaultBalanceUseCase) ListBalances(ctx context.Context, input ListBalanceInput) (*ListBalanceOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	var filters Filters
	filters.Year = input.Year
	storedBalancesCurrentYear, err := useCase.balanceStorage.All(ctx, filters)
	if err != nil {
		return nil, err
	}

	previousYear := *input.Year - 1
	filters.Year = &previousYear
	storedBalancesPreviousYear, err := useCase.balanceStorage.All(ctx, filters)
	if err != nil {
		return nil, err
	}

	_ = storedBalancesPreviousYear

	var balances []Balance
	for _, item := range storedBalancesCurrentYear.Items {
		var balance Balance
		balance.ID = item.ID
		balance.Type = item.Type
		balance.Kind = item.Kind
		balance.Parent = item.Parent
		balance.Child = item.Child

		balances = append(balances, balance)
	}

	return &ListBalanceOutput{
		Items: balances,
	}, err
}
