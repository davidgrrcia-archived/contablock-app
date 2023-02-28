package businessdb

import (
	"context"
	"database/sql"
	"errors"

	business "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
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

func (repository *BusinessRepository) Get(ctx context.Context, project string, id string) (*business.BusinessStorageData, error) {
	query := "SELECT * FROM business WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var outputDb BusinessDb
	err := result.Scan(&outputDb.Project, &outputDb.ID, &outputDb.Name, &outputDb.SocialDenomination, &outputDb.RegisteredOffice, &outputDb.Cif, &outputDb.Logo, &outputDb.EthAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var output business.BusinessStorageData
	output.Project = outputDb.Project
	output.ID = outputDb.ID
	output.Name = outputDb.Name
	output.SocialDenomination = outputDb.SocialDenomination
	output.RegisteredOffice = outputDb.RegisteredOffice
	output.Cif = outputDb.Cif
	output.Logo = outputDb.Logo
	output.EthAddress = outputDb.EthAddress

	return &output, nil
}

func (repository *BusinessRepository) All(ctx context.Context, project string) (*business.BusinessStorageCollectionData, error) {
	query := "SELECT * FROM business WHERE project = ?"

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedBusiness []BusinessDb
	for rows.Next() {
		var businessDb BusinessDb
		err = rows.Scan(&businessDb.Project, &businessDb.ID, &businessDb.Name, &businessDb.SocialDenomination, &businessDb.RegisteredOffice, &businessDb.Cif, &businessDb.Logo, &businessDb.EthAddress)
		if err != nil {
			return nil, err
		}
		storedBusiness = append(storedBusiness, businessDb)
	}

	var items []business.BusinessStorageData
	for _, storedBusiness := range storedBusiness {
		var business business.BusinessStorageData
		business.Project = storedBusiness.Project
		business.ID = storedBusiness.ID
		business.Name = storedBusiness.Name
		business.SocialDenomination = storedBusiness.SocialDenomination
		business.RegisteredOffice = storedBusiness.RegisteredOffice
		business.Cif = storedBusiness.Cif
		business.Logo = storedBusiness.Logo
		business.EthAddress = storedBusiness.EthAddress
		items = append(items, business)
	}

	return &business.BusinessStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *BusinessRepository) Add(ctx context.Context, input business.BusinessStorageData) (*business.BusinessStorageData, error) {
	var businessDb BusinessDb
	businessDb.Project = input.Project
	businessDb.ID = input.ID
	businessDb.Name = input.Name
	businessDb.SocialDenomination = input.SocialDenomination
	businessDb.RegisteredOffice = input.RegisteredOffice
	businessDb.Cif = input.Cif
	businessDb.Logo = input.Logo
	businessDb.EthAddress = input.EthAddress

	insertQuery := "INSERT INTO business (project, id, name, social_denomination, registered_office, cif, logo, eth_address) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(insertQuery, businessDb.Project, businessDb.ID, businessDb.Name, businessDb.SocialDenomination, businessDb.RegisteredOffice, businessDb.Cif, businessDb.Logo, businessDb.EthAddress)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM business WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, businessDb.Project, businessDb.ID)

	var storedBusiness BusinessDb
	err = result.Scan(&storedBusiness.Project, &storedBusiness.ID, &storedBusiness.Name, &storedBusiness.SocialDenomination, &storedBusiness.RegisteredOffice, &storedBusiness.Cif, &storedBusiness.Logo, &storedBusiness.EthAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("business with project " + businessDb.Project + " and id " + businessDb.ID + " not added")
		}
		return nil, err
	}

	var output business.BusinessStorageData
	output.Project = storedBusiness.Project
	output.ID = storedBusiness.ID
	output.Name = storedBusiness.Name
	output.SocialDenomination = storedBusiness.SocialDenomination
	output.RegisteredOffice = storedBusiness.RegisteredOffice
	output.Cif = storedBusiness.Cif
	output.Logo = storedBusiness.Logo
	output.EthAddress = storedBusiness.EthAddress

	return &output, nil
}
