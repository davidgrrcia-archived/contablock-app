package account

import (
	"context"
	"errors"

	"github.com/asaskevich/govalidator"
)

type AccountUseCase interface {
	DescribeAccount(ctx context.Context, input DescribeAccountInput) (*DescribeAccountOutput, error)
	ListAccounts(ctx context.Context, input ListAccountInput) (*ListAccountOutput, error)
}

func NewDefaultAccountUseCase(options DefaultAccountUseCaseOptions) (*DefaultAccountUseCase, error) {
	if options.AccountStorage == nil {
		return nil, errors.New("cannot create an instance of AccountUseCase because provided AccountStorage is nil")
	}

	return &DefaultAccountUseCase{
		accountStorage: options.AccountStorage,
	}, nil
}

type DefaultAccountUseCase struct {
	accountStorage AccountStorage
}

type DefaultAccountUseCaseOptions struct {
	AccountStorage AccountStorage
}

var _ AccountUseCase = (*DefaultAccountUseCase)(nil)

func (useCase *DefaultAccountUseCase) DescribeAccount(ctx context.Context, input DescribeAccountInput) (*DescribeAccountOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedAccount, err := useCase.accountStorage.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	var output DescribeAccountOutput
	output.ID = storedAccount.ID
	output.Name = storedAccount.Name
	output.Level = storedAccount.Level
	output.Type = storedAccount.Type
	output.DisplayName = storedAccount.ID + " - " + storedAccount.Name

	return &output, err
}

func (useCase *DefaultAccountUseCase) ListAccounts(ctx context.Context, input ListAccountInput) (*ListAccountOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input")
	}

	storedAccounts, err := useCase.accountStorage.All(ctx)
	if err != nil {
		return nil, err
	}

	var items []Account
	for _, item := range storedAccounts.Items {
		var account Account
		account.ID = item.ID
		account.Name = item.Name
		account.Level = item.Level
		account.Type = item.Type
		account.DisplayName = item.ID + " - " + item.Name

		items = append(items, account)
	}

	return &ListAccountOutput{
		Items: items,
	}, err
}
