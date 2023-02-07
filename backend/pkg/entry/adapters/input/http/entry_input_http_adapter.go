package entryInputHttpAdapter

import (
	"context"
	"errors"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type EntryInputAdapter interface {
	AdaptCreateEntry(ctx context.Context, data CreateEntryRequest) (*CreateEntryResponseWrapper, error)
	AdaptDescribeEntry(ctx context.Context, data DescribeEntryRequest) (*DescribeEntryResponseWrapper, error)
	AdaptListEntries(ctx context.Context, data ListEntryRequest) (*ListEntryResponseWrapper, error)
	AdaptDeleteEntry(ctx context.Context, data DeleteEntryRequest) (*DeleteEntryResponseWrapper, error)
	AdaptPublishEntry(ctx context.Context, data PublishEntryRequest) (*PublishEntryResponseWrapper, error)
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
	response.ID = outputData.Entry.ID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate
	response.LastUpdate = outputData.Entry.LastUpdate

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
	response.ID = outputData.Entry.ID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate
	response.LastUpdate = outputData.Entry.LastUpdate

	responseWrapper := DescribeEntryResponseWrapper{}
	responseWrapper.DescribeEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListEntries(ctx context.Context, data ListEntryRequest) (*ListEntryResponseWrapper, error) {
	var inputData entry.ListEntryInput

	outputData, err := adapter.entryUseCase.ListEntries(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListEntryResponse
	response.Items = make([]DescribeEntryResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeEntryResponse
		responseItem.ID = item.ID
		responseItem.AccountID = item.AccountID
		responseItem.MultipleEntryID = item.MultipleEntryID
		responseItem.Currency = item.Currency
		responseItem.Concept = item.Concept //todo
		responseItem.Date = item.Date
		responseItem.Quarter = item.Quarter
		responseItem.DebitBalance = item.DebitBalance
		responseItem.CreditBalance = item.CreditBalance
		responseItem.TotalBalance = item.TotalBalance
		responseItem.State = item.State
		responseItem.StateReason = item.StateReason
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListEntryResponseWrapper{}
	responseWrapper.ListEntryResponse = response
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
	response.ID = outputData.Entry.ID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate
	response.LastUpdate = outputData.Entry.LastUpdate

	responseWrapper := DeleteEntryResponseWrapper{}
	responseWrapper.DeleteEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptPublishEntry(ctx context.Context, data PublishEntryRequest) (*PublishEntryResponseWrapper, error) {
	var inputData entry.PublishEntryInput
	inputData.ID = data.ID

	outputData, err := adapter.entryUseCase.PublishEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response PublishEntryResponse
	response.ID = outputData.Entry.ID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate
	response.LastUpdate = outputData.Entry.LastUpdate

	responseWrapper := PublishEntryResponseWrapper{}
	responseWrapper.PublishEntryResponse = response

	return &responseWrapper, nil
}
