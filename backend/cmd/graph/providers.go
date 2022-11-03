package graph

import (
	"log"
	"time"

	inputEntry "github.com/davidevloper/HoldContract-go/pkg/adapters/input/http/entry"
	entryOutputDbAdapter "github.com/davidevloper/HoldContract-go/pkg/adapters/output/database/entry"
	"github.com/davidevloper/HoldContract-go/pkg/infra/input"
	"github.com/davidevloper/HoldContract-go/pkg/usecases/entry"
	"github.com/gorilla/mux"
)

func (graph *Graph) provideRouter() {
	graph.Router = mux.NewRouter().StrictSlash(true)
}

func (graph *Graph) provideEntryDB() {
	graph.EntryDb = entry.EntryCollection{
		{
			ID:           "1",
			Email:        "test@test.com",
			Password:     "123",
			CreationDate: time.Now(),
			LastUpdate:   time.Now(),
		},
		{
			ID:           "2",
			Email:        "test@test.com",
			Password:     "123",
			CreationDate: time.Now(),
			LastUpdate:   time.Now(),
		},
	}
}

func (graph *Graph) provideEntryStorage() entry.EntryStorage {
	entryStorageOpts := entryOutputDbAdapter.InMemoryEntryStorageOptions{
		EntryDatabase: graph.GetEntryDB(),
	}
	entryStorage, err := entryOutputDbAdapter.NewInMemoryEntryStorage(entryStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryStorage = entryStorage

	return entryStorage
}

func (graph *Graph) provideEntryUseCase() {
	entryUseCaseOpts := entry.DefaultEntryUseCaseOptions{
		EntryStorage: graph.GetEntryStorage(),
	}
	entryUseCase, err := entry.NewDefaultEntryUseCase(entryUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryUseCase = entryUseCase
}

func (graph *Graph) provideEntryInputAdapter() *inputEntry.DefaultHttpInputAdapter {
	entryInputAdapterOpts := inputEntry.DefaultHttpInputAdapterOptions{
		EntryUseCase: graph.GetEntryUseCase(),
	}
	entryInputAdapter, err := inputEntry.NewDefaultHttpInputAdapter(entryInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryInputAdapter = entryInputAdapter

	return entryInputAdapter
}

func (graph *Graph) provideEntryEndpoints() {
	var options input.DefaultEntriesHttpHandlerOptions
	options.Adapter = graph.GetEntryInputAdapter()
	handler, err := input.NewDefaultEntriesHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterEntryRoutes(graph.GetRouter())
}
