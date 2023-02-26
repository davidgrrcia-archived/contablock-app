package business

import (
	"context"
)

type Business struct {
	Project            string
	ID                 string
	Name               string
	SocialDenomination string
	RegisteredOffice   string
	Cif                string
	Logo               string
	EthAddress         string
}

type DescribeBusinessInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeBusinessOutput struct {
	Business Business
}

type ListBusinessInput struct {
	Project string `valid:"required,length(1|64)"`
}

type ListBusinessOutput struct {
	Items []Business
}

type CreateBusinessInput struct {
	Project            string `valid:"required,length(1|64)"`
	ID                 string `valid:"required,length(1|64)"`
	Name               string `valid:"required,length(1|256)"`
	SocialDenomination string `valid:"required,length(1|256)"`
	RegisteredOffice   string `valid:"required,length(1|256)"`
	Cif                string `valid:"required,length(1|9)"`
	Logo               string `valid:"required"`
	EthAddress         string `valid:"required,length(1|42)"`
}

type CreateBusinessOutput struct {
	Business Business
}

type BusinessStorage interface {
	Get(ctx context.Context, project string, id string) (*BusinessStorageData, error)
	All(ctx context.Context, project string) (*BusinessStorageCollectionData, error)
	Add(ctx context.Context, input BusinessStorageData) (*BusinessStorageData, error)
}

type BusinessStorageData struct {
	Project            string
	ID                 string
	Name               string
	SocialDenomination string
	RegisteredOffice   string
	Cif                string
	Logo               string
	EthAddress         string
}

type BusinessStorageCollectionData struct {
	Items []BusinessStorageData
}
