package eventSubscriptiondb

import (
	"context"
	"database/sql"
	"errors"
	"time"

	eventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/usecases"
)

func NewEventSubscriptionRepository(opts EventSubscriptionRepositoryOptions) (*EventSubscriptionRepository, error) {
	if opts.DB == nil {
		return nil, errors.New("cannot create an instance of EventSubscriptionRepository with a nil sql.DB")
	}

	return &EventSubscriptionRepository{
		db: opts.DB,
	}, nil
}

type EventSubscriptionRepositoryOptions struct {
	DB *sql.DB
}

type EventSubscriptionRepository struct {
	db *sql.DB
}

var _ eventSubscription.EventSubscriptionStorage = (*EventSubscriptionRepository)(nil)

func (repository *EventSubscriptionRepository) Get(ctx context.Context, project string, id string) (*eventSubscription.EventSubscriptionStorageData, error) {
	query := "SELECT * FROM event_subscriptions WHERE project = ? AND id = ?"

	result := repository.db.QueryRow(query, project, id)

	var storedEventSubscription EventSubscriptionDb
	err := result.Scan(&storedEventSubscription.Project, &storedEventSubscription.ID, &storedEventSubscription.ContractID, &storedEventSubscription.Network, &storedEventSubscription.CreationDate, &storedEventSubscription.LastUpdate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	var output eventSubscription.EventSubscriptionStorageData
	output.Project = storedEventSubscription.Project
	output.ID = storedEventSubscription.ID
	output.ContractID = storedEventSubscription.ContractID
	output.Network = storedEventSubscription.Network
	creationDate, err := time.Parse(time.RFC3339, storedEventSubscription.CreationDate)
	if err != nil {
		return nil, err
	}
	output.CreationDate = creationDate
	lastUpdate, err := time.Parse(time.RFC3339, storedEventSubscription.LastUpdate)
	if err != nil {
		return nil, err
	}
	output.LastUpdate = lastUpdate

	return &output, nil
}

func (repository *EventSubscriptionRepository) All(ctx context.Context, project string, filters eventSubscription.Filters) (*eventSubscription.EventSubscriptionStorageCollectionData, error) {
	query := BuildQueryWithFilters("SELECT * FROM event_subscriptions WHERE project = ?", filters)

	rows, err := repository.db.Query(query, project)
	if err != nil {
		return nil, err
	}

	var storedEventSubscriptions []EventSubscriptionDb
	for rows.Next() {
		var eventSubscriptionDb EventSubscriptionDb
		err = rows.Scan(&eventSubscriptionDb.Project, &eventSubscriptionDb.ID, &eventSubscriptionDb.ContractID, &eventSubscriptionDb.Network, &eventSubscriptionDb.CreationDate, &eventSubscriptionDb.LastUpdate)
		if err != nil {
			return nil, err
		}

		storedEventSubscriptions = append(storedEventSubscriptions, eventSubscriptionDb)
	}

	var items []eventSubscription.EventSubscriptionStorageData
	for _, storedEventSubscription := range storedEventSubscriptions {
		var eventSubscription eventSubscription.EventSubscriptionStorageData
		eventSubscription.Project = storedEventSubscription.Project
		eventSubscription.ID = storedEventSubscription.ID
		eventSubscription.ContractID = storedEventSubscription.ContractID
		eventSubscription.Network = storedEventSubscription.Network
		creationDate, err := time.Parse(time.RFC3339, storedEventSubscription.CreationDate)
		if err != nil {
			return nil, err
		}
		eventSubscription.CreationDate = creationDate
		lastUpdate, err := time.Parse(time.RFC3339, storedEventSubscription.LastUpdate)
		if err != nil {
			return nil, err
		}
		eventSubscription.LastUpdate = lastUpdate
		items = append(items, eventSubscription)
	}

	return &eventSubscription.EventSubscriptionStorageCollectionData{
		Items: items,
	}, nil
}

func (repository *EventSubscriptionRepository) Add(ctx context.Context, input eventSubscription.EventSubscriptionStorageData) (*eventSubscription.EventSubscriptionStorageData, error) {
	var eventSubscriptionDb EventSubscriptionDb
	eventSubscriptionDb.Project = input.Project
	eventSubscriptionDb.ID = input.ID
	eventSubscriptionDb.ContractID = input.ContractID
	eventSubscriptionDb.Network = input.Network
	eventSubscriptionDb.CreationDate = input.CreationDate.Format(time.RFC3339)
	eventSubscriptionDb.LastUpdate = input.LastUpdate.Format(time.RFC3339)

	insertQuery := "INSERT INTO event_subscriptions (project, id, contract_id, network, creation_date, last_update) VALUES (?, ?, ?, ?, ?, ?)"

	_, err := repository.db.Exec(insertQuery, eventSubscriptionDb.Project, eventSubscriptionDb.ID, eventSubscriptionDb.ContractID, eventSubscriptionDb.Network, eventSubscriptionDb.CreationDate, eventSubscriptionDb.LastUpdate)
	if err != nil {
		return nil, err
	}

	output, err := repository.Get(ctx, eventSubscriptionDb.Project, eventSubscriptionDb.ID)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func BuildQueryWithFilters(query string, filters eventSubscription.Filters) string {
	return query
}
