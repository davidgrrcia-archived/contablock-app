package balanceInputHttpAdapter

import (
	"context"
	"errors"
	"strconv"

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
	year, err := strconv.Atoi(data.Year)
	if err != nil {
		return nil, err
	}

	var inputData balance.ListBalanceInput
	inputData.Year = &year
	outputData, err := adapter.balanceUseCase.ListBalances(ctx, inputData)
	if err != nil {
		return nil, err
	}

	mapping := make(map[string]Balance)

	for _, item := range outputData.Items {
		var typeBalance Balance
		typeBalance.Name = item.Type
		typeBalance.CurrentYear = item.CurrentYear
		typeBalance.PreviousYear = item.PreviousYear
		if mapping[item.Type].Map == nil {
			typeBalance.Map = make(map[string]Balance)
			typeBalance.CurrentAmount = item.CurrentAmount
			typeBalance.PreviousAmount = item.PreviousAmount
			mapping[item.Type] = typeBalance
		} else {
			typeBalance.Map = mapping[item.Type].Map
			typeBalance.CurrentAmount = mapping[item.Type].CurrentAmount + item.CurrentAmount
			typeBalance.PreviousAmount = mapping[item.Type].PreviousAmount + item.PreviousAmount
			mapping[item.Type] = typeBalance
		}

		if item.Kind != nil && *item.Kind != "" {
			//balance.Name = *item.Kind
			//typesMap[item.Type] = balance
			//
			//if item.Parent != nil && *item.Parent != "" {
			//	parentBalance := Balance{
			//		Name: *item.Parent,
			//		Map:  make(map[string]Balance),
			//	}
			//	parentBalance.Map[*item.Kind] = balance
			//	typesMap[item.Type] = parentBalance
			//
			//	if item.Child != nil && *item.Child != "" {
			//		childBalance := Balance{
			//			Name: *item.Child,
			//			Map:  make(map[string]Balance),
			//		}
			//		childBalance.Map[*item.Parent] = parentBalance
			//		typesMap[item.Type] = childBalance
			//	}
			//}
			//continue
		}
		if item.Parent != nil && *item.Parent != "" {
			var parentBalance Balance
			parentBalance.Name = *item.Parent
			if mapping[item.Type].Map[*item.Parent].Map == nil {
				parentBalance.Map = make(map[string]Balance)
				parentBalance.CurrentAmount = item.CurrentAmount
				parentBalance.PreviousAmount = item.PreviousAmount
				mapping[item.Type].Map[*item.Parent] = parentBalance
			} else {
				parentBalance.Map = mapping[item.Type].Map[*item.Parent].Map
				parentBalance.CurrentAmount = mapping[item.Type].Map[*item.Parent].CurrentAmount + item.CurrentAmount
				parentBalance.PreviousAmount = mapping[item.Type].Map[*item.Parent].PreviousAmount + item.PreviousAmount
				mapping[item.Type].Map[*item.Parent] = parentBalance
			}

			if item.Child != nil && *item.Child != "" {
				var childBalance Balance
				childBalance.Name = *item.Child
				childBalance.CurrentAmount = item.CurrentAmount
				childBalance.PreviousAmount = item.PreviousAmount
				mapping[item.Type].Map[*item.Parent].Map[*item.Child] = childBalance
			}
			continue
		}
		if item.Child != nil && *item.Child != "" {
			childBalance := Balance{
				Name: *item.Child,
			}
			mapping[item.Type] = childBalance
			continue
		}
	}

	var response ListBalanceResponse
	response.Items = mapping

	responseWrapper := ListBalanceResponseWrapper{}
	responseWrapper.ListBalanceResponse = response
	return &responseWrapper, nil
}
