package contablock

import (
	"math/big"
)

type ContaBlockContractAdapter interface {
	CreateEntry(input CreateEntryAdapterInput) error
	GetEntryData(input GetEntryDataAdapterInput) (*GetEntryDataAdapterOutput, error)
	ListMultipleEntriesIds() (*ListMultipleEntriesIdsAdapterOutput, error)
	ListAccountIds() (*ListAccountIdsAdapterOutput, error)
	PublishEntry(input PublishEntryAdapterInput) error
}

type CreateEntryAdapterInput struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
}

type GetEntryDataAdapterInput struct {
	ID string
}

type GetEntryDataAdapterOutput struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	State           string
}

type ListMultipleEntriesIdsAdapterInput struct {
	ID string
}

type ListMultipleEntriesIdsAdapterOutput struct {
	MultipleEntriesIds []string
}

type ListAccountIdsAdapterInput struct {
	ID string
}

type ListAccountIdsAdapterOutput struct {
	AccountIds []string
}

type PublishEntryAdapterInput struct {
	ID string
}
