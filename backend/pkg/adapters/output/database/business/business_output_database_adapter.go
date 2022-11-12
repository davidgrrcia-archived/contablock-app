package businessdb

import (
	"context"
	"database/sql"
	"errors"

	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/business"
)

func NewBusinessRepository(opts BusinessRepositoryOptions) (*BusinessRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of BusinessRepository with a nil sql.DB")
	}

	return &BusinessRepository{
		db: opts.DB,
	}, nil
}

type BusinessRepositoryOptions struct {
	DB *sql.DB
}

type BusinessRepository struct {
	db *sql.DB
}

var _ business.BusinessStorage = (*BusinessRepository)(nil)

func (repository *BusinessRepository) Get(ctx context.Context, id string) (*business.BusinessStorageData, error) {
	query := "SELECT * FROM business WHERE id = ?"

	result := repository.db.QueryRow(query, id)

	var output BusinessDb
	err := result.Scan(&output.ID, &output.Name, &output.SocialDenomination, &output.RegisteredOffice, &output.Cif, &output.Logo, &output.EthAddress)

	storedBusiness, err := toBusiness(output)
	if err != nil {
		return nil, err
	}

	return &business.BusinessStorageData{
		Business: *storedBusiness,
	}, nil
}

func (repository *BusinessRepository) All(ctx context.Context) (*business.BusinessStorageCollectionData, error) {
	query := "SELECT * FROM business"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var output []BusinessDb
	for rows.Next() {
		var businessDb BusinessDb
		err = rows.Scan(&businessDb.ID, &businessDb.Name, &businessDb.SocialDenomination, &businessDb.RegisteredOffice, &businessDb.Cif, &businessDb.Logo, &businessDb.EthAddress)

		output = append(output, businessDb)
	}

	items, err := toBusinessItems(output)
	if err != nil {
		return nil, err
	}

	return &business.BusinessStorageCollectionData{
		Items: items,
	}, nil
}
