package accountdb

import (
	"context"
	"database/sql"
	"errors"

	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
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

func (repository *AccountRepository) Get(ctx context.Context, project string, id string) (*account.AccountStorageData, error) {
	query := "SELECT * FROM accounts WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedAccount AccountDb
	err := result.Scan(&storedAccount.Project, &storedAccount.ID, &storedAccount.Name, &storedAccount.Level, &storedAccount.BalanceID)
	if err != nil {
		return nil, err
	}

	var output account.AccountStorageData
	output.Project = storedAccount.Project
	output.ID = storedAccount.ID
	output.Name = storedAccount.Name
	output.Level = storedAccount.Level
	output.BalanceID = storedAccount.BalanceID

	return &output, nil
}

func (repository *AccountRepository) All(ctx context.Context, project string) (*account.AccountStorageCollectionData, error) {
	query := "SELECT * FROM accounts WHERE project = ?"

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedAccounts []AccountDb
	for rows.Next() {
		var accountDb AccountDb
		err = rows.Scan(&accountDb.Project, &accountDb.ID, &accountDb.Name, &accountDb.Level, &accountDb.BalanceID)
		if err != nil {
			return nil, err
		}

		storedAccounts = append(storedAccounts, accountDb)
	}

	var items []account.AccountStorageData
	for _, storedAccount := range storedAccounts {
		var account account.AccountStorageData
		account.Project = storedAccount.Project
		account.ID = storedAccount.ID
		account.Name = storedAccount.Name
		account.Level = storedAccount.Level
		account.BalanceID = storedAccount.BalanceID
		items = append(items, account)
	}

	return &account.AccountStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *AccountRepository) Edit(ctx context.Context, input account.AccountStorageData) (*account.AccountStorageData, error) {
	var accountDb AccountDb
	accountDb.Project = input.Project
	accountDb.ID = input.ID
	accountDb.Name = input.Name
	accountDb.Level = input.Level
	accountDb.BalanceID = input.BalanceID

	updateQuery := "UPDATE accounts SET name = ?, level = ?, balance_id = ? WHERE project = ? AND id = ?"

	repository.db.QueryRow(updateQuery, accountDb.Name, accountDb.Level, accountDb.BalanceID, accountDb.Project, accountDb.ID)

	query := "SELECT * FROM accounts WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, accountDb.Project, accountDb.ID)

	var storedAccount AccountDb
	err := result.Scan(&storedAccount.Project, &storedAccount.ID, &storedAccount.Name, &storedAccount.Level, &storedAccount.BalanceID)
	if err != nil {
		return nil, err
	}

	var output account.AccountStorageData
	output.Project = storedAccount.Project
	output.ID = storedAccount.ID
	output.Name = storedAccount.Name
	output.Level = storedAccount.Level
	output.BalanceID = storedAccount.BalanceID

	return &output, nil
}

func (repository *AccountRepository) Add(ctx context.Context, input account.AccountStorageData) (*account.AccountStorageData, error) {
	var accountDb AccountDb
	accountDb.Project = input.Project
	accountDb.ID = input.ID
	accountDb.Name = input.Name
	accountDb.Level = input.Level
	accountDb.BalanceID = input.BalanceID

	insertQuery := "INSERT INTO accounts (project, id, name, level, balance_id) VALUES (?, ?, ?, ?, ?)"

	repository.db.Exec(insertQuery, accountDb.Project, accountDb.ID, accountDb.Name, accountDb.Level, accountDb.BalanceID)

	query := "SELECT * FROM accounts WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, accountDb.Project, accountDb.ID)

	var storedAccount AccountDb
	err := result.Scan(&storedAccount.Project, &storedAccount.ID, &storedAccount.Name, &storedAccount.Level, &storedAccount.BalanceID)
	if err != nil {
		return nil, err
	}

	var output account.AccountStorageData
	output.Project = storedAccount.Project
	output.ID = storedAccount.ID
	output.Name = storedAccount.Name
	output.Level = storedAccount.Level
	output.BalanceID = storedAccount.BalanceID

	return &output, nil
}
