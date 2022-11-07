package accountdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/account"
)

func NewAccountRepository(opts AccountRepositoryOptions) (*AccountRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of AccountRepository with a nil sql.DB")
	}

	return &AccountRepository{
		db: opts.DB,
	}, nil
}

type AccountRepositoryOptions struct {
	DB *sql.DB
}

type AccountRepository struct {
	db *sql.DB
}

var _ account.AccountStorage = (*AccountRepository)(nil)

func (repository *AccountRepository) Get(ctx context.Context, id string) (*account.AccountStorageData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "SELECT * FROM accounts WHERE id = ?"

	result := db.QueryRow(query, id)

	var output AccountDb
	err = result.Scan(&output.ID, &output.Name, &output.Level, &output.Type)

	storedAccount, err := toAccount(output)
	if err != nil {
		return nil, err
	}

	return &account.AccountStorageData{
		Account: *storedAccount,
	}, nil
}

func (repository *AccountRepository) All(ctx context.Context) (*account.AccountStorageCollectionData, error) {
	db, err := repository.openConnection()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "SELECT * FROM accounts"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	var output []AccountDb
	for rows.Next() {
		var accountDb AccountDb
		err = rows.Scan(&accountDb.ID, &accountDb.Name, &accountDb.Level, &accountDb.Type)

		output = append(output, accountDb)
	}

	items, err := toAccountItems(output)
	if err != nil {
		return nil, err
	}

	return &account.AccountStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *AccountRepository) openConnection() (*sql.DB, error) {
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
