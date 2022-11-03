package graph

import (
	"github.com/gorilla/mux"

	inputEntry "github.com/davidevloper/HoldContract-go/pkg/adapters/input/http/entry"
	"github.com/davidevloper/HoldContract-go/pkg/usecases/entry"
)

type Graph struct {
	*mux.Router

	//holdContract
	//holdContractAdapter.HoldContractAdapter
	//holdContract.HoldContractInterface

	//entries
	EntryDb entry.EntryCollection
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
