package entryInputHttpAdapter

import (
	"time"

	accountInputHttpAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
)

type Entry struct {
	ID              string    `json:"id"`
	AccountID       string    `json:"accountId"`
	MultipleEntryID string    `json:"multipleEntryId"`
	Currency        string    `json:"currency"`
	Concept         string    `json:"concept"`
	Date            time.Time `json:"date"`
	Quarter         int       `json:"quarter"`
	DebitBalance    float64   `json:"debitBalance"`
	CreditBalance   float64   `json:"creditBalance"`
	TotalBalance    float64   `json:"totalBalance"`
	State           string    `json:"state"`
	StateReason     string    `json:"stateReason"`
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
	DebitBalance    *float64  `json:"debitBalance"`
	CreditBalance   *float64  `json:"creditBalance"`
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
	Expand string `json:"expand"`
}

type EntryExpandResponse struct {
	Entry   Entry                            `json:"entry"`
	Account *accountInputHttpAdapter.Account `json:"account"`
}

type ListEntryResponse struct {
	Items []EntryExpandResponse `json:"items"`
}
type ListEntryResponseWrapper struct {
	ListEntryResponse ListEntryResponse
}

type UpdateEntryRequest struct {
	ID            string     `json:"id"`
	AccountID     *string    `json:"accountId"`
	Currency      *string    `json:"currency"`
	Concept       *string    `json:"concept"`
	Date          *time.Time `json:"date"`
	DebitBalance  *float64   `json:"debitBalance"`
	CreditBalance *float64   `json:"creditBalance"`
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
