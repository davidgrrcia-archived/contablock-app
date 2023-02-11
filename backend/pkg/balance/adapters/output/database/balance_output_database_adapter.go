package balancedb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
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

func (repository *BalanceRepository) All(ctx context.Context, filters balance.Filters) (*balance.BalanceStorageCollectionData, error) {
	//query := BuildQueryWithFilters("SELECT * FROM balances INNER JOIN balances_by_year ON balances_by_year.balance_id=balances.id", filters)
	query := "SELECT * FROM `balances`"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var storedBalances []BalanceDb
	for rows.Next() {
		var balanceDb BalanceDb
		err = rows.Scan(&balanceDb.ID, &balanceDb.Type, &balanceDb.Kind, &balanceDb.Parent, &balanceDb.Child, &balanceDb.Accounts)
		if err != nil {
			return nil, err
		}

		storedBalances = append(storedBalances, balanceDb)
	}

	var items []balance.BalanceStorageData
	for _, storedBalance := range storedBalances {
		var balance balance.BalanceStorageData
		balance.ID = storedBalance.ID
		balance.Type = storedBalance.Type
		balance.Kind = storedBalance.Kind
		balance.Parent = storedBalance.Parent
		balance.Child = storedBalance.Child
		balance.Accounts = storedBalance.Accounts
		items = append(items, balance)
	}

	return &balance.BalanceStorageCollectionData{
		Items: items,
	}, nil
}

func BuildQueryWithFilters(query string, filters balance.Filters) string {
	if filters.Year != nil {
		query = query + " WHERE year = \"" + *filters.Year + "\""
	}

	return query
}
