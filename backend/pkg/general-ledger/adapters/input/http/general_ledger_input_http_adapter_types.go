package generalLedgerInputHttpAdapter

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
)

type ListGeneralLedgerRequest struct{}

type ListGeneralLedgerResponse struct {
	Items []Record `json:"items"`
}
type ListGeneralLedgerResponseWrapper struct {
	ListGeneralLedgerResponse ListGeneralLedgerResponse
}

type Record struct {
	Entries []entry.Entry   `json:"entries"`
	Account account.Account `json:"account"`
}
