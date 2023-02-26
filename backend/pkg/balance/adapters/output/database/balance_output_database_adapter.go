package balancedb

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
)

func NewBalanceRepository(opts BalanceRepositoryOptions) (*BalanceRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of BalanceRepository with a nil sql.DB")
	}

	return &BalanceRepository{
		db: opts.DB,
	}, nil
}

type BalanceRepositoryOptions struct {
	DB *sql.DB
}

type BalanceRepository struct {
	db *sql.DB
}

var _ balance.BalanceStorage = (*BalanceRepository)(nil)

func (repository *BalanceRepository) AddByYear(ctx context.Context, input balance.BalanceByYearStorageData) (*balance.BalanceByYearStorageData, error) {
	var balanceDb BalanceByYearDb
	balanceDb.Project = input.Project
	balanceDb.ID = input.ID
	balanceDb.BalanceID = input.BalanceID
	balanceDb.Amount = input.Amount
	balanceDb.Year = input.Year

	insertQuery := "INSERT INTO balances_by_year (project, id, balance_id, amount, year) VALUES (?, ?, ?, ?, ?)"

	repository.db.Exec(insertQuery, balanceDb.ID, balanceDb.BalanceID, balanceDb.Amount, balanceDb.Year)

	query := "SELECT * FROM balances_by_year WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, balanceDb.Project, balanceDb.ID)

	var addedBalance BalanceByYearDb
	err := result.Scan(&addedBalance.Project, &addedBalance.ID, &addedBalance.BalanceID, &addedBalance.Amount, &addedBalance.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("balance [" + balanceDb.ID + "] not added")
		}
		return nil, err
	}

	var output balance.BalanceByYearStorageData
	output.Project = addedBalance.Project
	output.ID = addedBalance.ID
	output.BalanceID = addedBalance.BalanceID
	output.Amount = addedBalance.Amount
	output.Year = addedBalance.Year

	return &output, nil
}

func (repository *BalanceRepository) All(ctx context.Context, project string, filters balance.Filters) (*balance.BalanceStorageCollectionData, error) {
	query := BuildQueryWithFilters(filters)

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedBalances []BalanceDb
	for rows.Next() {
		var balanceDb BalanceDb
		err = rows.Scan(&balanceDb.Project, &balanceDb.ID, &balanceDb.Type, &balanceDb.Kind, &balanceDb.Parent, &balanceDb.Child, &balanceDb.Year, &balanceDb.Amount)
		if err != nil {
			return nil, err
		}

		storedBalances = append(storedBalances, balanceDb)
	}

	var items []balance.BalanceStorageData
	for _, storedBalance := range storedBalances {
		var balance balance.BalanceStorageData
		balance.Project = storedBalance.Project
		balance.ID = storedBalance.ID
		balance.Type = storedBalance.Type
		balance.Kind = storedBalance.Kind
		balance.Parent = storedBalance.Parent
		balance.Child = storedBalance.Child
		balance.Year = storedBalance.Year
		balance.Amount = storedBalance.Amount
		items = append(items, balance)
	}

	return &balance.BalanceStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *BalanceRepository) Remove(ctx context.Context, project string, id string) (*balance.BalanceByYearStorageData, error) {
	query := "SELECT * FROM balances_by_year WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedBalance BalanceByYearDb
	err := result.Scan(&storedBalance.Project, &storedBalance.ID, &storedBalance.BalanceID, &storedBalance.Amount, &storedBalance.Year)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("balance [" + id + "] not found")
		}
		return nil, err
	}

	query = "DELETE FROM balances_by_year WHERE project = ? AND id = ?"

	_ = repository.db.QueryRow(query, project, id)

	var output balance.BalanceByYearStorageData
	output.Project = storedBalance.Project
	output.ID = storedBalance.ID
	output.BalanceID = storedBalance.BalanceID
	output.Amount = storedBalance.Amount
	output.Year = storedBalance.Year

	return &output, nil
}

func (repository *BalanceRepository) Add(ctx context.Context, input balance.AddBalanceStorageData) (*balance.AddBalanceStorageData, error) {
	var addBalanceDb AddBalanceDb
	addBalanceDb.Project = input.Project
	addBalanceDb.ID = input.ID
	addBalanceDb.Type = input.Type
	addBalanceDb.Kind = input.Kind
	addBalanceDb.Parent = input.Parent
	addBalanceDb.Child = input.Child
	accounts, err := json.Marshal(input.Accounts)
	if err != nil {
		return nil, err
	}
	addBalanceDb.Accounts = string(accounts)

	insertQuery := "INSERT INTO balances (project, id, type, kind, parent, child, accounts) VALUES (?, ?, ?, ?, ?, ?, ?)"

	_, err = repository.db.Exec(insertQuery, addBalanceDb.Project, addBalanceDb.ID, addBalanceDb.Type, addBalanceDb.Kind, addBalanceDb.Parent, addBalanceDb.Child, addBalanceDb.Accounts)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM balances WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, addBalanceDb.Project, addBalanceDb.ID)

	var storedBalance AddBalanceDb
	err = result.Scan(&storedBalance.Project, &storedBalance.ID, &storedBalance.Type, &storedBalance.Kind, &storedBalance.Parent, &storedBalance.Child, &storedBalance.Accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("balance [" + addBalanceDb.ID + "] not added")
		}
		return nil, err
	}

	var output balance.AddBalanceStorageData
	output.Project = storedBalance.Project
	output.ID = storedBalance.ID
	output.Type = storedBalance.Type
	output.Kind = storedBalance.Kind
	output.Parent = storedBalance.Parent
	output.Child = storedBalance.Child
    err = json.Unmarshal([]byte(storedBalance.Accounts), &output.Accounts)
    if err != nil {
		return nil, err
    }

	return &output, nil
}

func (repository *BalanceRepository) Get(ctx context.Context, project string, id string) (*balance.AddBalanceStorageData, error) {
	query := "SELECT * FROM balances WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedBalance AddBalanceDb
	err := result.Scan(&storedBalance.Project, &storedBalance.ID, &storedBalance.Type, &storedBalance.Kind, &storedBalance.Parent, &storedBalance.Child, &storedBalance.Accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var output balance.AddBalanceStorageData
	output.Project = storedBalance.Project
	output.ID = storedBalance.ID
	output.Type = storedBalance.Type
	output.Kind = storedBalance.Kind
	output.Parent = storedBalance.Parent
	output.Child = storedBalance.Child
	err = json.Unmarshal([]byte(storedBalance.Accounts), &output.Accounts)
    if err != nil {
		return nil, err
    }

	return &output, nil
}

func BuildQueryWithFilters(filters balance.Filters) string {
	if filters.Year != nil {
		year := strconv.Itoa(*filters.Year)
		query := "SELECT balances.project, balances.id, type, kind, parent, child, year, SUM(amount) FROM balances INNER JOIN balances_by_year ON balances_by_year.balance_id = balances.id WHERE year = \"" + year + "\" AND project = ? GROUP BY balances.id, type, kind, parent, child, year;"
		return query
	}

	return ""
}
