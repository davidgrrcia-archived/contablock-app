package entry

import (
	"context"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	"time"
)

const (
	// States
	EntryStateCreated = "CREATED"
)

type Entry struct {
	ID              string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         *string
	Date            time.Time
	Quarter         int
	DebitBalance    float64
	CreditBalance   float64
	TotalBalance    float64
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
	DebitBalance    *float64  `valid:"optional"`
	CreditBalance   *float64  `valid:"optional"`
}

type CreateEntryOutput struct {
	Entry
}

type DescribeEntryInput struct {
	ID string `valid:"required,length(1|64)"`
}

type DescribeEntryOutput struct {
	Entry
}

type ListEntryInput struct {
	Expand *bool `valid:"optional"`
}

type ListEntryOutput struct {
	Items []EntryExpand
}

type EntryExpand struct {
	Entry   Entry
	Account *account.Account
}

type UpdateEntryInput struct {
	ID            string     `valid:"required,length(1|64)"`
	AccountID     *string    `valid:"optional,length(1|64)"`
	Currency      *string    `valid:"optional,length(3|3)"`
	Concept       *string    `valid:"optional"`
	Date          *time.Time `valid:"optional"`
	DebitBalance  *float64   `valid:"optional"`
	CreditBalance *float64   `valid:"optional"`
	State         *string    `valid:"optional"`
	StateReason   *string    `valid:"optional"`
}

type UpdateEntryOutput struct {
	Entry
}

type RemoveEntryInput struct {
	ID string `valid:"required,length(1|64)"`
}

type RemoveEntryOutput struct {
	Entry
}

type EntryStorage interface {
	Add(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Get(ctx context.Context, id string) (*EntryStorageData, error)
	All(ctx context.Context) (*EntryStorageCollectionData, error)
	Edit(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Remove(ctx context.Context, id string) (*EntryStorageData, error)
}

type EntryStorageData struct {
	Entry
}

type EntryStorageCollectionData struct {
	Items []Entry
}
