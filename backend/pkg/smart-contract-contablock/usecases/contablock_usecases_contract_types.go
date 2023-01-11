package contablock

import (
	"math/big"
)

type ContaBlockContractAdapter interface {
	CreateEntry(input CreateEntryAdapterInput) error
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
