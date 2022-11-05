package entryInputHttpAdapter

import (
	"time"
)

type Entry struct {
	ID              string    `json:"id"`
	AccountID       string    `json:"account_id"`
	MultipleEntryID string    `json:"multiple_entry_id"`
	Currency        string    `json:"currency"`
	Concept         string    `json:"concept"`
	Date            time.Time `json:"date"`
	Quarter         int       `json:"quarter"`
	DebitBalance    float64   `json:"debit_balance"`
	CreditBalance   float64   `json:"credit_balance"`
	TotalBalance    float64   `json:"total_balance"`
	State           string    `json:"state"`
	StateReason     string    `json:"state_reason"`
	CreationDate    time.Time `json:"creation-date"`
	LastUpdate      time.Time `json:"last-update"`
}

type CreateEntryRequest struct {
	ID              *string   `json:"id"`
	AccountID       string    `json:"account_id"`
	MultipleEntryID string    `json:"multiple_entry_id"`
	Currency        string    `json:"currency"`
	Concept         *string   `json:"concept"`
	Date            time.Time `json:"date"`
	DebitBalance    *float64  `json:"debit_balance"`
	CreditBalance   *float64  `json:"credit_balance"`
}

type CreateEntryResponse struct {
	Entry
}
type CreateEntryResponseWrapper struct {
	CreateEntryResponse CreateEntryResponse
}

type DescribeEntryRequest struct {
	ID string `json:"id"`
}
type DescribeEntryResponse struct {
	Entry
}
type DescribeEntryResponseWrapper struct {
	DescribeEntryResponse DescribeEntryResponse
}

type ListEntryRequest struct{}

type ListEntryResponse struct {
	Items []DescribeEntryResponse `json:"items"`
}
type ListEntryResponseWrapper struct {
	ListEntryResponse ListEntryResponse
}

type UpdateEntryRequest struct {
	ID            string     `json:"id"`
	AccountID     *string    `json:"account_id"`
	Currency      *string    `json:"currency"`
	Concept       *string    `json:"concept"`
	Date          *time.Time `json:"date"`
	DebitBalance  *float64   `json:"debit_balance"`
	CreditBalance *float64   `json:"credit_balance"`
}
type UpdateEntryResponse struct {
	Entry
}
type UpdateEntryResponseWrapper struct {
	UpdateEntryResponse UpdateEntryResponse
}

type DeleteEntryRequest struct {
	ID string `json:"id"`
}
type DeleteEntryResponse struct {
	Entry
}
type DeleteEntryResponseWrapper struct {
	DeleteEntryResponse DeleteEntryResponse
}
