package graph

import (
	"database/sql"

	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/account"
	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/entry"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/account"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
	"github.com/gorilla/mux"
)

type Graph struct {
	*mux.Router

	//holdContract
	//holdContractAdapter.HoldContractAdapter
	//holdContract.HoldContractInterface

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
}

func NewGraph() Graph {
	return Graph{}
}

func (graph *Graph) Build() {
	// output infra
	graph.provideEntryDB()
	graph.provideAccountDB()

	// output adapters
	graph.provideEntryStorage()
	graph.provideAccountStorage()
	//graph.provideHoldContractInputAdapter()

	// use cases
	graph.provideEntryUseCase()
	graph.provideAccountUseCase()

	// input adapters
	graph.provideEntryInputAdapter()
	graph.provideAccountInputAdapter()

	// input infra
	graph.provideRouter()
	graph.provideEntryEndpoints()
	graph.provideAccountEndpoints()
}
