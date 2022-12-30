package accountdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
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
	query := "SELECT * FROM accounts WHERE id = ?"

	result := repository.db.QueryRow(query, id)

	var storedAccount AccountDb
	err := result.Scan(&storedAccount.ID, &storedAccount.Name, &storedAccount.Level, &storedAccount.Type)
	if err != nil {
		return nil, err
	}

	var output account.AccountStorageData
	output.ID = storedAccount.ID
	output.Name = storedAccount.Name
	output.Level = storedAccount.Level
	output.Type = &storedAccount.Type

	return &output, nil
}

func (repository *AccountRepository) All(ctx context.Context) (*account.AccountStorageCollectionData, error) {
	query := "SELECT * FROM accounts"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var storedAccounts []AccountDb
	for rows.Next() {
		var accountDb AccountDb
		err = rows.Scan(&accountDb.ID, &accountDb.Name, &accountDb.Level, &accountDb.Type)
		if err != nil {
			return nil, err
		}

		storedAccounts = append(storedAccounts, accountDb)
	}

	var items []account.AccountStorageData
	for _, storedAccount := range storedAccounts {
		var account account.AccountStorageData
		account.ID = storedAccount.ID
		account.Name = storedAccount.Name
		account.Level = storedAccount.Level
		account.Type = &storedAccount.Type
		items = append(items, account)
	}

	return &account.AccountStorageCollectionData{
		Items: items,
	}, nil
}
