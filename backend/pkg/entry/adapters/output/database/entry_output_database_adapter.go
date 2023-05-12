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
	entryDb.Project = input.Project
	entryDb.ID = input.ID
	entryDb.ContractID = input.ContractID
	entryDb.UserID = input.UserID
	entryDb.AccountID = input.AccountID
	entryDb.MultipleEntryID = input.MultipleEntryID
	entryDb.BalanceByYearID = input.BalanceByYearID
	entryDb.Currency = input.Currency
	entryDb.Concept = input.Concept
	entryDb.Date = input.Date.Format(time.RFC3339)
	entryDb.Quarter = input.Quarter
	entryDb.DebitBalance = input.DebitBalance
	entryDb.CreditBalance = input.CreditBalance
	entryDb.State = input.State
	entryDb.StateReason = input.StateReason
	entryDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	entryDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	insertQuery := "INSERT INTO entries (project, id, contract_id, user_id, account_id, multiple_entry_id, balance_by_year_id, currency, concept, date, quarter, debit_balance, credit_balance, state, state_reason, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(insertQuery, entryDb.Project, entryDb.ID, entryDb.ContractID, entryDb.UserID, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.BalanceByYearID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate)
	if err != nil {
		return nil, err
	}

	addedEntry, err := repository.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	return addedEntry, nil
}

func (repository *EntryRepository) Get(ctx context.Context, project string, id string) (*entry.EntryStorageData, error) {
	query := "SELECT * FROM entries WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedEntry EntryDb
	err := result.Scan(&storedEntry.Project, &storedEntry.ID, &storedEntry.ContractID, &storedEntry.UserID, &storedEntry.AccountID, &storedEntry.MultipleEntryID, &storedEntry.BalanceByYearID, &storedEntry.Currency, &storedEntry.Concept, &storedEntry.Date, &storedEntry.Quarter, &storedEntry.DebitBalance, &storedEntry.CreditBalance, &storedEntry.State, &storedEntry.StateReason, &storedEntry.CreationDate, &storedEntry.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("entry with project [" + project + "] and id [" + id + "] not found")
		}
		return nil, err
	}

	var output entry.EntryStorageData
	output.Project = storedEntry.Project
	output.ID = storedEntry.ID
	output.ContractID = storedEntry.ContractID
	output.UserID = storedEntry.UserID
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

func (repository *EntryRepository) All(ctx context.Context, project string, filters entry.Filters) (*entry.EntryStorageCollectionData, error) {
	query := BuildQueryWithFilters("SELECT * FROM entries WHERE project = ?", filters)

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedEntries []EntryDb
	for rows.Next() {
		var entryDb EntryDb
		err = rows.Scan(&entryDb.Project, &entryDb.ID, &entryDb.ContractID, &entryDb.UserID, &entryDb.AccountID, &entryDb.MultipleEntryID, &entryDb.BalanceByYearID, &entryDb.Currency, &entryDb.Concept, &entryDb.Date, &entryDb.Quarter, &entryDb.DebitBalance, &entryDb.CreditBalance, &entryDb.State, &entryDb.StateReason, &entryDb.CreationDate, &entryDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		storedEntries = append(storedEntries, entryDb)
	}

	var items []entry.EntryStorageData
	for _, storedEntry := range storedEntries {
		var entry entry.EntryStorageData
		entry.Project = storedEntry.Project
		entry.ID = storedEntry.ID
		entry.ContractID = storedEntry.ContractID
		entry.UserID = storedEntry.UserID
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
	entryDb.Project = input.Project
	entryDb.ID = input.ID
	entryDb.ContractID = input.ContractID
	entryDb.UserID = input.UserID
	entryDb.AccountID = input.AccountID
	entryDb.MultipleEntryID = input.MultipleEntryID
	entryDb.BalanceByYearID = input.BalanceByYearID
	entryDb.Currency = input.Currency
	entryDb.Concept = input.Concept
	entryDb.Date = input.Date.Format(time.RFC3339)
	entryDb.Quarter = input.Quarter
	entryDb.DebitBalance = input.DebitBalance
	entryDb.CreditBalance = input.CreditBalance
	entryDb.State = input.State
	entryDb.StateReason = input.StateReason
	entryDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	entryDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	updateQuery := "UPDATE entries SET contract_id = ?, user_id = ?, account_id = ?, multiple_entry_id = ?, balance_by_year_id = ?, currency = ?, concept = ?, date = ?, quarter = ?, debit_balance = ?, credit_balance = ?, state = ?, state_reason = ?, creation_date = ?, last_update = ? WHERE project = ? AND id = ?"

	_, err := repository.db.Exec(updateQuery, entryDb.ContractID, entryDb.UserID, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.BalanceByYearID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate, entryDb.Project, entryDb.ID)
	if err != nil {
		return nil, err
	}

	storedEntry, err := repository.Get(ctx, input.Project, input.ID)
	if err != nil {
		return nil, err
	}

	return storedEntry, nil
}

func (repository *EntryRepository) Remove(ctx context.Context, project string, id string) (*entry.EntryStorageData, error) {
	storedEntry, err := repository.Get(ctx, project, id)
	if err != nil {
		return nil, err
	}

	query := "DELETE FROM entries WHERE project = ? AND id = ?"

	_, err = repository.db.Exec(query, project, id)
	if err != nil {
		return nil, err
	}

	return storedEntry, nil
}

func BuildQueryWithFilters(query string, filters entry.Filters) string {
	if filters.MultipleEntryID != nil {
		query = query + " AND multiple_entry_id = \"" + *filters.MultipleEntryID + "\""
		return query
	}
	if filters.AccountID != nil {
		query = query + " AND account_id = \"" + *filters.AccountID + "\""
		return query
	}

	return query
}
