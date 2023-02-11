package balance

import (
	"context"
)

type Balance struct {
	ID             string
	Type           string
	Kind           *string
	Parent         *string
	Child          *string
	CurrentYear    int
	PreviousYear   int
	CurrentAmount  float64
	PreviousAmount float64
}

type CreateBalanceInput struct {
	ID        *string
	BalanceID string
	Amount    float64
	Year      int
}

type CreateBalanceOutput struct {
	ID        string
	BalanceID string
	Amount    float64
	Year      int
}

type ListBalanceInput struct {
	Year *int
}

type ListBalanceOutput struct {
	Items []Balance
}

type Filters struct {
	Year *int
}

type BalanceStorage interface {
	Add(ctx context.Context, input BalanceByYearStorageData) (*BalanceByYearStorageData, error)
	All(ctx context.Context, filters Filters) (*BalanceStorageCollectionData, error)
}

type BalanceByYearStorageData struct {
	ID        string
	BalanceID string
	Amount    float64
	Year      int
}

type BalanceStorageData struct {
	ID             string
	Type           string
	Kind           *string
	Parent         *string
	Child          *string
	CurrentYear    int
	PreviousYear   int
	CurrentAmount  float64
	PreviousAmount float64
}

type BalanceStorageCollectionData struct {
	Items []BalanceStorageData
}
