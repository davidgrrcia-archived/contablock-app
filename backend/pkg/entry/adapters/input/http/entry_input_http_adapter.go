package entryInputHttpAdapter

import (
	"context"
	"errors"
	"time"

	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type EntryInputHttpAdapter interface {
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

var _ EntryInputHttpAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptCreateEntry(ctx context.Context, data CreateEntryRequest) (*CreateEntryResponseWrapper, error) {
	var inputData entry.CreateEntryInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.ContractID = data.ContractID
	inputData.UserID = data.UserID
	inputData.AccountID = data.AccountID
	inputData.MultipleEntryID = data.MultipleEntryID
	inputData.Currency = data.Currency
	inputData.Concept = data.Concept
	inputData.Date = time.Unix(data.Date, 0)
	inputData.DebitBalance = data.DebitBalance
	inputData.CreditBalance = data.CreditBalance

	outputData, err := adapter.entryUseCase.CreateEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateEntryResponse
	response.Project = outputData.Entry.Project
	response.ID = outputData.Entry.ID
	response.ContractID = outputData.Entry.ContractID
	response.UserID = outputData.Entry.UserID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.BalanceByYearID = outputData.Entry.BalanceByYearID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date.Unix()
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate.Unix()
	response.LastUpdate = outputData.Entry.LastUpdate.Unix()

	responseWrapper := CreateEntryResponseWrapper{}
	responseWrapper.CreateEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDescribeEntry(ctx context.Context, data DescribeEntryRequest) (*DescribeEntryResponseWrapper, error) {
	var inputData entry.DescribeEntryInput
	inputData.Project = data.Project
	inputData.ID = data.ID

	outputData, err := adapter.entryUseCase.DescribeEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DescribeEntryResponse
	response.Project = outputData.Entry.Project
	response.ID = outputData.Entry.ID
	response.ContractID = outputData.Entry.ContractID
	response.UserID = outputData.Entry.UserID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.BalanceByYearID = outputData.Entry.BalanceByYearID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date.Unix()
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate.Unix()
	response.LastUpdate = outputData.Entry.LastUpdate.Unix()

	responseWrapper := DescribeEntryResponseWrapper{}
	responseWrapper.DescribeEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptListEntries(ctx context.Context, data ListEntryRequest) (*ListEntryResponseWrapper, error) {
	var inputData entry.ListEntryInput
	inputData.Project = data.Project

	outputData, err := adapter.entryUseCase.ListEntries(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response ListEntryResponse
	response.Items = make([]DescribeEntryResponse, 0)

	for _, item := range outputData.Items {
		var responseItem DescribeEntryResponse
		responseItem.Project = item.Project
		responseItem.ID = item.ID
		responseItem.ContractID = item.ContractID
		responseItem.UserID = item.UserID
		responseItem.AccountID = item.AccountID
		responseItem.MultipleEntryID = item.MultipleEntryID
		responseItem.BalanceByYearID = item.BalanceByYearID
		responseItem.Currency = item.Currency
		responseItem.Concept = item.Concept
		responseItem.Date = item.Date.Unix()
		responseItem.Quarter = item.Quarter
		responseItem.DebitBalance = item.DebitBalance
		responseItem.CreditBalance = item.CreditBalance
		responseItem.TotalBalance = item.TotalBalance
		responseItem.State = item.State
		responseItem.StateReason = item.StateReason
		responseItem.CreationDate = item.CreationDate.Unix()
		responseItem.LastUpdate = item.LastUpdate.Unix()
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListEntryResponseWrapper{}
	responseWrapper.ListEntryResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptDeleteEntry(ctx context.Context, data DeleteEntryRequest) (*DeleteEntryResponseWrapper, error) {
	var inputData entry.RemoveEntryInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.ContractID = data.ContractID
	inputData.UserID = data.UserID

	outputData, err := adapter.entryUseCase.RemoveEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response DeleteEntryResponse
	response.Project = outputData.Entry.Project
	response.ID = outputData.Entry.ID
	response.ContractID = outputData.Entry.ContractID
	response.UserID = outputData.Entry.UserID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.BalanceByYearID = outputData.Entry.BalanceByYearID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date.Unix()
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate.Unix()
	response.LastUpdate = outputData.Entry.LastUpdate.Unix()

	responseWrapper := DeleteEntryResponseWrapper{}
	responseWrapper.DeleteEntryResponse = response

	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptPublishEntry(ctx context.Context, data PublishEntryRequest) (*PublishEntryResponseWrapper, error) {
	var inputData entry.PublishEntryInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.ContractID = data.ContractID
	inputData.UserID = data.UserID

	outputData, err := adapter.entryUseCase.PublishEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response PublishEntryResponse
	response.Project = outputData.Entry.Project
	response.ID = outputData.Entry.ID
	response.ContractID = outputData.Entry.ContractID
	response.UserID = outputData.Entry.UserID
	response.AccountID = outputData.Entry.AccountID
	response.MultipleEntryID = outputData.Entry.MultipleEntryID
	response.BalanceByYearID = outputData.Entry.BalanceByYearID
	response.Currency = outputData.Entry.Currency
	response.Concept = outputData.Entry.Concept
	response.Date = outputData.Entry.Date.Unix()
	response.Quarter = outputData.Entry.Quarter
	response.DebitBalance = outputData.Entry.DebitBalance
	response.CreditBalance = outputData.Entry.CreditBalance
	response.TotalBalance = outputData.Entry.TotalBalance
	response.State = outputData.Entry.State
	response.StateReason = outputData.Entry.StateReason
	response.CreationDate = outputData.Entry.CreationDate.Unix()
	response.LastUpdate = outputData.Entry.LastUpdate.Unix()

	responseWrapper := PublishEntryResponseWrapper{}
	responseWrapper.PublishEntryResponse = response

	return &responseWrapper, nil
}
