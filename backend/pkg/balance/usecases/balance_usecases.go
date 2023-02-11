package balance

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
)

type BalanceUseCase interface {
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

func (useCase *DefaultBalanceUseCase) ListBalances(ctx context.Context, input ListBalanceInput) (*ListBalanceOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedBalances, err := useCase.balanceStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	var balances []Balance
	for _, item := range storedBalances.Items {
		var balance Balance
		balance.ID = item.ID
		balance.Type = item.Type
		balance.Kind = item.Kind
		balance.Parent = item.Parent
		balance.Child = item.Child
		balance.Accounts = item.Accounts
		balance.CurrentYear = item.CurrentYear
		balance.PreviousYear = item.PreviousYear
		balance.CurrentAmount = item.CurrentAmount
		balance.PreviousAmount = item.PreviousAmount

		balances = append(balances, balance)
	}

	return &ListBalanceOutput{
		Items: balances,
	}, err
}
