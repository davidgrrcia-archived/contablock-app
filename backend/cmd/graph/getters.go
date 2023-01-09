package graph

import (
	"database/sql"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	inputAdapterAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputAdapterBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputAdapterEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/gorilla/mux"
)

func (graph *Graph) GetRouter() *mux.Router {
	if graph.Router == nil {
		panic("graph Router is nil")
	}
	return graph.Router
}

/////// SMART CONTRACT CONTABLOCK ///////

func (graph *Graph) GetContablockClient() *contablockClient.Contablock {
	return &graph.ContablockClient
}

func (graph *Graph) GetContablockAuth() *bind.TransactOpts {
	return &graph.ContablockAuth
}

func (graph *Graph) GetContablockAdapter() contablock.ContaBlockContractAdapter {
	return graph.ContaBlockContractAdapter
}

func (graph *Graph) GetContaBlockUseCase() contablock.ContaBlockUseCase {
	return graph.ContaBlockUseCase
}

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
