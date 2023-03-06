package eventSubscription

import (
	"context"
	"time"
)

type EventSubscription struct {
	Project      string
	ID           string
	ContractID   string
	Network      string
	CreationDate time.Time
	LastUpdate   time.Time
}

type DescribeEventSubscriptionInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeEventSubscriptionOutput struct {
	EventSubscription EventSubscription
}

type ListEventSubscriptionInput struct {
	Project string `valid:"required,length(1|64)"`
}

type ListEventSubscriptionOutput struct {
	Items []EventSubscription
}

type CreateEventSubscriptionInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	Network    string `valid:"required,length(1|64)"`
}

type CreateEventSubscriptionOutput struct {
	EventSubscription EventSubscription
}

type Filters struct {
}

type EventSubscriptionStorage interface {
	Get(ctx context.Context, project string, id string) (*EventSubscriptionStorageData, error)
	All(ctx context.Context, project string, filters Filters) (*EventSubscriptionStorageCollectionData, error)
	Add(ctx context.Context, input EventSubscriptionStorageData) (*EventSubscriptionStorageData, error)
}

type EventSubscriptionStorageData struct {
	Project      string
	ID           string
	ContractID   string
	Network      string
	CreationDate time.Time
	LastUpdate   time.Time
}

type EventSubscriptionStorageCollectionData struct {
	Items []EventSubscriptionStorageData
}
