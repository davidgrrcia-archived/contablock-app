package businessdb

import (
	"context"
	"database/sql"
	"errors"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
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

	var outputDb BusinessDb
	err := result.Scan(&outputDb.ID, &outputDb.Name, &outputDb.SocialDenomination, &outputDb.RegisteredOffice, &outputDb.Cif, &outputDb.Logo, &outputDb.EthAddress)
	if err != nil {
		return nil, err
	}

	var output business.BusinessStorageData
	output.ID = outputDb.ID
	output.Name = outputDb.Name
	output.SocialDenomination = outputDb.SocialDenomination
	output.RegisteredOffice = outputDb.RegisteredOffice
	output.Cif = outputDb.Cif
	output.Logo = outputDb.Logo
	output.EthAddress = outputDb.EthAddress

	return &output, nil
}

func (repository *BusinessRepository) All(ctx context.Context) (*business.BusinessStorageCollectionData, error) {
	query := "SELECT * FROM business"

	rows, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	var storedBusiness []BusinessDb
	for rows.Next() {
		var businessDb BusinessDb
		err = rows.Scan(&businessDb.ID, &businessDb.Name, &businessDb.SocialDenomination, &businessDb.RegisteredOffice, &businessDb.Cif, &businessDb.Logo, &businessDb.EthAddress)
		if err != nil {
			return nil, err
		}
		storedBusiness = append(storedBusiness, businessDb)
	}

	var items []business.BusinessStorageData
	for _, storedAccount := range storedBusiness {
		var business business.BusinessStorageData
		business.ID = storedAccount.ID
		business.Name = storedAccount.Name
		business.SocialDenomination = storedAccount.SocialDenomination
		business.RegisteredOffice = storedAccount.RegisteredOffice
		business.Cif = storedAccount.Cif
		business.Logo = storedAccount.Logo
		business.EthAddress = storedAccount.EthAddress
		items = append(items, business)
	}

	return &business.BusinessStorageCollectionData{
		Items: items,
	}, nil
}
