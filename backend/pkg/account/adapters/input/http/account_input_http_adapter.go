package accountInputHttpAdapter

import (
	"context"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
)

type AccountInputAdapter interface {
	AdaptDescribeAccount(ctx context.Context, data DescribeAccountRequest) (*DescribeAccountResponseWrapper, error)
	AdaptListAccounts(ctx context.Context, data ListAccountRequest) (*ListAccountResponseWrapper, error)
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
	inputData.ID = data.ID

	outputData, err := adapter.accountUseCase.DescribeAccount(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeAccountResponse
	response.ID = outputData.Account.ID
	response.Name = outputData.Account.Name
	response.Level = outputData.Account.Level
	response.Type = outputData.Account.Type
	response.DisplayName = outputData.Account.DisplayName

	responseWrapper := DescribeAccountResponseWrapper{}
	responseWrapper.DescribeAccountResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListAccounts(ctx context.Context, data ListAccountRequest) (*ListAccountResponseWrapper, error) {
	var inputData account.ListAccountInput

	outputData, err := adapter.accountUseCase.ListAccounts(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListAccountResponse
	response.Items = make([]Group, 0)

	for _, item := range outputData.Items {
		var responseItem Group
		responseItem.ID = item.ID
		responseItem.Name = item.Name
		responseItem.Level = item.Level
		responseItem.Type = item.Type
		responseItem.DisplayName = item.DisplayName

		for i, subgroup := range item.Subgroups {
			var subgroupItem Subgroup
			subgroupItem.ID = subgroup.ID
			subgroupItem.Name = subgroup.Name
			subgroupItem.Level = subgroup.Level
			subgroupItem.Type = subgroup.Type
			subgroupItem.DisplayName = subgroup.DisplayName
			responseItem.Subgroups = append(responseItem.Subgroups, subgroupItem)
			for _, account := range subgroup.Accounts {
				var accountItem Account
				accountItem.ID = account.ID
				accountItem.Name = account.Name
				accountItem.Level = account.Level
				accountItem.Type = account.Type
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
