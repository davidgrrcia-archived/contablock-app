package account

import (
	"context"
	"errors"
	"github.com/asaskevich/govalidator"
)

type AccountUseCase interface {
	DescribeAccount(ctx context.Context, input DescribeAccountInput) (*DescribeAccountOutput, error)
	ListAccounts(ctx context.Context, input ListAccountInput) (*ListAccountOutput, error)
	CreateAccount(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error)
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
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedAccount, err := useCase.accountStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedAccount == nil {
		return nil, errors.New("account not found")
	}

	var output DescribeAccountOutput
	output.Account.Project = storedAccount.Project
	output.Account.ID = storedAccount.ID
	output.Account.Name = storedAccount.Name
	output.Account.Level = storedAccount.Level
	output.Account.BalanceID = storedAccount.BalanceID
	output.Account.DisplayName = storedAccount.ID + " - " + storedAccount.Name

	return &output, err
}

func (useCase *DefaultAccountUseCase) ListAccounts(ctx context.Context, input ListAccountInput) (*ListAccountOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedAccounts, err := useCase.accountStorage.All(ctx, input.Project)
	if err != nil {
		return nil, err
	}

	var groups []Group
	var accounts []Account

	for _, item := range storedAccounts.Items {
		if item.Level == LevelGroup || item.Level == LevelSubGroup {
			var group Group
			group.Project = item.Project
			group.ID = item.ID
			group.Name = item.Name
			group.Level = item.Level
			group.BalanceID = item.BalanceID
			group.DisplayName = item.ID + " - " + item.Name
			groups = append(groups, group)
		}
		if item.Level == LevelAccount || item.Level == LevelSubAccount {
			var account Account
			account.Project = item.Project
			account.ID = item.ID
			account.Name = item.Name
			account.Level = item.Level
			account.BalanceID = item.BalanceID
			account.DisplayName = item.ID + " - " + item.Name
			accounts = append(accounts, account)
		}
	}

	for _, account := range accounts {
		for i, group := range groups {
			if len(group.ID) == 1 {
				continue
			}

			if group.ID == account.ID[0:2] {
				groups[i].Accounts = append(groups[i].Accounts, account)
			}

		}
	}

	return &ListAccountOutput{
		Items: groups,
	}, err
}

func (useCase *DefaultAccountUseCase) CreateAccount(ctx context.Context, input CreateAccountInput) (*CreateAccountOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	//TODO if the project in the path is the admin, then allow to add

	storedAccount, err := useCase.accountStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedAccount != nil {
		return nil, errors.New("account [" + storedAccount.ID + "] already exists")
	}

	var desiredAccount AccountStorageData
	desiredAccount.Project = input.Project
	desiredAccount.ID = input.ID
	desiredAccount.Name = input.Name
	desiredAccount.Level = input.Level
	desiredAccount.BalanceID = input.BalanceID

	addedAccount, err := useCase.accountStorage.Add(ctx, desiredAccount)
	if err != nil {
		return nil, err
	}

	var output CreateAccountOutput
	output.Account.Project = addedAccount.Project
	output.Account.ID = addedAccount.ID
	output.Account.Name = addedAccount.Name
	output.Account.Level = addedAccount.Level
	output.Account.BalanceID = addedAccount.BalanceID
	output.Account.DisplayName = addedAccount.ID + " - " + addedAccount.Name

	return &output, nil
}
