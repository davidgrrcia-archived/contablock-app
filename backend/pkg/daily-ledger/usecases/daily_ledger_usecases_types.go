package dailyledger

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type ListDailyLedgerInput struct {
	Project    string
	ContractID string
	UserID     string
}

type ListDailyLedgerOutput struct {
	Items []Record
}

type Record struct {
	Project           string
	ID                string
	TotalDebitAmount  float64
	TotalCreditAmount float64
	EntriesExpanded   []EntryExpanded
}

type EntryExpanded struct {
	Entry   entry.Entry
	Account account.Account
}
