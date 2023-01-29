package contablock

import (
	"math/big"
)

const (
	// States
	EntryStateCreated = "CREATED"
)

type CreateEntryInput struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
}

type GetEntryDataInput struct {
	ID string
}

type GetEntryDataOutput struct {
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

type ListMultipleEntriesIdsOutput struct {
	multipleEntriesIds []string
}

type ListAccountIdsOutput struct {
	accountIds []string
}

type PublishEntryInput struct {
	ID string
}

type RemoveEntryInput struct {
	ID string
}
