package entrydb

import "math/big"

type EntryDb struct {
	ID              string  `storage:"id"`
	AccountID       string  `storage:"account_id"`
	MultipleEntryID string  `storage:"multiple_entry_id"`
	Currency        string  `storage:"currency"`
	Concept         *string `storage:"concept"`
	Date            string  `storage:"date"`
	Quarter         int     `storage:"quarter"`
	DebitBalance    big.Int `storage:"debit_balance"`
	CreditBalance   big.Int `storage:"credit_balance"`
	TotalBalance    big.Int `storage:"total_balance"`
	State           string  `storage:"state"`
	StateReason     *string `storage:"state_reason"`
	CreationDate    string  `storage:"creation_date"`
	LastUpdate      string  `storage:"last_update"`
}
