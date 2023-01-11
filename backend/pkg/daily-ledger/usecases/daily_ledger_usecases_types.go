package dailyledger

import (
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

type ListDailyLedgerInput struct{}

type ListDailyLedgerOutput struct {
	Items []AccountingRecord
}

type AccountingRecord struct {
	Entries []Record
}

type Record struct {
	Entry   []entry.Entry
	Account []account.Account
}
