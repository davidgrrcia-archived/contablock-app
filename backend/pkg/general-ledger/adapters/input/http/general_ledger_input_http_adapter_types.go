package generalLedgerInputHttpAdapter

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
)

type ListGeneralLedgerRequest struct {
	Project    string `json:"project"`
	ContractID string `json:"contractId"`
	UserID     string `json:"userId"`
}

type ListGeneralLedgerResponse struct {
	Items []Record `json:"items"`
}
type ListGeneralLedgerResponseWrapper struct {
	ListGeneralLedgerResponse ListGeneralLedgerResponse
}

type Record struct {
	TotalDebitAmount  float64         `json:"totalDebitAmount"`
	TotalCreditAmount float64         `json:"totalCreditAmount"`
	Entries           []entry.Entry   `json:"entries"`
	Account           account.Account `json:"account"`
}
