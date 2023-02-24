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
	Project     string
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
}

type Subgroup struct {
	Project     string
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
	Accounts    []Account
}

type Group struct {
	Project     string
	ID          string
	Name        string
	Level       string
	BalanceID   *string
	DisplayName string
	Subgroups   []Subgroup
}

type DescribeAccountInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeAccountOutput struct {
	Account Account
}

type ListAccountInput struct {
	Project string `valid:"required,length(1|64)"`
}

type ListAccountOutput struct {
	Items []Group
}

type CreateAccountInput struct {
	Project   string  `valid:"required,length(1|64)"`
	ID        string  `valid:"required,length(1|64)"`
	Name      string  `valid:"required,length(1|256)"`
	Level     string  `valid:"required,length(1|64)"`
	BalanceID *string `valid:"optional,length(1|64)"`
}

type CreateAccountOutput struct {
	Account Account
}

type AccountStorage interface {
	Get(ctx context.Context, project string, id string) (*AccountStorageData, error)
	All(ctx context.Context, project string) (*AccountStorageCollectionData, error)
	Edit(ctx context.Context, input AccountStorageData) (*AccountStorageData, error)
	Add(ctx context.Context, input AccountStorageData) (*AccountStorageData, error)
}

type AccountStorageData struct {
	Project   string
	ID        string
	Name      string
	Level     string
	BalanceID *string
}

type AccountStorageCollectionData struct {
	Items []AccountStorageData
}
