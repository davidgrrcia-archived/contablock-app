package contablock

import "time"

const (
	// States
	EntryStateCreated       = "CREATED"
	EntryStateUpdated       = "UPDATED"
	EntryStateRemoved       = "REMOVED"
	EntryStatePendingReview = "PENDING_REVIEW"
	EntryStateApproved      = "APPROVED"
	EntryStateRejected      = "REJECTED"
	EntryStatePublished     = "PUBLISHED"
)

type ClientCommonInput struct {
	ContractAddress string
	UserAddress     string
	UserPrivateKey  string
}

type GrantAccessInput struct {
	ClientCommonInput
	AdminAddress string
}

type CreateEntryInput struct {
	ClientCommonInput
	ID              string
	ContractID      string
	UserID          string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            time.Time
	DebitBalance    *float64
	CreditBalance   *float64
}

type GetEntryDataInput struct {
	ClientCommonInput
	ID string
}

type GetEntryDataOutput struct {
	ID              string
	ContractID      string
	UserID          string
	AccountID       string
	MultipleEntryID string
	Currency        string
	Concept         string
	Date            time.Time
	DebitBalance    *float64
	CreditBalance   *float64
	State           string
}

type ListMultipleEntriesIdsInput struct {
	ClientCommonInput
}

type ListMultipleEntriesIdsOutput struct {
	MultipleEntriesIds []string
}

type ListAccountIdsInput struct {
	ClientCommonInput
}

type ListAccountIdsOutput struct {
	AccountIds []string
}

type PublishEntryInput struct {
	ClientCommonInput
	ID string
}

type RemoveEntryInput struct {
	ClientCommonInput
	ID string
}

type ReviewEntryInput struct {
	ClientCommonInput
	ID string
}

type ApproveEntryInput struct {
	ClientCommonInput
	ID string
}

type RejectEntryInput struct {
	ClientCommonInput
	ID string
}
