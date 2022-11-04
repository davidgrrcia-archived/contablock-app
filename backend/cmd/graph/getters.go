package graph

import (
	"database/sql"
	inputAdapterEntry "github.com/davidevloper/HoldContract-go/pkg/adapters/input/http/entry"
	"github.com/davidevloper/HoldContract-go/pkg/usecases/entry"
	"github.com/gorilla/mux"
)

func (graph *Graph) GetRouter() *mux.Router {
	if graph.Router == nil {
		panic("graph Router is nil")
	}
	return graph.Router
}

/*
func (graph *Graph) GetHoldContract() *holdContract.HoldContract {
	return graph.HoldContract
}*/

/*
func (graph *Graph) GetHoldContractAdapter() holdContractAdapter.HoldContractAdapter {
	return graph.HoldContractAdapter
}*/

/* ENTRIES */

func (graph *Graph) GetEntryInputAdapter() inputAdapterEntry.EntryInputAdapter {
	return graph.EntryInputAdapter
}

func (graph *Graph) GetEntryDB() *sql.DB {
	return &graph.EntryDb
}

func (graph *Graph) GetEntryStorage() entry.EntryStorage {
	if graph.EntryStorage == nil {
		panic("graph EntryWalletStorage is nil")
	}
	return graph.EntryStorage
}

func (graph *Graph) GetEntryUseCase() entry.EntryUseCase {
	if graph.EntryUseCase == nil {
		panic("graph EntryWalletUseCase case is nil")
	}
	return graph.EntryUseCase
}
