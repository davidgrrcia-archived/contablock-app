package account

import (
	"context"
)

type Account struct {
	ID          string
	Name        string
	Level       string
	Type        *string
	DisplayName string
}

type DescribeAccountInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeAccountOutput struct {
	Account Account
}

type ListAccountInput struct{}

type ListAccountOutput struct {
	Items []Account
}

type AccountStorage interface {
	Get(ctx context.Context, id string) (*AccountStorageData, error)
	All(ctx context.Context) (*AccountStorageCollectionData, error)
}

type AccountStorageData struct {
	ID    string
	Name  string
	Level string
	Type  *string
}

type AccountStorageCollectionData struct {
	Items []AccountStorageData
}
