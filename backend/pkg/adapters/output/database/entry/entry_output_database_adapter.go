package entryOutputDbAdapter

import (
	"context"
	"errors"

	"github.com/davidevloper/HoldContract-go/pkg/usecases/entry"
)

func NewInMemoryEntryStorage(opts InMemoryEntryStorageOptions) (*InMemoryEntryStorage, error) {
	if opts.EntryDatabase == nil {
		return nil, errors.New("cannot create an instance of InMemoryEntryStorage because provided EntryDatabase is nil")
	}

	return &InMemoryEntryStorage{
		entryDatabase: opts.EntryDatabase,
	}, nil
}

type InMemoryEntryStorageOptions struct {
	EntryDatabase entry.EntryCollection
}

type InMemoryEntryStorage struct {
	entryDatabase entry.EntryCollection
}

var _ entry.EntryStorage = (*InMemoryEntryStorage)(nil)

func (adapter *InMemoryEntryStorage) Add(ctx context.Context, input entry.AddInput) (*entry.AddOutput, error) {
	var u entry.Entry
	u.ID = input.ID
	u.Email = input.Email
	u.Password = input.Password
	u.CreationDate = input.CreationDate
	u.LastUpdate = input.LastUpdate

	adapter.entryDatabase = append(adapter.entryDatabase, u)
	return &entry.AddOutput{
		Entry: u,
	}, nil
}

func (adapter *InMemoryEntryStorage) Get(ctx context.Context, input entry.GetInput) (*entry.GetOutput, error) {
	for _, u := range adapter.entryDatabase {
		if u.ID != input.ID {
			continue
		}

		return &entry.GetOutput{
			Entry: u,
		}, nil
	}

	return nil, errors.New("entry not found")
}

func (adapter *InMemoryEntryStorage) All(ctx context.Context, input entry.AllInput) (*entry.AllOutput, error) {
	return &entry.AllOutput{
		Items: adapter.entryDatabase,
	}, nil
}

func (adapter *InMemoryEntryStorage) Edit(ctx context.Context, input entry.EditInput) (*entry.EditOutput, error) {

	for _, u := range adapter.entryDatabase {
		if u.ID == input.ID {
			u.Email = input.Email
			u.Password = input.Password
			u.LastUpdate = input.LastUpdate

			return &entry.EditOutput{
				Entry: u,
			}, nil
		}
	}

	return nil, errors.New("entry not found")
}

func (adapter *InMemoryEntryStorage) Remove(ctx context.Context, input entry.RemoveInput) (*entry.RemoveOutput, error) {
	for i, u := range adapter.entryDatabase {
		if u.ID != input.ID {
			continue
		}

		adapter.entryDatabase = append(adapter.entryDatabase[:i], adapter.entryDatabase[i+1:]...)

		return &entry.RemoveOutput{
			Entry: u,
		}, nil
	}

	return nil, errors.New("entry not found")
}
