package eventSubscriptionInputHttpAdapter

import (
	"context"
	"errors"

	eventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/usecases"
)

type EventSubscriptionInputAdapter interface {
	AdaptCreateEventSubscription(ctx context.Context, data CreateEventSubscriptionRequest) (*CreateEventSubscriptionResponseWrapper, error)
}

func NewDefaultHttpInputAdapter(options DefaultHttpInputAdapterOptions) (*DefaultHttpInputAdapter, error) {
	if options.EventSubscriptionUseCase == nil {
		return nil, errors.New("invalid options adapter")
	}

	return &DefaultHttpInputAdapter{
		eventSubscriptionUseCase: options.EventSubscriptionUseCase,
	}, nil
}

type DefaultHttpInputAdapter struct {
	eventSubscriptionUseCase eventSubscription.EventSubscriptionUseCase
}

type DefaultHttpInputAdapterOptions struct {
	EventSubscriptionUseCase eventSubscription.EventSubscriptionUseCase
}

var _ EventSubscriptionInputAdapter = (*DefaultHttpInputAdapter)(nil)

func (adapter *DefaultHttpInputAdapter) AdaptCreateEventSubscription(ctx context.Context, data CreateEventSubscriptionRequest) (*CreateEventSubscriptionResponseWrapper, error) {
	var inputData eventSubscription.CreateEventSubscriptionInput
	inputData.Project = data.Project
	inputData.ID = data.ID
	inputData.ContractID = data.ContractID
	inputData.Network = data.Network

	outputData, err := adapter.eventSubscriptionUseCase.CreateEventSubscription(ctx, inputData)
	if err != nil {
		return nil, err
	}

	var response CreateEventSubscriptionResponse
	response.Project = outputData.EventSubscription.Project
	response.ID = outputData.EventSubscription.ID
	response.ContractID = outputData.EventSubscription.ContractID
	response.Network = outputData.EventSubscription.Network
	response.CreationDate = outputData.EventSubscription.CreationDate.Unix()
	response.LastUpdate = outputData.EventSubscription.LastUpdate.Unix()

	responseWrapper := CreateEventSubscriptionResponseWrapper{}
	responseWrapper.CreateEventSubscriptionResponse = response

	return &responseWrapper, nil
}
