package generalledger

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type ListGeneralLedgerInput struct {
	Project    string `valid:"required,length(1|64)"`
	ContractID string `valid:"required,length(1|64)"`
	UserID     string `valid:"required,length(1|64)"`
}

type ListGeneralLedgerOutput struct {
	Items []Record
}

type Record struct {
	TotalDebitAmount  float64
	TotalCreditAmount float64
	Entries           []entry.Entry
	Account           account.Account
}
