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

	typesMap := make(map[string]Balance)

	for _, item := range outputData.Items {
		kindsMap := make(map[string]Balance)
		parentsMap := make(map[string]Balance)

		var balance Balance
		balance.ID = item.ID
		balance.CurrentYear = item.CurrentYear
		balance.PreviousYear = item.PreviousYear
		balance.CurrentAmount = item.CurrentAmount
		balance.PreviousAmount = item.PreviousAmount
		balance.DifferenceAmount = item.DifferenceAmount
		balance.Increase = item.Increase

		if item.Kind != nil && *item.Kind != "" {
			balance.Name = *item.Kind
			typesMap[item.Type] = balance

			if item.Parent != nil && *item.Parent != "" {
				balance.Name = *item.Parent
				if balance.Map == nil {
					kindsMap[*item.Kind] = balance
					balance.Map = kindsMap
				} else {
					balance.Map[*item.Kind] = balance
				}
				typesMap[item.Type] = balance

				if item.Child != nil && *item.Child != "" {
					balance.Name = *item.Child
					parentsMap[*item.Parent] = balance
					balance.Map = parentsMap //todo review
					typesMap[item.Type] = balance
				}
			}
			continue
		}
		if item.Parent != nil && *item.Parent != "" {
			balance.Name = *item.Parent
			typesMap[item.Type] = balance

			if item.Child != nil && *item.Child != "" {
				balance.Name = *item.Child
				if typesMap[item.Type].Map == nil {
					parentsMap[*item.Parent] = balance
					balance.Map = parentsMap
				} else {
					balance.Map[*item.Parent] = balance
				}
				balance.Name = *item.Parent
				typesMap[item.Type] = balance
			}
			continue
		}
		if item.Child != nil && *item.Child != "" {
			balance.Name = *item.Child
			typesMap[item.Type] = balance
			continue
		}
	}

	var response ListBalanceResponse
	response.Items = make([]Type, 0)

	for _, balance := range outputData.Items {
		var child Child
		var parent Parent
		var kind Kind
		var item Type

		child.Balance.ID = balance.ID
		child.Balance.Name = *balance.Child
		child.Balance.CurrentYear = balance.CurrentYear
		child.Balance.PreviousYear = balance.PreviousYear
		child.Balance.CurrentAmount = balance.CurrentAmount
		child.Balance.PreviousAmount = balance.PreviousAmount
		child.Balance.DifferenceAmount = balance.DifferenceAmount
		child.Balance.Increase = balance.Increase

		if balance.Child != nil {
			child.Balance.ID = *balance.Child
			if balance.Parent != nil {
				parent.Balance.ID = *balance.Parent
				parent.Childs = []Child{child}
				if balance.Kind != nil {
					kind.Balance.ID = *balance.Kind
					kind.Parents = []Parent{parent}
					item.Balance.ID = balance.ID
					item.Kinds = []Kind{kind}
					response.Items = append(response.Items, item)
				}
			}
		}
	}

	responseWrapper := ListBalanceResponseWrapper{}
	responseWrapper.ListBalanceResponse = response
	return &responseWrapper, nil
}
