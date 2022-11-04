package entryInputHttpAdapter

import (
	"context"
	"errors"
	"github.com/davidevloper/HoldContract-go/pkg/usecases/entry"
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
	inputData.Entryname = data.Entryname
	inputData.Password = data.Password
	inputData.Email = data.Email

	outputData, err := adapter.entryUseCase.CreateEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateEntryResponse
	response.ID = outputData.ID
	response.Password = outputData.Password
	response.Email = outputData.Email
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
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

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
		responseItem.Password = item.Password
		responseItem.Email = item.Email
		responseItem.CreationDate = item.CreationDate
		responseItem.LastUpdate = item.LastUpdate
		response.Items = append(response.Items, responseItem)
	}

	responseWrapper := ListEntryResponseWrapper{}
	responseWrapper.ListEntryResponse = response
	return &responseWrapper, nil
}

func (adapter *DefaultHttpInputAdapter) AdaptUpdateEntry(ctx context.Context, data UpdateEntryRequest) (*UpdateEntryResponseWrapper, error) {
	var inputData entry.UpdateEntryInput
	inputData.ID = data.ID
	inputData.Entryname = data.Entryname
	inputData.Password = data.Password
	inputData.Email = data.Email

	outputData, err := adapter.entryUseCase.UpdateEntry(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response UpdateEntryResponse
	response.ID = outputData.ID
	response.Password = outputData.Password
	response.Email = outputData.Email
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
	response.Password = outputData.Password
	response.Email = outputData.Email
	response.CreationDate = outputData.CreationDate
	response.LastUpdate = outputData.LastUpdate

	responseWrapper := DeleteEntryResponseWrapper{}
	responseWrapper.DeleteEntryResponse = response

	return &responseWrapper, nil
}
