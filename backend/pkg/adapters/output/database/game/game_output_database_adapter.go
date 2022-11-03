package gameOutputDbAdapter

import (
	"context"
	"errors"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/game"
)

/* AUX VAR */

var _ game.GameStorage = (*InMemoryGameStorage)(nil)

/* INSTANCE TYPES */

type InMemoryGameStorageOptions struct {
	GameDatabase game.GameCollection
}

type InMemoryGameStorage struct {
	gameDatabase game.GameCollection
}

/* INSTANCE CONSTRUCTOR */

func NewInMemoryGameStorage(opts InMemoryGameStorageOptions) (*InMemoryGameStorage, error) {
	if opts.GameDatabase == nil {
		return nil, errors.New("cannot create an instance of InMemoryGameStorage because provided GameDatabase is nil")
	}

	return &InMemoryGameStorage{
		gameDatabase: opts.GameDatabase,
	}, nil
}

func (adapter *InMemoryGameStorage) Add(ctx context.Context, input game.AddInput) (*game.AddOutput, error) {

	//map use case to storage
	var u game.Game
	u.ID = input.ID
	u.Name = input.Name
	u.OwnerUserID = input.OwnerUserID
	u.CreationDate = input.CreationDate
	u.LastUpdate = input.LastUpdate

	adapter.gameDatabase = append(adapter.gameDatabase, u)
	return &game.AddOutput{
		Game: u,
	}, nil
}

func (adapter *InMemoryGameStorage) Get(ctx context.Context, input game.GetInput) (*game.GetOutput, error) {

	for _, u := range adapter.gameDatabase {
		if u.ID != input.ID {
			continue
		}

		return &game.GetOutput{
			Game: u,
		}, nil
	}

	return nil, errors.New("game not found")
}

func (adapter *InMemoryGameStorage) All(ctx context.Context, input game.AllInput) (*game.AllOutput, error) {
	return &game.AllOutput{
		Items: adapter.gameDatabase,
	}, nil
}

func (adapter *InMemoryGameStorage) Edit(ctx context.Context, input game.EditInput) (*game.EditOutput, error) {

	for _, u := range adapter.gameDatabase {
		if u.ID == input.ID {
			//map use case to storage
			u.Name = input.Name
			u.OwnerUserID = input.OwnerUserID
			u.LastUpdate = input.LastUpdate

			return &game.EditOutput{
				Game: u,
			}, nil
		}
	}

	return nil, errors.New("game not found")
}

func (adapter *InMemoryGameStorage) Remove(ctx context.Context, input game.RemoveInput) (*game.RemoveOutput, error) {
	for i, u := range adapter.gameDatabase {
		if u.ID != input.ID {
			continue
		}

		adapter.gameDatabase = append(adapter.gameDatabase[:i], adapter.gameDatabase[i+1:]...)

		return &game.RemoveOutput{
			Game: u,
		}, nil
	}

	return nil, errors.New("game not found")
}
