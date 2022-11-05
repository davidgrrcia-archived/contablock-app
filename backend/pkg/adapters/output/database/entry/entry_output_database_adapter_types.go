package entrydb

import (
	"time"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
)

type EntryDb struct {
	ID              string  `storage:"id"`
	AccountID       string  `storage:"account_id"`
	MultipleEntryID string  `storage:"multiple_entry_id"`
	Currency        string  `storage:"currency"`
	Concept         *string `storage:"concept"`
	Date            string  `storage:"date"`
	Quarter         int     `storage:"quarter"`
	DebitBalance    float64 `storage:"debit_balance"`
	CreditBalance   float64 `storage:"credit_balance"`
	TotalBalance    float64 `storage:"total_balance"`
	State           string  `storage:"state"`
	StateReason     *string `storage:"state_reason"`
	CreationDate    string  `storage:"creation_date"`
	LastUpdate      string  `storage:"last_update"`
}

func toEntryDb(entry entry.Entry) (*EntryDb, error) {
	var db EntryDb
	db.ID = entry.ID
	db.AccountID = entry.AccountID
	db.MultipleEntryID = entry.MultipleEntryID
	db.Currency = entry.Currency
	db.Concept = entry.Concept
	db.Date = entry.Date.Format(time.RFC3339)
	db.Quarter = entry.Quarter
	db.DebitBalance = entry.DebitBalance
	db.CreditBalance = entry.CreditBalance
	db.TotalBalance = entry.TotalBalance
	db.State = entry.State
	db.StateReason = entry.StateReason
	db.CreationDate = entry.CreationDate.Format(time.RFC3339)
	db.LastUpdate = entry.LastUpdate.Format(time.RFC3339)
	return &db, nil
}

func toEntry(entryDb EntryDb) (*entry.Entry, error) {
	var entry entry.Entry
	entry.ID = entryDb.ID
	entry.AccountID = entryDb.AccountID
	entry.MultipleEntryID = entryDb.MultipleEntryID
	entry.Currency = entryDb.Currency
	entry.Concept = entryDb.Concept
	date, err := time.Parse(time.RFC3339, entryDb.Date)
	if err != nil {
		return nil, err
	}
	entry.Date = date
	entry.Quarter = entryDb.Quarter
	entry.DebitBalance = entryDb.DebitBalance
	entry.CreditBalance = entryDb.CreditBalance
	entry.TotalBalance = entryDb.TotalBalance
	entry.State = entryDb.State
	entry.StateReason = entryDb.StateReason
	creationDate, err := time.Parse(time.RFC3339, entryDb.CreationDate)
	if err != nil {
		return nil, err
	}
	entry.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, entryDb.LastUpdate)
	if err != nil {
		return nil, err
	}
	entry.LastUpdate = lastUpdate
	return &entry, nil
}

func toEntryItems(entrySlice []EntryDb) ([]entry.Entry, error) {
	entryItems := make([]entry.Entry, len(entrySlice))
	for index := range entrySlice {
		entry, err := toEntry(entrySlice[index])
		if err != nil {
			return nil, err
		}
		entryItems[index] = *entry
	}
	return entryItems, nil
}
