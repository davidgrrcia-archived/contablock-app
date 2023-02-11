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
	Accounts       []string
	CurrentYear    *string
	PreviousYear   *string
	CurrentAmount  *string
	PreviousAmount *string
}

type ListBalanceInput struct{}

type ListBalanceOutput struct {
	Items []Balance
}

type BalanceStorage interface {
	All(ctx context.Context) (*BalanceStorageCollectionData, error)
}

type BalanceStorageData struct {
	ID             string
	Type           string
	Kind           *string
	Parent         *string
	Child          *string
	Accounts       []string
	CurrentYear    *string
	PreviousYear   *string
	CurrentAmount  *string
	PreviousAmount *string
}

type BalanceStorageCollectionData struct {
	Items []BalanceStorageData
}
