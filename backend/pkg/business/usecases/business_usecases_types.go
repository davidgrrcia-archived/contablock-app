package business

import (
	"context"
)

type Business struct {
	ID                 string
	Name               string
	SocialDenomination string
	RegisteredOffice   string
	Cif                string
	Logo               string
	EthAddress         string
}

type DescribeBusinessInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeBusinessOutput struct {
	Business Business
}

type ListBusinessInput struct{}

type ListBusinessOutput struct {
	Items []Business
}

type BusinessStorage interface {
	Get(ctx context.Context, id string) (*BusinessStorageData, error)
	All(ctx context.Context) (*BusinessStorageCollectionData, error)
}

type BusinessStorageData struct {
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
