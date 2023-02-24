package account

import (
	"context"
	"errors"
	"strconv"

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
		return nil, errors.New("invalid input")
	}

	storedAccount, err := useCase.accountStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
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
		return nil, errors.New("invalid input")
	}

	storedAccounts, err := useCase.accountStorage.All(ctx, input.Project)
	if err != nil {
		return nil, err
	}

	var groups []Group
	var accounts []Account

	for _, item := range storedAccounts.Items {
		if item.Level == LevelGroup {
			var group Group
			group.Project = item.Project
			group.ID = item.ID
			group.Name = item.Name
			group.Level = item.Level
			group.BalanceID = item.BalanceID
			group.DisplayName = item.ID + " - " + item.Name
			groups = append(groups, group)
		}
		if item.Level == LevelSubGroup {
			var subgroup Subgroup
			subgroup.Project = item.Project
			subgroup.ID = item.ID
			subgroup.Name = item.Name
			subgroup.Level = item.Level
			subgroup.BalanceID = item.BalanceID
			subgroup.DisplayName = item.ID + " - " + item.Name
			indexString := subgroup.ID[0:1]
			index, err := strconv.Atoi(indexString)
			if err != nil {
				return nil, err
			}
			groups[index-1].Subgroups = append(groups[index-1].Subgroups, subgroup)
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
			if group.ID != account.ID[0:1] {
				continue
			}

			for j, subgroup := range group.Subgroups {
				if subgroup.ID != account.ID[0:2] {
					continue
				}

				groups[i].Subgroups[j].Accounts = append(groups[i].Subgroups[j].Accounts, account)
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
		return nil, errors.New("invalid input")
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
