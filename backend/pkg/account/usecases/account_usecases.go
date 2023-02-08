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
	output.Account.ID = storedAccount.ID
	output.Account.Name = storedAccount.Name
	output.Account.Level = storedAccount.Level
	output.Account.Type = storedAccount.Type
	output.Account.DisplayName = storedAccount.ID + " - " + storedAccount.Name

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

	var groups []Group
	var accounts []Account

	for _, item := range storedAccounts.Items {
		if item.Level == LevelGroup {
			var group Group
			group.ID = item.ID
			group.Name = item.Name
			group.Level = item.Level
			group.Type = item.Type
			group.DisplayName = item.ID + " - " + item.Name
			groups = append(groups, group)
		}
		if item.Level == LevelSubGroup {
			var subgroup Subgroup
			subgroup.ID = item.ID
			subgroup.Name = item.Name
			subgroup.Level = item.Level
			subgroup.Type = item.Type
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
			account.ID = item.ID
			account.Name = item.Name
			account.Level = item.Level
			account.Type = item.Type
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
