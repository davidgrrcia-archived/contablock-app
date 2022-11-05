package entrydb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
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
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	entryDb, err := toEntryDb(input.Entry)
	if err != nil {
		return nil, err
	}

	query := "INSERT INTO entries (id, account_id, multiple_entry_id, currency, concept, date, quarter, debit_balance, credit_balance, total_balance, state, state_reason, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result := db.QueryRow(query, entryDb.ID, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.TotalBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate)

	query = "SELECT * FROM entries WHERE id = ?"

	result = db.QueryRow(query, entryDb.ID)

	var output EntryDb
	err = result.Scan(&output.ID, &output.AccountID, &output.MultipleEntryID, &output.Currency, &output.Concept, &output.Date, &output.Quarter, &output.DebitBalance, &output.CreditBalance, &output.TotalBalance, &output.State, &output.StateReason, &output.CreationDate, &output.LastUpdate)
	if err != nil {
		return nil, err
	}

	storedEntry, err := toEntry(output)
	if err != nil {
		return nil, err
	}

	return &entry.EntryStorageData{
		Entry: *storedEntry,
	}, nil
}

func (repository *EntryRepository) Get(ctx context.Context, id string) (*entry.EntryStorageData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "SELECT * FROM entries WHERE id = ?"

	result := db.QueryRow(query, id)

	var output EntryDb
	err = result.Scan(&output.ID, &output.AccountID, &output.MultipleEntryID, &output.Currency, &output.Concept, &output.Date, &output.Quarter, &output.DebitBalance, &output.CreditBalance, &output.TotalBalance, &output.State, &output.StateReason, &output.CreationDate, &output.LastUpdate)
	if err != nil {
		return nil, err
	}

	storedEntry, err := toEntry(output)
	if err != nil {
		return nil, err
	}

	return &entry.EntryStorageData{
		Entry: *storedEntry,
	}, nil
}

func (repository *EntryRepository) All(ctx context.Context) (*entry.EntryStorageCollectionData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "SELECT * FROM entries"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var output []EntryDb
	for rows.Next() {
		var entryDb EntryDb
		err = rows.Scan(&entryDb.ID, &entryDb.AccountID, &entryDb.MultipleEntryID, &entryDb.Currency, &entryDb.Concept, &entryDb.Date, &entryDb.Quarter, &entryDb.DebitBalance, &entryDb.CreditBalance, &entryDb.TotalBalance, &entryDb.State, &entryDb.StateReason, &entryDb.CreationDate, &entryDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		output = append(output, entryDb)
	}

	items, err := toEntryItems(output)
	if err != nil {
		return nil, err
	}

	return &entry.EntryStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *EntryRepository) Edit(ctx context.Context, input entry.EntryStorageData) (*entry.EntryStorageData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	entryDb, err := toEntryDb(input.Entry)
	if err != nil {
		return nil, err
	}

	query := "UPDATE entries SET account_id = ?, multiple_entry_id = ?, currency = ?, concept = ?, date = ?, quarter = ?, debit_balance = ?, credit_balance = ?, total_balance = ?, state = ?, state_reason = ?, creation_date = ?, last_update = ? WHERE id = ?"

	result := db.QueryRow(query, entryDb.AccountID, entryDb.MultipleEntryID, entryDb.Currency, entryDb.Concept, entryDb.Date, entryDb.Quarter, entryDb.DebitBalance, entryDb.CreditBalance, entryDb.TotalBalance, entryDb.State, entryDb.StateReason, entryDb.CreationDate, entryDb.LastUpdate, entryDb.ID)

	query = "SELECT * FROM entries WHERE id = ?"

	result = db.QueryRow(query, entryDb.ID)

	var output EntryDb
	err = result.Scan(&output.ID, &output.AccountID, &output.MultipleEntryID, &output.Currency, &output.Concept, &output.Date, &output.Quarter, &output.DebitBalance, &output.CreditBalance, &output.TotalBalance, &output.State, &output.StateReason, &output.CreationDate, &output.LastUpdate)
	if err != nil {
		return nil, err
	}

	storedEntry, err := toEntry(output)
	if err != nil {
		return nil, err
	}

	return &entry.EntryStorageData{
		Entry: *storedEntry,
	}, nil
}

func (repository *EntryRepository) Remove(ctx context.Context, id string) (*entry.EntryStorageData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "SELECT * FROM entries WHERE id = ?"

	result := db.QueryRow(query, id)

	query = "DELETE FROM entries WHERE id = ?"

	_ = db.QueryRow(query, id)

	var output EntryDb
	err = result.Scan(&output.ID, &output.AccountID, &output.MultipleEntryID, &output.Currency, &output.Concept, &output.Date, &output.Quarter, &output.DebitBalance, &output.CreditBalance, &output.TotalBalance, &output.State, &output.StateReason, &output.CreationDate, &output.LastUpdate)
	if err != nil {
		return nil, err
	}

	storedEntry, err := toEntry(output)
	if err != nil {
		return nil, err
	}

	return &entry.EntryStorageData{
		Entry: *storedEntry,
	}, nil
}

func (repository *EntryRepository) openConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_accounting")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
