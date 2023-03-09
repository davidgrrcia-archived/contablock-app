package entry

import (
	"context"
	"time"
)

const (
	// States
	EntryStateCreated       = "CREATED"
	EntryStateUpdated       = "UPDATED"
	EntryStatePendingReview = "PENDING_REVIEW"
	EntryStateApproved      = "APPROVED"
	EntryStateRejected      = "REJECTED"
	EntryStatePublished     = "PUBLISHED"

	// Roles
	AccountantRole = "ACCOUNTANT"
	AuditorRole    = "AUDITOR"
	ViewerRole     = "VIEWER"
)

type Entry struct {
	Project         string
	ID              string
	ContractID      string
	UserID          string
	AccountID       string
	MultipleEntryID string
	BalanceByYearID *string
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
	Project         string    `valid:"required,length(1|64)"`
	ID              *string   `valid:"optional,length(1|64)"`
	ContractID      string    `valid:"required,length(1|64)"`
	UserID          string    `valid:"required,length(1|64)"`
	AccountID       string    `valid:"required,length(1|64)"`
	MultipleEntryID string    `valid:"required,length(1|64)"`
	Currency        string    `valid:"required,length(3|3)"`
	Concept         *string   `valid:"optional"`
	Date            time.Time `valid:"required"`
	DebitBalance    *float64  `valid:"optional"`
	CreditBalance   *float64  `valid:"optional"`
}

type CreateEntryOutput struct {
	Entry Entry
}

type DescribeEntryInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type DescribeEntryOutput struct {
	Entry Entry
}

type ListEntryInput struct {
	Project         string
	MultipleEntryID *string
	AccountID       *string
}

type ListEntryOutput struct {
	Items []Entry
}

type UpdateEntryInput struct {
	Project string  `valid:"required,length(1|64)"`
	ID      string  `valid:"required,length(1|64)"`
	State   *string `valid:"optional"`
}

type UpdateEntryOutput struct {
	Entry Entry
}

type RemoveEntryInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type RemoveEntryOutput struct {
	Entry Entry
}

type PublishEntryInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type PublishEntryOutput struct {
	Entry Entry
}

type ReviewEntryInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type ReviewEntryOutput struct {
	Entry Entry
}

type ApproveEntryInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type ApproveEntryOutput struct {
	Entry Entry
}

type RejectEntryInput struct {
	Project    string `valid:"required,length(1|64)"`
	ID         string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type RejectEntryOutput struct {
	Entry Entry
}

type Filters struct {
	MultipleEntryID *string
	AccountID       *string
}

type EntryStorage interface {
	Add(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Get(ctx context.Context, project string, id string) (*EntryStorageData, error)
	All(ctx context.Context, project string, filters Filters) (*EntryStorageCollectionData, error)
	Edit(ctx context.Context, input EntryStorageData) (*EntryStorageData, error)
	Remove(ctx context.Context, project string, id string) (*EntryStorageData, error)
}

type EntryStorageData struct {
	Project         string
	ID              string
	ContractID      string
	UserID          string
	AccountID       string
	MultipleEntryID string
	BalanceByYearID *string
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

type EntryStorageCollectionData struct {
	Items []EntryStorageData
}

// SC Events
type EntryCreatedEventInput struct {
	Project         string    `valid:"required,length(1|64)"`
	ID              string    `valid:"required,length(1|64)"`
	ContractID      string    `valid:"required,length(1|64)"`
	UserID          string    `valid:"required,length(1|64)"`
	AccountID       string    `valid:"required,length(1|64)"`
	MultipleEntryID string    `valid:"required,length(1|64)"`
	Currency        string    `valid:"required,length(3|3)"`
	Concept         string    `valid:"optional"`
	Date            time.Time `valid:"required"`
	DebitBalance    *float64  `valid:"optional"`
	CreditBalance   *float64  `valid:"optional"`
}

type EntryReviewRequestedEventInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type EntryRemovedEventInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type EntryPublishedEventInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type EntryApprovedEventInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}

type EntryRejectedEventInput struct {
	Project string `valid:"required,length(1|64)"`
	ID      string `valid:"required,length(1|64)"`
}
