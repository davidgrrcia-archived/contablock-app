package dailyledger

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type ListDailyLedgerInput struct {
	Project    string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
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
