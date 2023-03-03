package contract

import (
	"context"
	"time"
)

type Contract struct {
	Project      string
	ID           string
	Name         string
	Description  string
	EthAddress   string
	Network      string
	AdminID      string
	CreationDate time.Time
	LastUpdate   time.Time
}

type DescribeContractInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeContractOutput struct {
	Contract Contract
}

type ListContractInput struct {
	Project string `valid:"required,length(1|64)"`
}

type ListContractOutput struct {
	Items []Contract
}

type CreateContractInput struct {
	Project     string `valid:"required,length(1|64)"`
	ID          string `valid:"optional,length(1|64)"`
	Name        string `valid:"required,length(1|256)"`
	Description string `valid:"required"`
	EthAddress  string `valid:"required,length(1|42)"`
	Network     string `valid:"required,length(1|64)"`
	AdminID     string `valid:"required,length(1|64)"`
}

type CreateContractOutput struct {
	Contract Contract
}

type Filters struct {
}

type ContractStorage interface {
	Get(ctx context.Context, project string, id string) (*ContractStorageData, error)
	All(ctx context.Context, project string, filters Filters) (*ContractStorageCollectionData, error)
	Add(ctx context.Context, input ContractStorageData) (*ContractStorageData, error)
}

type ContractStorageData struct {
	Project      string
	ID           string
	Name         string
	Description  string
	EthAddress   string
	Network      string
	AdminID      string
	CreationDate time.Time
	LastUpdate   time.Time
}

type ContractStorageCollectionData struct {
	Items []ContractStorageData
}
