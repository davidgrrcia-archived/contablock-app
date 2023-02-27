package generalledger

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type ListGeneralLedgerInput struct{
	Project string
}

type ListGeneralLedgerOutput struct {
	Items []Record
}

type Record struct {
	Entries []entry.Entry
	Account account.Account
}
