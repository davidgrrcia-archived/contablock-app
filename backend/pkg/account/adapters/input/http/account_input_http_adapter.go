package accountInputHttpAdapter

import (
	"context"
	"errors"

	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
)

type AccountInputAdapter interface {
	AdaptDescribeAccount(ctx context.Context, data DescribeAccountRequest) (*DescribeAccountResponseWrapper, error)
	AdaptListAccounts(ctx context.Context, data ListAccountRequest) (*ListAccountResponseWrapper, error)
	AdaptCreateAccount(ctx context.Context, data CreateAccountRequest) (*CreateAccountResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.AccountUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		accountUseCase: options.AccountUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	accountUseCase account.AccountUseCase
}

type DefaultHttpInputAdapterOptions struct {
	AccountUseCase account.AccountUseCase
}

var _ AccountInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptDescribeAccount(ctx context.Context, data DescribeAccountRequest) (*DescribeAccountResponseWrapper, error) {
	var inputData account.DescribeAccountInput
	inputData.Project = data.Project
	inputData.ID = data.ID

	outputData, err := adapter.accountUseCase.DescribeAccount(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeAccountResponse
	response.Project = outputData.Account.Project
	response.ID = outputData.Account.ID
	response.Name = outputData.Account.Name
	response.Level = outputData.Account.Level
	response.BalanceID = outputData.Account.BalanceID
	response.DisplayName = outputData.Account.DisplayName

	responseWrapper := DescribeAccountResponseWrapper{}
	responseWrapper.DescribeAccountResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListAccounts(ctx context.Context, data ListAccountRequest) (*ListAccountResponseWrapper, error) {
	var inputData account.ListAccountInput
	inputData.Project = data.Project

	outputData, err := adapter.accountUseCase.ListAccounts(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListAccountResponse
	response.Items = make([]Group, 0)

	for _, item := range outputData.Items {
		var responseItem Group
		responseItem.Project = item.Project
		responseItem.ID = item.ID
		responseItem.Name = item.Name
		responseItem.Level = item.Level
		responseItem.BalanceID = item.BalanceID
		responseItem.DisplayName = item.DisplayName

		for i, subgroup := range item.Subgroups {
			var subgroupItem Subgroup
			subgroupItem.Project = subgroup.Project
			subgroupItem.ID = subgroup.ID
			subgroupItem.Name = subgroup.Name
			subgroupItem.Level = subgroup.Level
			subgroupItem.BalanceID = subgroup.BalanceID
			subgroupItem.DisplayName = subgroup.DisplayName
			responseItem.Subgroups = append(responseItem.Subgroups, subgroupItem)
			for _, account := range subgroup.Accounts {
				var accountItem Account
				accountItem.Project = account.Project
				accountItem.ID = account.ID
				accountItem.Name = account.Name
				accountItem.Level = account.Level
				accountItem.BalanceID = account.BalanceID
				accountItem.DisplayName = account.DisplayName
				responseItem.Subgroups[i].Accounts = append(responseItem.Subgroups[i].Accounts, accountItem)
			}
		}

		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListAccountResponseWrapper{}
	responseWrapper.ListAccountResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptCreateAccount(ctx context.Context, data CreateAccountRequest) (*CreateAccountResponseWrapper, error) {
	var inputData account.CreateAccountInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.Name = data.Name
	inputData.Level = data.Level
	inputData.BalanceID = data.BalanceID

	outputData, err := adapter.accountUseCase.CreateAccount(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateAccountResponse
	response.Project = outputData.Account.Project
	response.ID = outputData.Account.ID
	response.Name = outputData.Account.Name
	response.Level = outputData.Account.Level
	response.BalanceID = outputData.Account.BalanceID
	response.DisplayName = outputData.Account.DisplayName

	responseWrapper := CreateAccountResponseWrapper{}
	responseWrapper.CreateAccountResponse = response

	return &responseWrapper, nil
}