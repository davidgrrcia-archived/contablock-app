package dailyLedgerInputHttpAdapter

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
)

type ListDailyLedgerRequest struct{}

type ListDailyLedgerResponse struct {
	Items []Record `json:"items"`
}
type ListDailyLedgerResponseWrapper struct {
	ListDailyLedgerResponse ListDailyLedgerResponse
}

type Record struct {
	ID              string
	EntriesExpanded []EntryExpanded
}

type EntryExpanded struct {
	Entry   entry.Entry
	Account account.Account
}
