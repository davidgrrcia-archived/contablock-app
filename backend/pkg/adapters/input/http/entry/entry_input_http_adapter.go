package entryInputHttpAdapter

import (
	"context"
	"errors"
	accountInputHttpAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/account"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
	"strconv"
)

type EntryInputAdapter interface {
	AdaptCreateEntry(ctx context.Context, data CreateEntryRequest) (*CreateEntryResponseWrapper, error)
	AdaptDescribeEntry(ctx context.Context, data DescribeEntryRequest) (*DescribeEntryResponseWrapper, error)
	AdaptListEntries(ctx context.Context, data ListEntryRequest) (*ListEntryResponseWrapper, error)
	AdaptUpdateEntry(ctx context.Context, data UpdateEntryRequest) (*UpdateEntryResponseWrapper, error)
	AdaptDeleteEntry(ctx context.Context, data DeleteEntryRequest) (*DeleteEntryResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.EntryUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		entryUseCase: options.EntryUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	entryUseCase entry.EntryUseCase
}

type DefaultHttpInputAdapterOptions struct {
	EntryUseCase entry.EntryUseCase
}

var _ EntryInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptCreateEntry(ctx context.Context, data CreateEntryRequest) (*CreateEntryResponseWrapper, error) {
	var inputData entry.CreateEntryInput
	inputData.ID = data.ID
	inputData.AccountID = data.AccountID
	inputData.MultipleEntryID = data.MultipleEntryID
	inputData.Currency = data.Currency
	inputData.Concept = data.Concept
	inputData.Date = data.Date
	inputData.DebitBalance = data.DebitBalance
	inputData.CreditBalance = data.CreditBalance

	outputData, err := adapter.entryUseCase.CreateEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateEntryResponse
	response.ID = outputData.ID
	response.AccountID = outputData.AccountID
	response.MultipleEntryID = outputData.MultipleEntryID
	response.Currency = outputData.Currency
	if outputData.Concept == nil {
		response.Concept = ""
	} else {
		response.Concept = *outputData.Concept
	}
	response.Concept = *outputData.Concept
	response.Date = outputData.Date
	response.Quarter = outputData.Quarter
	response.DebitBalance = outputData.DebitBalance
	response.CreditBalance = outputData.CreditBalance
	response.TotalBalance = outputData.TotalBalance
	response.State = outputData.State
	if outputData.StateReason == nil {
		response.StateReason = ""
	} else {
		response.StateReason = *outputData.StateReason
	}
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := CreateEntryResponseWrapper{}
	responseWrapper.CreateEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeEntry(ctx context.Context, data DescribeEntryRequest) (*DescribeEntryResponseWrapper, error) {
	var inputData entry.DescribeEntryInput
	inputData.ID = data.ID

	outputData, err := adapter.entryUseCase.DescribeEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeEntryResponse
	response.ID = outputData.ID
	response.AccountID = outputData.AccountID
	response.MultipleEntryID = outputData.MultipleEntryID
	response.Currency = outputData.Currency
	if outputData.Concept == nil {
		response.Concept = ""
	} else {
		response.Concept = *outputData.Concept
	}
	response.Concept = *outputData.Concept
	response.Date = outputData.Date
	response.Quarter = outputData.Quarter
	response.DebitBalance = outputData.DebitBalance
	response.CreditBalance = outputData.CreditBalance
	response.TotalBalance = outputData.TotalBalance
	response.State = outputData.State
	if outputData.StateReason == nil {
		response.StateReason = ""
	} else {
		response.StateReason = *outputData.StateReason
	}
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DescribeEntryResponseWrapper{}
	responseWrapper.DescribeEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListEntries(ctx context.Context, data ListEntryRequest) (*ListEntryResponseWrapper, error) {
	var inputData entry.ListEntryInput

	if data.Expand != "" {
		expand, err := strconv.ParseBool(data.Expand)
		if err != nil {
			return nil, err
		}
		inputData.Expand = &expand
	}

	outputData, err := adapter.entryUseCase.ListEntries(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListEntryResponse
	response.Items = make([]EntryExpandResponse, 0)

	for _, item := range outputData.Items {
		var responseItem EntryExpandResponse
		responseItem.Entry.ID = item.Entry.ID
		responseItem.Entry.AccountID = item.Entry.AccountID
		responseItem.Entry.MultipleEntryID = item.Entry.MultipleEntryID
		responseItem.Entry.Currency = item.Entry.Currency
		if item.Entry.Concept == nil {
			responseItem.Entry.Concept = ""
		} else {
			responseItem.Entry.Concept = *item.Entry.Concept
		}
		responseItem.Entry.Concept = *item.Entry.Concept
		responseItem.Entry.Date = item.Entry.Date
		responseItem.Entry.Quarter = item.Entry.Quarter
		responseItem.Entry.DebitBalance = item.Entry.DebitBalance
		responseItem.Entry.CreditBalance = item.Entry.CreditBalance
		responseItem.Entry.TotalBalance = item.Entry.TotalBalance
		responseItem.Entry.State = item.Entry.State
		if item.Entry.StateReason == nil {
			responseItem.Entry.StateReason = ""
		} else {
			responseItem.Entry.StateReason = *item.Entry.StateReason
		}
		responseItem.Entry.CreationDate = item.Entry.CreationDate
		responseItem.Entry.LastUpdate = item.Entry.LastUpdate
		if item.Account != nil {
			var account accountInputHttpAdapter.Account
			account.ID = item.Account.ID
			account.Name = item.Account.Name
			account.Level = item.Account.Level
			account.Type = *item.Account.Type
			responseItem.Account = &account
		}
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListEntryResponseWrapper{}
	responseWrapper.ListEntryResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateEntry(ctx context.Context, data UpdateEntryRequest) (*UpdateEntryResponseWrapper, error) {
	var inputData entry.UpdateEntryInput
	inputData.ID = data.ID
	inputData.AccountID = data.AccountID
	inputData.Currency = data.Currency
	inputData.Concept = data.Concept
	inputData.Date = data.Date
	inputData.DebitBalance = data.DebitBalance
	inputData.CreditBalance = data.CreditBalance

	outputData, err := adapter.entryUseCase.UpdateEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response UpdateEntryResponse
	response.ID = outputData.ID
	response.AccountID = outputData.AccountID
	response.MultipleEntryID = outputData.MultipleEntryID
	response.Currency = outputData.Currency
	if outputData.Concept == nil {
		response.Concept = ""
	} else {
		response.Concept = *outputData.Concept
	}
	response.Concept = *outputData.Concept
	response.Date = outputData.Date
	response.Quarter = outputData.Quarter
	response.DebitBalance = outputData.DebitBalance
	response.CreditBalance = outputData.CreditBalance
	response.TotalBalance = outputData.TotalBalance
	response.State = outputData.State
	if outputData.StateReason == nil {
		response.StateReason = ""
	} else {
		response.StateReason = *outputData.StateReason
	}
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := UpdateEntryResponseWrapper{}
	responseWrapper.UpdateEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteEntry(ctx context.Context, data DeleteEntryRequest) (*DeleteEntryResponseWrapper, error) {
	var inputData entry.RemoveEntryInput
	inputData.ID = data.ID

	outputData, err := adapter.entryUseCase.RemoveEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DeleteEntryResponse
	response.ID = outputData.ID
	response.AccountID = outputData.AccountID
	response.MultipleEntryID = outputData.MultipleEntryID
	response.Currency = outputData.Currency
	if outputData.Concept == nil {
		response.Concept = ""
	} else {
		response.Concept = *outputData.Concept
	}
	response.Concept = *outputData.Concept
	response.Date = outputData.Date
	response.Quarter = outputData.Quarter
	response.DebitBalance = outputData.DebitBalance
	response.CreditBalance = outputData.CreditBalance
	response.TotalBalance = outputData.TotalBalance
	response.State = outputData.State
	if outputData.StateReason == nil {
		response.StateReason = ""
	} else {
		response.StateReason = *outputData.StateReason
	}
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteEntryResponseWrapper{}
	responseWrapper.DeleteEntryResponse = response

	return &responseWrapper, nil
}
