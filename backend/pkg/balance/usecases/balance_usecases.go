package balance

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
)

type BalanceUseCase interface {
	CreateBalanceByYear(ctx context.Context, input CreateBalanceByYearInput) (*CreateBalanceByYearOutput, error)
	ListBalances(ctx context.Context, input ListBalanceInput) (*ListBalanceOutput, error)
	RemoveBalance(ctx context.Context, input RemoveBalanceInput) (*RemoveBalanceOutput, error)
	CreateBalance(ctx context.Context, input CreateBalanceInput) (*CreateBalanceOutput, error)
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

func (useCase *DefaultBalanceUseCase) CreateBalanceByYear(ctx context.Context, input CreateBalanceByYearInput) (*CreateBalanceByYearOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	id := uuid.New().String()

	var desiredBalance BalanceByYearStorageData
	if input.ID == nil {
		input.ID = &id
	}
	desiredBalance.Project = input.Project
	desiredBalance.ID = *input.ID
	desiredBalance.BalanceID = input.BalanceID
	desiredBalance.Amount = input.Amount
	desiredBalance.Year = input.Year
	storedBalance, err := useCase.balanceStorage.AddByYear(ctx, desiredBalance)
	if err != nil {
		return nil, err
	}

	var output CreateBalanceByYearOutput
	output.Project = storedBalance.Project
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
	storedBalancesCurrentYear, err := useCase.balanceStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	previousYear := *input.Year - 1
	filters.Year = &previousYear
	storedBalancesPreviousYear, err := useCase.balanceStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	var balances []Balance
	for _, item := range storedBalancesCurrentYear.Items {
		var balance Balance
		balance.Project = item.Project
		balance.ID = item.ID
		balance.Type = item.Type
		balance.Kind = item.Kind
		balance.Parent = item.Parent
		balance.Child = item.Child
		balance.CurrentYear = item.Year
		balance.PreviousYear = previousYear
		balance.CurrentAmount = item.Amount

		balances = append(balances, balance)
	}

	for _, item := range storedBalancesPreviousYear.Items {
		alreadyAdded := false

		for j, balance := range balances {
			if item.ID == balance.ID {
				balances[j].PreviousAmount = item.Amount
				alreadyAdded = true
			}
		}

		if !alreadyAdded {
			var balance Balance
			balance.Project = item.Project
			balance.ID = item.ID
			balance.Type = item.Type
			balance.Kind = item.Kind
			balance.Parent = item.Parent
			balance.Child = item.Child
			balance.CurrentYear = previousYear + 1
			balance.PreviousYear = previousYear
			balance.PreviousAmount = item.Amount

			balances = append(balances, balance)
		}
	}

	return &ListBalanceOutput{
		Items: balances,
	}, err
}

func (useCase *DefaultBalanceUseCase) RemoveBalance(ctx context.Context, input RemoveBalanceInput) (*RemoveBalanceOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedBalance, err := useCase.balanceStorage.Remove(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	var output RemoveBalanceOutput
	output.Project = storedBalance.Project
	output.ID = storedBalance.ID
	output.BalanceID = storedBalance.BalanceID
	output.Amount = storedBalance.Amount
	output.Year = storedBalance.Year

	return &output, nil
}

func (useCase *DefaultBalanceUseCase) CreateBalance(ctx context.Context, input CreateBalanceInput) (*CreateBalanceOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedBalance, err := useCase.balanceStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedBalance != nil {
		return nil, errors.New("balance [" + storedBalance.ID + "] already exists")
	}

	var desiredBalance AddBalanceStorageData
	desiredBalance.Project = input.Project
	desiredBalance.ID = input.ID
	desiredBalance.Type = input.Type
	desiredBalance.Kind = input.Kind
	desiredBalance.Parent = input.Parent
	desiredBalance.Child = input.Child
	desiredBalance.Accounts = input.Accounts

	addedBalance, err := useCase.balanceStorage.Add(ctx, desiredBalance)
	if err != nil {
		return nil, err
	}

	var output CreateBalanceOutput
	output.Project = addedBalance.Project
	output.ID = addedBalance.ID
	output.Type = addedBalance.Type
	output.Kind = addedBalance.Kind
	output.Parent = addedBalance.Parent
	output.Child = addedBalance.Child
	output.Accounts = addedBalance.Accounts

	return &output, nil
}
