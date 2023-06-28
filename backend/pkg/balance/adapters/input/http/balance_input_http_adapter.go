package balanceInputHttpAdapter

import (
	"context"
	"errors"
	"strconv"

	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
)

type BalanceInputAdapter interface {
	AdaptCreateBalance(ctx context.Context, data CreateBalanceRequest) (*CreateBalanceResponseWrapper, error)
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
	inputData.Project = data.Project
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
			typeBalance.Accounts = mapping[item.Type].Accounts
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
				kindBalance.Accounts = item.Accounts
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
					parentBalance.Accounts = item.Accounts
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
				childBalance.Accounts = item.Accounts
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
				parentBalance.Accounts = item.Accounts
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
				childBalance.Accounts = item.Accounts
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
	response.CurrentDebitTotalAmount = getCurrentDebitAmount(response.Items)
	response.PreviousDebitTotalAmount = getPreviousDebitAmount(response.Items)
	response.CurrentCreditTotalAmount = getCurrentCreditAmount(response.Items)
	response.PreviousCreditTotalAmount = getPreviousCreditAmount(response.Items)

	responseWrapper := ListBalanceResponseWrapper{}
	responseWrapper.ListBalanceResponse = response
	return &responseWrapper, nil
}

func mapToList(mapa map[string]Balance) []BalanceItem {
	lista := []BalanceItem{}
	for _, valor := range mapa {
		item := BalanceItem{
			Project:        valor.Project,
			ID:             valor.ID,
			Name:           valor.Name,
			CurrentYear:    valor.CurrentYear,
			PreviousYear:   valor.PreviousYear,
			CurrentAmount:  valor.CurrentAmount,
			PreviousAmount: valor.PreviousAmount,
			Child:          nil,
			Accounts:       valor.Accounts,
		}
		if valor.Map != nil {
			item.Child = mapToList(valor.Map)
		}
		lista = append(lista, item)
	}
	return lista
}

func getCurrentDebitAmount(balance []BalanceItem) float64 {
	var amount float64

	for _, item := range balance {
		if item.Name[0:2] == "A)" || item.Name[0:2] == "B)" {
			amount += item.CurrentAmount
		}
	}

	return amount
}

func getPreviousDebitAmount(balance []BalanceItem) float64 {
	var amount float64

	for _, item := range balance {
		if item.Name[0:2] == "A)" || item.Name[0:2] == "B)" {
			amount += item.PreviousAmount
		}
	}

	return amount
}

func getCurrentCreditAmount(balance []BalanceItem) float64 {
	var amount float64

	for _, item := range balance {
		if item.Name[0:2] == "C)" || item.Name[0:2] == "D)" || item.Name[0:2] == "E)" {
			amount += item.CurrentAmount
		}
	}

	return amount
}

func getPreviousCreditAmount(balance []BalanceItem) float64 {
	var amount float64

	for _, item := range balance {
		if item.Name[0:2] == "C)" || item.Name[0:2] == "D)" || item.Name[0:2] == "E)" {
			amount += item.PreviousAmount
		}
	}

	return amount
}

func (adapter *DefaultHttpInputAdapter) AdaptCreateBalance(ctx context.Context, data CreateBalanceRequest) (*CreateBalanceResponseWrapper, error) {
	var inputData balance.CreateBalanceInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.Type = data.Type
	inputData.Kind = data.Kind
	inputData.Parent = data.Parent
	inputData.Child = data.Child
	inputData.Accounts = data.Accounts

	outputData, err := adapter.balanceUseCase.CreateBalance(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateBalanceResponse
	response.Project = outputData.Project
	response.ID = outputData.ID
	response.Type = outputData.Type
	response.Kind = outputData.Kind
	response.Parent = outputData.Parent
	response.Child = outputData.Child
	response.Accounts = outputData.Accounts

	responseWrapper := CreateBalanceResponseWrapper{}
	responseWrapper.CreateBalanceResponse = response

	return &responseWrapper, nil
}
