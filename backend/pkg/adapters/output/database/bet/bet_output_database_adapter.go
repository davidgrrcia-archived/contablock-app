package betOutputDbAdapter

import (
	"context"
	"errors"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
)

/* AUX VAR */

var _ bet.BetStorage = (*InMemoryBetStorage)(nil)

/* INSTANCE TYPES */

type InMemoryBetStorageOptions struct {
	BetDatabase bet.BetCollection
}

type InMemoryBetStorage struct {
	betDatabase bet.BetCollection
}

/* INSTANCE CONSTRUCTOR */

func NewInMemoryBetStorage(opts InMemoryBetStorageOptions) (*InMemoryBetStorage, error) {
	if opts.BetDatabase == nil {
		return nil, errors.New("cannot create an instance of InMemoryBetStorage because provided BetDatabase is nil")
	}

	return &InMemoryBetStorage{
		betDatabase: opts.BetDatabase,
	}, nil
}

func (adapter *InMemoryBetStorage) Add(ctx context.Context, input bet.AddInput) (*bet.AddOutput, error) {

	//map use case to storage
	var u bet.Bet
	u.ID = input.ID
	u.Amount = input.Amount
	u.GameID = input.GameID
	u.UserID = input.UserID
	u.CreationDate = input.CreationDate
	u.LastUpdate = input.LastUpdate

	adapter.betDatabase = append(adapter.betDatabase, u)
	return &bet.AddOutput{
		Bet: u,
	}, nil
}

func (adapter *InMemoryBetStorage) Get(ctx context.Context, input bet.GetInput) (*bet.GetOutput, error) {

	for _, u := range adapter.betDatabase {
		if u.ID != input.ID && u.UserID != input.ID {
			continue
		}

		return &bet.GetOutput{
			Bet: u,
		}, nil
	}

	return nil, errors.New("bet not found")
}

func (adapter *InMemoryBetStorage) GetByUserID(ctx context.Context, input bet.GetByUserIDInput) (*bet.GetByUserIDOutput, error) {

	for _, u := range adapter.betDatabase {
		if u.UserID != input.UserID && u.GameID != input.GameID {
			continue
		}

		return &bet.GetByUserIDOutput{
			Bet: u,
		}, nil
	}

	return nil, errors.New("bet not found")
}

func (adapter *InMemoryBetStorage) All(ctx context.Context, input bet.AllInput) (*bet.AllOutput, error) {
	return &bet.AllOutput{
		Items: adapter.betDatabase,
	}, nil
}

func (adapter *InMemoryBetStorage) Edit(ctx context.Context, input bet.EditInput) (*bet.EditOutput, error) {

	for _, u := range adapter.betDatabase {
		if u.ID == input.ID {
			//map use case to storage
			u.Amount = input.Amount
			u.GameID = input.GameID
			u.UserID = input.UserID
			u.LastUpdate = input.LastUpdate

			return &bet.EditOutput{
				Bet: u,
			}, nil
		}
	}

	return nil, errors.New("bet not found")
}

func (adapter *InMemoryBetStorage) Remove(ctx context.Context, input bet.RemoveInput) (*bet.RemoveOutput, error) {
	for i, u := range adapter.betDatabase {
		if u.ID != input.ID {
			continue
		}

		adapter.betDatabase = append(adapter.betDatabase[:i], adapter.betDatabase[i+1:]...)

		return &bet.RemoveOutput{
			Bet: u,
		}, nil
	}

	return nil, errors.New("bet not found")
}
