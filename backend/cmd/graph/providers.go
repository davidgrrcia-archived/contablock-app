package graph

import (
	"database/sql"
	"log"

	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	accountOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/output/database"
	accountInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/infra/input"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	businessOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/output/database"
	businessInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/infra/input"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	entryOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/output/database"
	entryInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/infra/input"
	"github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func (graph *Graph) provideRouter() {
	graph.Router = mux.NewRouter().StrictSlash(true)
}

/////// ENTRIES ///////

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
		EntryStorage:   graph.GetEntryStorage(),
		AccountUseCase: graph.GetAccountUseCase(),
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
	var options entryInputInfra.DefaultEntriesHttpHandlerOptions
	options.Adapter = graph.GetEntryInputAdapter()
	handler, err := entryInputInfra.NewDefaultEntriesHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterEntryRoutes(graph.GetRouter())
}

/////// ACCOUNTS ///////

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
	var options accountInputInfra.DefaultAccountsHttpHandlerOptions
	options.Adapter = graph.GetAccountInputAdapter()
	handler, err := accountInputInfra.NewDefaultAccountsHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterAccountRoutes(graph.GetRouter())
}

/////// BUSINESS ///////

func (graph *Graph) provideBusinessDB() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/db_accounting")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.BusinessDb = *db
}

func (graph *Graph) provideBusinessStorage() business.BusinessStorage {
	businessStorageOpts := businessOutputDbAdapter.BusinessRepositoryOptions{
		DB: graph.GetBusinessDB(),
	}
	businessStorage, err := businessOutputDbAdapter.NewBusinessRepository(businessStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.BusinessStorage = businessStorage

	return businessStorage
}

func (graph *Graph) provideBusinessUseCase() {
	businessUseCaseOpts := business.DefaultBusinessUseCaseOptions{
		BusinessStorage: graph.GetBusinessStorage(),
	}
	businessUseCase, err := business.NewDefaultBusinessUseCase(businessUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.BusinessUseCase = businessUseCase
}

func (graph *Graph) provideBusinessInputAdapter() *inputBusiness.DefaultHttpInputAdapter {
	businessInputAdapterOpts := inputBusiness.DefaultHttpInputAdapterOptions{
		BusinessUseCase: graph.GetBusinessUseCase(),
	}
	businessInputAdapter, err := inputBusiness.NewDefaultHttpInputAdapter(businessInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.BusinessInputAdapter = businessInputAdapter

	return businessInputAdapter
}

func (graph *Graph) provideBusinessEndpoints() {
	var options businessInputInfra.DefaultBusinessHttpHandlerOptions
	options.Adapter = graph.GetBusinessInputAdapter()
	handler, err := businessInputInfra.NewDefaultBusinessHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterBusinessRoutes(graph.GetRouter())
}
