package entryInputHttpAdapter

import (
	"time"
)

type Entry struct {
	ID              string    `json:"id"`
	AccountID       string    `json:"accountId"`
	MultipleEntryID string    `json:"multipleEntryId"`
	Currency        string    `json:"currency"`
	Concept         *string   `json:"concept"`
	Date            time.Time `json:"date"`
	Quarter         int       `json:"quarter"`
	DebitBalance    big.Int   `json:"debitBalance"`
	CreditBalance   big.Int   `json:"creditBalance"`
	TotalBalance    big.Int   `json:"totalBalance"`
	State           string    `json:"state"`
	StateReason     *string   `json:"stateReason"`
	CreationDate    time.Time `json:"creationDate"`
	LastUpdate      time.Time `json:"lastUpdate"`
}

type CreateEntryRequest struct {
	ID              *string   `json:"id"`
	AccountID       string    `json:"accountId"`
	MultipleEntryID string    `json:"multipleEntryId"`
	Currency        string    `json:"currency"`
	Concept         *string   `json:"concept"`
	Date            time.Time `json:"date"`
	DebitBalance    *big.Int  `json:"debitBalance"`
	CreditBalance   *big.Int  `json:"creditBalance"`
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

type ListEntryRequest struct {
}

type ListEntryResponse struct {
	Items []DescribeEntryResponse `json:"items"`
}
type ListEntryResponseWrapper struct {
	ListEntryResponse ListEntryResponse
}

type UpdateEntryRequest struct {
	ID            string    `json:"id"`
	AccountID     string    `json:"accountId"`
	Currency      string    `json:"currency"`
	Concept       *string   `json:"concept"`
	Date          time.Time `json:"date"`
	DebitBalance  *big.Int  `json:"debitBalance"`
	CreditBalance *big.Int  `json:"creditBalance"`
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
