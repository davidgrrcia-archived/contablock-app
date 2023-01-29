package entry

import (
	"context"
	"math/big"
	"time"
)

const (
	// States
	EntryStateCreated   = "CREATED"
	EntryStatePublished = "PUBLISHED"
)

type Entry struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         *string
	Date            time.Time
	Quarter         int
	DebitBalance    big.Int
	CreditBalance   big.Int
	TotalBalance    big.Int
	State           string
	StateReason     *string
	CreationDate    time.Time
	LastUpdate      time.Time
}

type CreateEntryInput struct {
	ID              *string   `valid:"optional,length(1|64)"`
	AccountID       string    `valid:"required,length(1|64)"`
	MultipleEntryID string    `valid:"required,length(1|64)"`
	Currency        string    `valid:"required,length(3|3)"`
	Concept         *string   `valid:"optional"`
	Date            time.Time `valid:"required"`
	DebitBalance    *big.Int  `valid:"optional"`
	CreditBalance   *big.Int  `valid:"optional"`
}

type CreateEntryOutput struct {
	Entry Entry
}

type DescribeEntryInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeEntryOutput struct {
	Entry Entry
}

type ListEntryInput struct {
}

type ListEntryOutput struct {
	Items []Entry
}

type UpdateEntryInput struct {
	Entry Entry
}

type UpdateEntryOutput struct {
	Entry Entry
}

type RemoveEntryInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveEntryOutput struct {
	Entry Entry
}

type PublishEntryInput struct {
	ID string `valid:"required,length(1|64)"`
}

type PublishEntryOutput struct {
	Entry Entry
}

type EntryStorage interface {
	Add(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Get(ctx context.Context, id string) (*EntryStorageData, error)
	All(ctx context.Context) (*EntryStorageCollectionData, error)
	Edit(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Remove(ctx context.Context, id string) (*EntryStorageData, error)
}

type EntryStorageData struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         *string
	Date            time.Time
	Quarter         int
	DebitBalance    big.Int
	CreditBalance   big.Int
	TotalBalance    big.Int
	State           string
	StateReason     *string
	CreationDate    time.Time
	LastUpdate      time.Time
}

type EntryStorageCollectionData struct {
	Items []EntryStorageData
}
