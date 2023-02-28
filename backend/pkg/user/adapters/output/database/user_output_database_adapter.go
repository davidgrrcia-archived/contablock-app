package userdb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
)

func NewUserRepository(opts UserRepositoryOptions) (*UserRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of UserRepository with a nil sql.DB")
	}

	return &UserRepository{
		db: opts.DB,
	}, nil
}

type UserRepositoryOptions struct {
	DB *sql.DB
}

type UserRepository struct {
	db *sql.DB
}

var _ user.UserStorage = (*UserRepository)(nil)

func (repository *UserRepository) Get(ctx context.Context, project string, id string) (*user.UserStorageData, error) {
	query := "SELECT * FROM users WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedUser UserDb
	err := result.Scan(&storedUser.Project, &storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password, &storedUser.Role, &storedUser.EthAddress, &storedUser.CreationDate, &storedUser.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var output user.UserStorageData
	output.Project = storedUser.Project
	output.ID = storedUser.ID
	output.Name = storedUser.Name
	output.Email = storedUser.Email
	output.Password = storedUser.Password
	output.Role = storedUser.Role
	output.EthAddress = storedUser.EthAddress
	creationDate, err := time.Parse(time.RFC3339, storedUser.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedUser.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *UserRepository) All(ctx context.Context, project string, filters user.Filters) (*user.UserStorageCollectionData, error) {
	query := BuildQueryWithFilters("SELECT * FROM users WHERE project = ?", filters)

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedUsers []UserDb
	for rows.Next() {
		var userDb UserDb
		err = rows.Scan(&userDb.Project, &userDb.ID, &userDb.Name, &userDb.Email, &userDb.Password, &userDb.Role, &userDb.EthAddress, &userDb.CreationDate, &userDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		storedUsers = append(storedUsers, userDb)
	}

	var items []user.UserStorageData
	for _, storedUser := range storedUsers {
		var user user.UserStorageData
		user.Project = storedUser.Project
		user.ID = storedUser.ID
		user.Name = storedUser.Name
		user.Email = storedUser.Email
		user.Password = storedUser.Password
		user.Role = storedUser.Role
		user.EthAddress = storedUser.EthAddress
		creationDate, err := time.Parse(time.RFC3339, storedUser.CreationDate)
		if err != nil {
			return nil, err
		}
		user.CreationDate = creationDate
		lastUpdate, err := time.Parse(time.RFC3339, storedUser.LastUpdate)
		if err != nil {
			return nil, err
		}
		user.LastUpdate = lastUpdate
		items = append(items, user)
	}

	return &user.UserStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *UserRepository) Edit(ctx context.Context, input user.UserStorageData) (*user.UserStorageData, error) {
	var userDb UserDb
	userDb.Project = input.Project
	userDb.ID = input.ID
	userDb.Name = input.Name
	userDb.Email = input.Email
	userDb.Password = input.Password
	userDb.Role = input.Role
	userDb.EthAddress = input.EthAddress
	userDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	userDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	updateQuery := "UPDATE users SET name = ?, email = ?, password = ?, role = ?, eth_address = ?, creation_date = ?, last_update = ? WHERE project = ? AND id = ?"

	repository.db.Exec(updateQuery, userDb.Name, userDb.Email, userDb.Password, userDb.Role, userDb.EthAddress, userDb.CreationDate, userDb.LastUpdate, userDb.Project, userDb.ID)

	query := "SELECT * FROM users WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, userDb.Project, userDb.ID)

	var storedUser UserDb
	err := result.Scan(&storedUser.Project, &storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password, &storedUser.Role, &storedUser.EthAddress, &storedUser.CreationDate, &storedUser.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user with project [" + userDb.Project + "] and id [" + userDb.ID + "] not found")
		}
		return nil, err
	}

	var output user.UserStorageData
	output.Project = storedUser.Project
	output.ID = storedUser.ID
	output.Name = storedUser.Name
	output.Email = storedUser.Email
	output.Password = storedUser.Password
	output.Role = storedUser.Role
	output.EthAddress = storedUser.EthAddress
	creationDate, err := time.Parse(time.RFC3339, storedUser.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedUser.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *UserRepository) Add(ctx context.Context, input user.UserStorageData) (*user.UserStorageData, error) {
	var userDb UserDb
	userDb.Project = input.Project
	userDb.ID = input.ID
	userDb.Name = input.Name
	userDb.Email = input.Email
	userDb.Password = input.Password
	userDb.Role = input.Role
	userDb.EthAddress = input.EthAddress
	userDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	userDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	insertQuery := "INSERT INTO users (project, id, name, email, password, role, eth_address, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(insertQuery, userDb.Project, userDb.ID, userDb.Name, userDb.Email, userDb.Password, userDb.Role, userDb.EthAddress, userDb.CreationDate, userDb.LastUpdate)
	if err != nil {
		return nil, err
	}

	query := "SELECT * FROM users WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, userDb.Project, userDb.ID)

	var storedUser UserDb
	err = result.Scan(&storedUser.Project, &storedUser.ID, &storedUser.Name, &storedUser.Email, &storedUser.Password, &storedUser.Role, &storedUser.EthAddress, &storedUser.CreationDate, &storedUser.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user with project [" + userDb.Project + "] and id [" + userDb.ID + "] not added")
		}
		return nil, err
	}

	var output user.UserStorageData
	output.Project = storedUser.Project
	output.ID = storedUser.ID
	output.Name = storedUser.Name
	output.Email = storedUser.Email
	output.Password = storedUser.Password
	output.Role = storedUser.Role
	output.EthAddress = storedUser.EthAddress
	creationDate, err := time.Parse(time.RFC3339, storedUser.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedUser.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func BuildQueryWithFilters(query string, filters user.Filters) string {
	if filters.Email != nil {
		query = query + " AND email = \"" + *filters.Email + "\""
		return query
	}
	if filters.Password != nil {
		query = query + " AND password = \"" + *filters.Password + "\""
		return query
	}

	return query
}
