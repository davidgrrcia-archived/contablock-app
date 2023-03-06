package eventSubscription

import (
	"context"
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
)

type EventSubscriptionUseCase interface {
	DescribeEventSubscription(ctx context.Context, input DescribeEventSubscriptionInput) (*DescribeEventSubscriptionOutput, error)
	ListEventSubscriptions(ctx context.Context, input ListEventSubscriptionInput) (*ListEventSubscriptionOutput, error)
	CreateEventSubscription(ctx context.Context, input CreateEventSubscriptionInput) (*CreateEventSubscriptionOutput, error)
}

func NewDefaultEventSubscriptionUseCase(options DefaultEventSubscriptionUseCaseOptions) (*DefaultEventSubscriptionUseCase, error) {
	if options.EventSubscriptionStorage == nil {
		return nil, errors.New("cannot create an instance of EventSubscriptionUseCase because provided EventSubscriptionStorage is nil")
	}

	return &DefaultEventSubscriptionUseCase{
		eventSubscriptionStorage: options.EventSubscriptionStorage,
	}, nil
}

type DefaultEventSubscriptionUseCase struct {
	eventSubscriptionStorage EventSubscriptionStorage
}

type DefaultEventSubscriptionUseCaseOptions struct {
	EventSubscriptionStorage EventSubscriptionStorage
}

var _ EventSubscriptionUseCase = (*DefaultEventSubscriptionUseCase)(nil)

func (useCase *DefaultEventSubscriptionUseCase) DescribeEventSubscription(ctx context.Context, input DescribeEventSubscriptionInput) (*DescribeEventSubscriptionOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedEventSubscription, err := useCase.eventSubscriptionStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedEventSubscription == nil {
		return nil, errors.New("eventSubscription with project [" + input.Project + "] and id [" + input.ID + "] not found")
	}

	var output DescribeEventSubscriptionOutput
	output.EventSubscription.Project = storedEventSubscription.Project
	output.EventSubscription.ID = storedEventSubscription.ID
	output.EventSubscription.ContractID = storedEventSubscription.ContractID
	output.EventSubscription.Network = storedEventSubscription.Network
	output.EventSubscription.CreationDate = storedEventSubscription.CreationDate
	output.EventSubscription.LastUpdate = storedEventSubscription.LastUpdate

	return &output, err
}

func (useCase *DefaultEventSubscriptionUseCase) ListEventSubscriptions(ctx context.Context, input ListEventSubscriptionInput) (*ListEventSubscriptionOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	var filters Filters
	storedEventSubscriptions, err := useCase.eventSubscriptionStorage.All(ctx, input.Project, filters)
	if err != nil {
		return nil, err
	}

	var eventSubscriptions []EventSubscription

	for _, item := range storedEventSubscriptions.Items {
		var eventSubscription EventSubscription
		eventSubscription.Project = item.Project
		eventSubscription.ID = item.ID
		eventSubscription.ContractID = item.ContractID
		eventSubscription.Network = item.Network
		eventSubscription.CreationDate = item.CreationDate
		eventSubscription.LastUpdate = item.LastUpdate
		eventSubscriptions = append(eventSubscriptions, eventSubscription)
	}

	return &ListEventSubscriptionOutput{
		Items: eventSubscriptions,
	}, err
}

func (useCase *DefaultEventSubscriptionUseCase) CreateEventSubscription(ctx context.Context, input CreateEventSubscriptionInput) (*CreateEventSubscriptionOutput, error) {
	_, err := govalidator.ValidateStruct(input)
	if err != nil {
		return nil, errors.New("invalid input: " + err.Error())
	}

	storedEventSubscription, err := useCase.eventSubscriptionStorage.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	if storedEventSubscription != nil {
		return nil, errors.New("eventSubscription [" + storedEventSubscription.ID + "] already exists")
	}

	var desiredEventSubscription EventSubscriptionStorageData
	desiredEventSubscription.Project = input.Project
	desiredEventSubscription.ID = input.ID
	desiredEventSubscription.ContractID = input.ContractID
	desiredEventSubscription.Network = input.Network
	desiredEventSubscription.CreationDate = time.Now()
	desiredEventSubscription.LastUpdate = desiredEventSubscription.CreationDate

	addedEventSubscription, err := useCase.eventSubscriptionStorage.Add(ctx, desiredEventSubscription)
	if err != nil {
		return nil, err
	}

	var output CreateEventSubscriptionOutput
	output.EventSubscription.Project = addedEventSubscription.Project
	output.EventSubscription.ID = addedEventSubscription.ID
	output.EventSubscription.ContractID = addedEventSubscription.ContractID
	output.EventSubscription.Network = addedEventSubscription.Network
	output.EventSubscription.CreationDate = addedEventSubscription.CreationDate
	output.EventSubscription.LastUpdate = addedEventSubscription.LastUpdate

	return &output, nil
}
