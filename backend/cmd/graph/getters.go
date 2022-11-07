package graph

import (
	"database/sql"

	inputAdapterAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/account"
	inputAdapterBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/business"
	inputAdapterEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/entry"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/account"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/business"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
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

/////// ENTRIES ///////

func (graph *Graph) GetEntryInputAdapter() inputAdapterEntry.EntryInputAdapter {
	return graph.EntryInputAdapter
}

func (graph *Graph) GetEntryDB() *sql.DB {
	return &graph.EntryDb
}

func (graph *Graph) GetEntryStorage() entry.EntryStorage {
	if graph.EntryStorage == nil {
		panic("graph EntryStorage is nil")
	}
	return graph.EntryStorage
}

func (graph *Graph) GetEntryUseCase() entry.EntryUseCase {
	if graph.EntryUseCase == nil {
		panic("graph EntryUseCase case is nil")
	}
	return graph.EntryUseCase
}

/////// ACCOUNTS ///////

func (graph *Graph) GetAccountInputAdapter() inputAdapterAccount.AccountInputAdapter {
	return graph.AccountInputAdapter
}

func (graph *Graph) GetAccountDB() *sql.DB {
	return &graph.AccountDb
}

func (graph *Graph) GetAccountStorage() account.AccountStorage {
	if graph.AccountStorage == nil {
		panic("graph AccountStorage is nil")
	}
	return graph.AccountStorage
}

func (graph *Graph) GetAccountUseCase() account.AccountUseCase {
	if graph.AccountUseCase == nil {
		panic("graph AccountUseCase case is nil")
	}
	return graph.AccountUseCase
}

/////// BUSINESS ///////

func (graph *Graph) GetBusinessInputAdapter() inputAdapterBusiness.BusinessInputAdapter {
	return graph.BusinessInputAdapter
}

func (graph *Graph) GetBusinessDB() *sql.DB {
	return &graph.BusinessDb
}

func (graph *Graph) GetBusinessStorage() business.BusinessStorage {
	if graph.BusinessStorage == nil {
		panic("graph BusinessStorage is nil")
	}
	return graph.BusinessStorage
}

func (graph *Graph) GetBusinessUseCase() business.BusinessUseCase {
	if graph.BusinessUseCase == nil {
		panic("graph BusinessUseCase case is nil")
	}
	return graph.BusinessUseCase
}
