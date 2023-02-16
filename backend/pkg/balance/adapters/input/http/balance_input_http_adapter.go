package balanceInputHttpAdapter

import (
	"context"
	"errors"
	"strconv"

	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
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
			var kindBalance Balance
			kindBalance.Name = *item.Kind
			if mapping[item.Type].Map[*item.Kind].Map == nil {
				kindBalance.Map = make(map[string]Balance)
				kindBalance.CurrentAmount = item.CurrentAmount
				kindBalance.PreviousAmount = item.PreviousAmount
				mapping[item.Type].Map[*item.Kind] = kindBalance
			} else {
				kindBalance.Map = mapping[item.Type].Map[*item.Kind].Map
				kindBalance.CurrentAmount = mapping[item.Type].Map[*item.Kind].CurrentAmount + item.CurrentAmount
				kindBalance.PreviousAmount = mapping[item.Type].Map[*item.Kind].PreviousAmount + item.PreviousAmount
				mapping[item.Type].Map[*item.Kind] = kindBalance
			}

			if item.Parent != nil && *item.Parent != "" {
				var parentBalance Balance
				parentBalance.Name = *item.Parent
				if mapping[item.Type].Map[*item.Kind].Map[*item.Parent].Map == nil {
					parentBalance.Map = make(map[string]Balance)
					parentBalance.CurrentAmount = item.CurrentAmount
					parentBalance.PreviousAmount = item.PreviousAmount
					mapping[item.Type].Map[*item.Kind].Map[*item.Parent] = parentBalance
				} else {
					parentBalance.Map = mapping[item.Type].Map[*item.Kind].Map[*item.Parent].Map
					parentBalance.CurrentAmount = mapping[item.Type].Map[*item.Kind].Map[*item.Parent].CurrentAmount + item.CurrentAmount
					parentBalance.PreviousAmount = mapping[item.Type].Map[*item.Kind].Map[*item.Parent].PreviousAmount + item.PreviousAmount
					mapping[item.Type].Map[*item.Kind].Map[*item.Parent] = parentBalance
				}
			}

			if item.Child != nil && *item.Child != "" {
				var childBalance Balance
				childBalance.Name = *item.Child
				childBalance.CurrentAmount = item.CurrentAmount
				childBalance.PreviousAmount = item.PreviousAmount
				mapping[item.Type].Map[*item.Kind].Map[*item.Parent].Map[*item.Child] = childBalance
			}
			continue
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
	response.Items = mapToList(mapping)

	responseWrapper := ListBalanceResponseWrapper{}
	responseWrapper.ListBalanceResponse = response
	return &responseWrapper, nil
}

func mapToList(mapa map[string]Balance) []BalanceItem {
	lista := []BalanceItem{}
	for _, valor := range mapa {
		item := BalanceItem{
			ID:             valor.ID,
			Name:           valor.Name,
			CurrentYear:    valor.CurrentYear,
			PreviousYear:   valor.PreviousYear,
			CurrentAmount:  valor.CurrentAmount,
			PreviousAmount: valor.PreviousAmount,
			Child:          nil,
		}
		if valor.Map != nil {
			item.Child = mapToList(valor.Map)
		}
		lista = append(lista, item)
	}
	return lista
}

