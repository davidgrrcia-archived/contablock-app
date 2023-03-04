package dailyLedgerInputHttpAdapter

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
)

type ListDailyLedgerRequest struct {
	Project    string `json:"project"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}

type ListDailyLedgerResponse struct {
	Items []Record `json:"items"`
}
type ListDailyLedgerResponseWrapper struct {
	ListDailyLedgerResponse ListDailyLedgerResponse
}

type Record struct {
	Project           string          `json:"project"`
	ID                string          `json:"id"`
	TotalDebitAmount  float64         `json:"totalDebitAmount"`
	TotalCreditAmount float64         `json:"totalCreditAmount"`
	EntriesExpanded   []EntryExpanded `json:"entriesExpanded"`
}

type EntryExpanded struct {
	Entry   entry.Entry     `json:"entry"`
	Account account.Account `json:"account"`
}
