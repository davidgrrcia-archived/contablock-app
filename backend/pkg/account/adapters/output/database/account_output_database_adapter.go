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

	var output AccountDb
	err := result.Scan(&output.ID, &output.Name, &output.Level, &output.Type)

	storedAccount, err := toAccount(output)
	if err != nil {
		return nil, err
	}

	return &account.AccountStorageData{
		Account: *storedAccount,
	}, nil
}

func (repository *AccountRepository) All(ctx context.Context) (*account.AccountStorageCollectionData, error) {
	query := "SELECT * FROM accounts"

	rows, err := repository.db.Query(query)
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
