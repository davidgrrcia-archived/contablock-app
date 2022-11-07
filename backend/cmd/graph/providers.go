package graph

import (
	"database/sql"
	"log"

	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/account"
	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/input/http/entry"
	accountOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/output/database/account"
	entryOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/adapters/output/database/entry"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/infra/input"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/account"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/usecases/entry"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func (graph *Graph) provideRouter() {
	graph.Router = mux.NewRouter().StrictSlash(true)
}

func (graph *Graph) provideEntryDB() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_accounting")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.EntryDb = *db
}

func (graph *Graph) provideEntryStorage() entry.EntryStorage {
	entryStorageOpts := entryOutputDbAdapter.EntryRepositoryOptions{
		DB: graph.GetEntryDB(),
	}
	entryStorage, err := entryOutputDbAdapter.NewEntryRepository(entryStorageOpts)
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

func (graph *Graph) provideAccountDB() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_accounting")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.AccountDb = *db
}

func (graph *Graph) provideAccountStorage() account.AccountStorage {
	accountStorageOpts := accountOutputDbAdapter.AccountRepositoryOptions{
		DB: graph.GetAccountDB(),
	}
	accountStorage, err := accountOutputDbAdapter.NewAccountRepository(accountStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.AccountStorage = accountStorage

	return accountStorage
}

func (graph *Graph) provideAccountUseCase() {
	accountUseCaseOpts := account.DefaultAccountUseCaseOptions{
		AccountStorage: graph.GetAccountStorage(),
	}
	accountUseCase, err := account.NewDefaultAccountUseCase(accountUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.AccountUseCase = accountUseCase
}

func (graph *Graph) provideAccountInputAdapter() *inputAccount.DefaultHttpInputAdapter {
	accountInputAdapterOpts := inputAccount.DefaultHttpInputAdapterOptions{
		AccountUseCase: graph.GetAccountUseCase(),
	}
	accountInputAdapter, err := inputAccount.NewDefaultHttpInputAdapter(accountInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.AccountInputAdapter = accountInputAdapter

	return accountInputAdapter
}

func (graph *Graph) provideAccountEndpoints() {
	var options input.DefaultAccountsHttpHandlerOptions
	options.Adapter = graph.GetAccountInputAdapter()
	handler, err := input.NewDefaultAccountsHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterAccountRoutes(graph.GetRouter())
}
