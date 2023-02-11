package account

import (
	"context"
)

const (
	LevelGroup      = "Grupo"
	LevelSubGroup   = "Subgrupo"
	LevelAccount    = "Cuenta"
	LevelSubAccount = "Subcuenta"
)

type Account struct {
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
}

type Subgroup struct {
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
	Accounts    []Account
}

type Group struct {
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
	Subgroups   []Subgroup
}

type DescribeAccountInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeAccountOutput struct {
	Account Account
}

type ListAccountInput struct{}

type ListAccountOutput struct {
	Items []Group
}

type AccountStorage interface {
	Get(ctx context.Context, id string) (*AccountStorageData, error)
	All(ctx context.Context) (*AccountStorageCollectionData, error)
	Edit(ctx context.Context, input AccountStorageData) (*AccountStorageData, error)
}

type AccountStorageData struct {
	ID        string
	Name      string
	Level     string
	BalanceID *string
}

type AccountStorageCollectionData struct {
	Items []AccountStorageData
}
