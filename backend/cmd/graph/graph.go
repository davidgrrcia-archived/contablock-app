package graph

import (
	"database/sql"
	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gorilla/mux"
)

type Graph struct {
	*mux.Router

	//contablock
	ContablockClient contablockClient.Contablock
	ContablockAuth   bind.TransactOpts
	contablock.ContaBlockContractAdapter
	contablock.ContaBlockUseCase

	//entries
	EntryDb sql.DB
	entry.EntryStorage
	entry.EntryUseCase
	inputEntry.EntryInputAdapter

	//accounts
	AccountDb sql.DB
	account.AccountStorage
	account.AccountUseCase
	inputAccount.AccountInputAdapter

	//business
	BusinessDb sql.DB
	business.BusinessStorage
	business.BusinessUseCase
	inputBusiness.BusinessInputAdapter
}

func NewGraph() Graph {
	return Graph{}
}

func (graph *Graph) Build() {
	// output infra
	graph.provideContablockClient()
	graph.provideAccountDB()
	graph.provideBusinessDB()
	graph.provideEntryDB()

	// output adapters
	graph.provideContablockContractAdapter()
	graph.provideAccountStorage()
	graph.provideBusinessStorage()
	graph.provideEntryStorage()

	// use cases
	graph.provideContablockUseCase()
	graph.provideAccountUseCase()
	graph.provideBusinessUseCase()
	graph.provideEntryUseCase()

	// input adapters
	graph.provideAccountInputAdapter()
	graph.provideBusinessInputAdapter()
	graph.provideEntryInputAdapter()

	// input infra
	graph.provideRouter()
	graph.provideAccountEndpoints()
	graph.provideBusinessEndpoints()
	graph.provideEntryEndpoints()
}
