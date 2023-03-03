package contractdb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
)

func NewContractRepository(opts ContractRepositoryOptions) (*ContractRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of ContractRepository with a nil sql.DB")
	}

	return &ContractRepository{
		db: opts.DB,
	}, nil
}

type ContractRepositoryOptions struct {
	DB *sql.DB
}

type ContractRepository struct {
	db *sql.DB
}

var _ contract.ContractStorage = (*ContractRepository)(nil)

func (repository *ContractRepository) Get(ctx context.Context, project string, id string) (*contract.ContractStorageData, error) {
	query := "SELECT * FROM contracts WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedContract ContractDb
	err := result.Scan(&storedContract.Project, &storedContract.ID, &storedContract.Name, &storedContract.Description, &storedContract.EthAddress, &storedContract.Network, &storedContract.AdminID, &storedContract.CreationDate, &storedContract.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var output contract.ContractStorageData
	output.Project = storedContract.Project
	output.ID = storedContract.ID
	output.Name = storedContract.Name
	output.Description = storedContract.Description
	output.EthAddress = storedContract.EthAddress
	output.Network = storedContract.Network
	output.AdminID = storedContract.AdminID
	creationDate, err := time.Parse(time.RFC3339, storedContract.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedContract.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *ContractRepository) All(ctx context.Context, project string, filters contract.Filters) (*contract.ContractStorageCollectionData, error) {
	query := BuildQueryWithFilters("SELECT * FROM contracts WHERE project = ?", filters)

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedContracts []ContractDb
	for rows.Next() {
		var contractDb ContractDb
		err = rows.Scan(&contractDb.Project, &contractDb.ID, &contractDb.Name, &contractDb.Description, &contractDb.EthAddress, &contractDb.Network, &contractDb.AdminID, &contractDb.CreationDate, &contractDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		storedContracts = append(storedContracts, contractDb)
	}

	var items []contract.ContractStorageData
	for _, storedContract := range storedContracts {
		var contract contract.ContractStorageData
		contract.Project = storedContract.Project
		contract.ID = storedContract.ID
		contract.Name = storedContract.Name
		contract.Description = storedContract.Description
		contract.EthAddress = storedContract.EthAddress
		contract.Network = storedContract.Network
		contract.AdminID = storedContract.AdminID
		creationDate, err := time.Parse(time.RFC3339, storedContract.CreationDate)
		if err != nil {
			return nil, err
		}
		contract.CreationDate = creationDate
		lastUpdate, err := time.Parse(time.RFC3339, storedContract.LastUpdate)
		if err != nil {
			return nil, err
		}
		contract.LastUpdate = lastUpdate
		items = append(items, contract)
	}

	return &contract.ContractStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *ContractRepository) Add(ctx context.Context, input contract.ContractStorageData) (*contract.ContractStorageData, error) {
	var contractDb ContractDb
	contractDb.Project = input.Project
	contractDb.ID = input.ID
	contractDb.Name = input.Name
	contractDb.Description = input.Description
	contractDb.EthAddress = input.EthAddress
	contractDb.Network = input.Network
	contractDb.AdminID = input.AdminID
	contractDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	contractDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	insertQuery := "INSERT INTO contracts (project, id, name, description, eth_address, network, admin_id, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(insertQuery, contractDb.Project, contractDb.ID, contractDb.Name, contractDb.Description, contractDb.EthAddress, contractDb.Network, contractDb.AdminID, contractDb.CreationDate, contractDb.LastUpdate)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM contracts WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, contractDb.Project, contractDb.ID)

	var storedContract ContractDb
	err = result.Scan(&storedContract.Project, &storedContract.ID, &storedContract.Name, &storedContract.Description, &storedContract.EthAddress, &storedContract.Network, &storedContract.AdminID, &storedContract.CreationDate, &storedContract.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("contract with project [" + contractDb.Project + "] and id [" + contractDb.ID + "] not added")
		}
		return nil, err
	}

	var output contract.ContractStorageData
	output.Project = storedContract.Project
	output.ID = storedContract.ID
	output.Name = storedContract.Name
	output.Description = storedContract.Description
	output.EthAddress = storedContract.EthAddress
	output.Network = storedContract.Network
	output.AdminID = storedContract.AdminID
	creationDate, err := time.Parse(time.RFC3339, storedContract.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedContract.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func BuildQueryWithFilters(query string, filters contract.Filters) string {

	return query
}
