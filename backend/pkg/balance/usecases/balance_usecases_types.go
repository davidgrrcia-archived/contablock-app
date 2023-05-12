package balance

import (
	"context"
)

type Balance struct {
	Project  string
	ID       string
	Type     string
	Kind     *string
	Parent   *string
	Child    *string
	Accounts []string
}

type BalanceByYear struct {
	Project        string
	ID             string
	BalanceID      string
	CurrentYear    int
	PreviousYear   int
	CurrentAmount  float64
	PreviousAmount float64
}

type CreateBalanceByYearInput struct {
	Project   string
	ID        *string
	BalanceID string
	Amount    float64
	Year      int
}

type CreateBalanceByYearOutput struct {
	Project   string
	ID        string
	BalanceID string
	Amount    float64
	Year      int
}

type DescribeBalanceInput struct {
	Project string
	ID      string
}

type DescribeBalanceOutput struct {
	Project  string
	ID       string
	Type     string
	Kind     *string
	Parent   *string
	Child    *string
	Accounts []string
}

type ListBalanceInput struct {
	Project string
	Year    *int
}

type ListBalanceOutput struct {
	Items []Balance
}

type RemoveBalanceInput struct {
	Project string
	ID      string
}

type RemoveBalanceOutput struct {
	Project   string
	ID        string
	BalanceID string
	Amount    float64
	Year      int
}

type CreateBalanceInput struct {
	Project  string   `valid:"required,length(1|64)"`
	ID       string   `valid:"required,length(1|64)"`
	Type     string   `valid:"required,length(1|256)"`
	Kind     *string  `valid:"length(1|256)"`
	Parent   *string  `valid:"length(1|256)"`
	Child    *string  `valid:"length(1|256)"`
	Accounts []string `valid:"required"`
}

type CreateBalanceOutput struct {
	Project  string
	ID       string
	Type     string
	Kind     *string
	Parent   *string
	Child    *string
	Accounts []string
}

type Filters struct {
	Year *int
}

type BalanceStorage interface {
	AddByYear(ctx context.Context, input BalanceByYearStorageData) (*BalanceByYearStorageData, error)
	All(ctx context.Context, project string, filters Filters) (*BalanceStorageCollectionData, error)
	AllByYear(ctx context.Context, project string, filters Filters) (*BalanceByYearStorageCollectionData, error)
	Remove(ctx context.Context, project string, id string) (*BalanceByYearStorageData, error)
	Add(ctx context.Context, input AddBalanceStorageData) (*AddBalanceStorageData, error)
	Get(ctx context.Context, project string, id string) (*AddBalanceStorageData, error)
}

type BalanceByYearStorageData struct {
	Project   string
	ID        string
	BalanceID string
	Amount    float64
	Year      int
}

type BalanceStorageData struct {
	Project  string
	ID       string
	Type     string
	Kind     *string
	Parent   *string
	Child    *string
	Accounts []string
}

type AddBalanceStorageData struct {
	Project  string
	ID       string
	Type     string
	Kind     *string
	Parent   *string
	Child    *string
	Accounts []string
}

type BalanceStorageCollectionData struct {
	Items []BalanceStorageData
}

type BalanceByYearStorageCollectionData struct {
	Items []BalanceByYearStorageData
}
