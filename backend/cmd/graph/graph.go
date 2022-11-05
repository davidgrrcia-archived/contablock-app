package graph

import (
	"database/sql"
	"github.com/gorilla/mux"

	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/entry"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
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
}

func NewGraph() Graph {
	return Graph{}
}

func (graph *Graph) Build() {
	// output infra
	graph.provideEntryDB()

	// output adapters
	graph.provideEntryStorage()
	//graph.provideHoldContractInputAdapter()

	// use cases
	graph.provideEntryUseCase()

	// input adapters
	graph.provideEntryInputAdapter()

	// input infra
	graph.provideRouter()
	graph.provideEntryEndpoints()
}
