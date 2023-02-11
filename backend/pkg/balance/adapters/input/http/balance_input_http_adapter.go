package balanceInputHttpAdapter

import (
	"context"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
)

type BalanceInputAdapter interface {
	AdaptListBalances(ctx context.Context, data ListBalanceRequest) (*ListBalanceResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.BalanceUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		balanceUseCase: options.BalanceUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	balanceUseCase balance.BalanceUseCase
}

type DefaultHttpInputAdapterOptions struct {
	BalanceUseCase balance.BalanceUseCase
}

var _ BalanceInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptListBalances(ctx context.Context, data ListBalanceRequest) (*ListBalanceResponseWrapper, error) {
	var inputData balance.ListBalanceInput
	inputData.Year = &data.Year
	outputData, err := adapter.balanceUseCase.ListBalances(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListBalanceResponse
	response.Items = make([]Balance, 0)

	for _, item := range outputData.Items {
		var responseItem Balance
		responseItem.ID = item.ID
		responseItem.Type = item.Type
		responseItem.Kind = item.Kind
		responseItem.Parent = item.Parent
		responseItem.Child = item.Child
		responseItem.CurrentYear = item.CurrentYear
		responseItem.PreviousYear = item.PreviousYear
		responseItem.CurrentAmount = item.CurrentAmount
		responseItem.PreviousAmount = item.PreviousAmount

		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListBalanceResponseWrapper{}
	responseWrapper.ListBalanceResponse = response
	return &responseWrapper, nil
}
