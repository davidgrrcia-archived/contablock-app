package entrydb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
)

func NewEntryRepository(opts EntryRepositoryOptions) (*EntryRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of EntryRepository with a nil sql.DB")
	}

	return &EntryRepository{
		db: opts.DB,
	}, nil
}

type EntryRepositoryOptions struct {
	DB *sql.DB
}

type EntryRepository struct {
	db *sql.DB
}

var _ entry.EntryStorage = (*EntryRepository)(nil)

func (repository *EntryRepository) Add(ctx context.Context, input entry.EntryStorageData) (*entry.EntryStorageData, error) {
	var entryDb EntryDb
	entryDb.ID = input.ID
	entryDb.AccountID = input.AccountID
	entryDb.MultipleEntryID = input.MultipleEntryID
	entryDb.BalanceByYearID = input.BalanceByYearID
	entryDb.Currency = input.Currency
	entryDb.Concept = input.Concept
	entryDb.Date = input.Date.Format(time.RFC3339)
	entryDb.Quarter = input.Quarter
	entryDb.DebitBalance = input.DebitBalance
	entryDb.CreditBalance = input.CreditBalance
	entryDb.TotalBalance = input.TotalBalance
	entryDb.State = input.State
	entryDb.StateReason = input.StateReason
	entryDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	entryDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	insertQuery := "INSERT INTO entries (id, account_id, multiple_entry_id, balance_by_year_id, currency, concept, date, quarter, debit_balance, credit_balance, total_balance, state, state_reason, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	repository.db.QueryRow(insertQuery, entryDb.ID, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.BalanceByYearID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.TotalBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate)

	query := "SELECT * FROM entries WHERE id = ?"

	result := repository.db.QueryRow(query, entryDb.ID)

	var addedEntry EntryDb
	err := result.Scan(&addedEntry.ID, &addedEntry.AccountID, &addedEntry.MultipleEntryID, &addedEntry.BalanceByYearID, &addedEntry.Currency, &addedEntry.Concept, &addedEntry.Date, &addedEntry.Quarter, &addedEntry.DebitBalance, &addedEntry.CreditBalance, &addedEntry.TotalBalance, &addedEntry.State, &addedEntry.StateReason, &addedEntry.CreationDate, &addedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}

	var output entry.EntryStorageData
	output.ID = addedEntry.ID
	output.AccountID = addedEntry.AccountID
	output.MultipleEntryID = addedEntry.MultipleEntryID
	output.BalanceByYearID = addedEntry.BalanceByYearID
	output.Currency = addedEntry.Currency
	output.Concept = addedEntry.Concept
	date, err := time.Parse(time.RFC3339, addedEntry.Date)
	if err != nil {
		return nil, err
	}
	output.Date = date
	output.Quarter = addedEntry.Quarter
	output.DebitBalance = addedEntry.DebitBalance
	output.CreditBalance = addedEntry.CreditBalance
	output.TotalBalance = addedEntry.TotalBalance
	output.State = addedEntry.State
	output.StateReason = addedEntry.StateReason
	creationDate, err := time.Parse(time.RFC3339, addedEntry.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, addedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *EntryRepository) Get(ctx context.Context, id string) (*entry.EntryStorageData, error) {
	query := "SELECT * FROM entries WHERE id = ?"

	result := repository.db.QueryRow(query, id)

	var storedEntry EntryDb
	err := result.Scan(&storedEntry.ID, &storedEntry.AccountID, &storedEntry.MultipleEntryID, &storedEntry.BalanceByYearID, &storedEntry.Currency, &storedEntry.Concept, &storedEntry.Date, &storedEntry.Quarter, &storedEntry.DebitBalance, &storedEntry.CreditBalance, &storedEntry.TotalBalance, &storedEntry.State, &storedEntry.StateReason, &storedEntry.CreationDate, &storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}

	var output entry.EntryStorageData
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.BalanceByYearID = storedEntry.BalanceByYearID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	date, err := time.Parse(time.RFC3339, storedEntry.Date)
	if err != nil {
		return nil, err
	}
	output.Date = date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	creationDate, err := time.Parse(time.RFC3339, storedEntry.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *EntryRepository) All(ctx context.Context, filters entry.Filters) (*entry.EntryStorageCollectionData, error) {
	query := BuildQueryWithFilters("SELECT * FROM entries", filters)

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var storedEntries []EntryDb
	for rows.Next() {
		var entryDb EntryDb
		err = rows.Scan(&entryDb.ID, &entryDb.AccountID, &entryDb.MultipleEntryID, &entryDb.BalanceByYearID, &entryDb.Currency, &entryDb.Concept, &entryDb.Date, &entryDb.Quarter, &entryDb.DebitBalance, &entryDb.CreditBalance, &entryDb.TotalBalance, &entryDb.State, &entryDb.StateReason, &entryDb.CreationDate, &entryDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		storedEntries = append(storedEntries, entryDb)
	}

	var items []entry.EntryStorageData
	for _, storedEntry := range storedEntries {
		var entry entry.EntryStorageData
		entry.ID = storedEntry.ID
		entry.AccountID = storedEntry.AccountID
		entry.MultipleEntryID = storedEntry.MultipleEntryID
		entry.BalanceByYearID = storedEntry.BalanceByYearID
		entry.Currency = storedEntry.Currency
		entry.Concept = storedEntry.Concept
		date, err := time.Parse(time.RFC3339, storedEntry.Date)
		if err != nil {
			return nil, err
		}
		entry.Date = date
		entry.Quarter = storedEntry.Quarter
		entry.DebitBalance = storedEntry.DebitBalance
		entry.CreditBalance = storedEntry.CreditBalance
		entry.TotalBalance = storedEntry.TotalBalance
		entry.State = storedEntry.State
		entry.StateReason = storedEntry.StateReason
		creationDate, err := time.Parse(time.RFC3339, storedEntry.CreationDate)
		if err != nil {
			return nil, err
		}
		entry.CreationDate = creationDate
		lastUpdate, err := time.Parse(time.RFC3339, storedEntry.LastUpdate)
		if err != nil {
			return nil, err
		}
		entry.LastUpdate = lastUpdate
		items = append(items, entry)
	}

	return &entry.EntryStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *EntryRepository) Edit(ctx context.Context, input entry.EntryStorageData) (*entry.EntryStorageData, error) {
	var entryDb EntryDb
	entryDb.ID = input.ID
	entryDb.AccountID = input.AccountID
	entryDb.MultipleEntryID = input.MultipleEntryID
	entryDb.BalanceByYearID = input.BalanceByYearID
	entryDb.Currency = input.Currency
	entryDb.Concept = input.Concept
	entryDb.Date = input.Date.Format(time.RFC3339)
	entryDb.Quarter = input.Quarter
	entryDb.DebitBalance = input.DebitBalance
	entryDb.CreditBalance = input.CreditBalance
	entryDb.TotalBalance = input.TotalBalance
	entryDb.State = input.State
	entryDb.StateReason = input.StateReason
	entryDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	entryDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	updateQuery := "UPDATE entries SET account_id = ?, multiple_entry_id = ?, balance_by_year_id = ?, currency = ?, concept = ?, date = ?, quarter = ?, debit_balance = ?, credit_balance = ?, total_balance = ?, state = ?, state_reason = ?, creation_date = ?, last_update = ? WHERE id = ?"

	repository.db.QueryRow(updateQuery, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.BalanceByYearID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.TotalBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate, entryDb.ID)

	query := "SELECT * FROM entries WHERE id = ?"

	result := repository.db.QueryRow(query, entryDb.ID)

	var storedEntry EntryDb
	err := result.Scan(&storedEntry.ID, &storedEntry.AccountID, &storedEntry.MultipleEntryID, &storedEntry.BalanceByYearID, &storedEntry.Currency, &storedEntry.Concept, &storedEntry.Date, &storedEntry.Quarter, &storedEntry.DebitBalance, &storedEntry.CreditBalance, &storedEntry.TotalBalance, &storedEntry.State, &storedEntry.StateReason, &storedEntry.CreationDate, &storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}

	var output entry.EntryStorageData
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.BalanceByYearID = storedEntry.BalanceByYearID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	date, err := time.Parse(time.RFC3339, storedEntry.Date)
	if err != nil {
		return nil, err
	}
	output.Date = date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	creationDate, err := time.Parse(time.RFC3339, storedEntry.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *EntryRepository) Remove(ctx context.Context, id string) (*entry.EntryStorageData, error) {
	query := "SELECT * FROM entries WHERE id = ?"

	result := repository.db.QueryRow(query, id)

	var storedEntry EntryDb
	err := result.Scan(&storedEntry.ID, &storedEntry.AccountID, &storedEntry.MultipleEntryID, &storedEntry.BalanceByYearID, &storedEntry.Currency, &storedEntry.Concept, &storedEntry.Date, &storedEntry.Quarter, &storedEntry.DebitBalance, &storedEntry.CreditBalance, &storedEntry.TotalBalance, &storedEntry.State, &storedEntry.StateReason, &storedEntry.CreationDate, &storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}

	query = "DELETE FROM entries WHERE id = ?"

	_ = repository.db.QueryRow(query, id)

	var output entry.EntryStorageData
	output.ID = storedEntry.ID
	output.AccountID = storedEntry.AccountID
	output.MultipleEntryID = storedEntry.MultipleEntryID
	output.BalanceByYearID = storedEntry.BalanceByYearID
	output.Currency = storedEntry.Currency
	output.Concept = storedEntry.Concept
	date, err := time.Parse(time.RFC3339, storedEntry.Date)
	if err != nil {
		return nil, err
	}
	output.Date = date
	output.Quarter = storedEntry.Quarter
	output.DebitBalance = storedEntry.DebitBalance
	output.CreditBalance = storedEntry.CreditBalance
	output.TotalBalance = storedEntry.TotalBalance
	output.State = storedEntry.State
	output.StateReason = storedEntry.StateReason
	creationDate, err := time.Parse(time.RFC3339, storedEntry.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedEntry.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func BuildQueryWithFilters(query string, filters entry.Filters) string {
	if filters.MultipleEntryID != nil {
		query = query + " WHERE multiple_entry_id = \"" + *filters.MultipleEntryID + "\""
	}

	return query
}
