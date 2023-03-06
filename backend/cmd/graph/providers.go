package graph

import (
	"database/sql"
	contablockOutputHttpAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/adapters/output/http"
	"log"
	"os"

	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	accountOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/output/database"
	accountInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/infra/input"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputBalance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/adapters/input/http"
	balanceOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/adapters/output/database"
	balanceInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/infra/input"
	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
	inputBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	businessOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/output/database"
	businessInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/infra/input"
	business "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputContract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/adapters/input/http"
	contractOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/adapters/output/database"
	contractInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/infra/input"
	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
	inputDailyLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/adapters/input/http"
	dailyLedgerInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/infra/input"
	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
	inputAsyncEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/async"
	inputHttpEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	entryOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/output/database"
	entryInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/infra/input"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	inputEventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/adapters/input/http"
	eventSubscriptionOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/adapters/output/database"
	eventSubscriptionInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/infra/input"
	eventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/usecases"
	inputGeneralLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/adapters/input/http"
	generalLedgerInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/infra/input"
	generalledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	inputUser "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/adapters/input/http"
	userOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/adapters/output/database"
	userInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/infra/input"
	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
	"github.com/ethereum/go-ethereum/ethclient"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func (graph *Graph) provideRouter() {
	graph.Router = mux.NewRouter().StrictSlash(true)
}

/////// ENTRIES ///////

func (graph *Graph) provideEntryDB() {

	println(os.Getenv("DB_USER"))

	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
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
		EntryStorage:      graph.GetEntryStorage(),
		ContaBlockUseCase: graph.GetContaBlockUseCase(),
		BalanceUseCase:    graph.GetBalanceUseCase(),
		AccountUseCase:    graph.GetAccountUseCase(),
		ContractUseCase:   graph.GetContractUseCase(),
		UserUseCase:       graph.GetUserUseCase(),
	}
	entryUseCase, err := entry.NewDefaultEntryUseCase(entryUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryUseCase = entryUseCase
}

func (graph *Graph) provideEntryInputHttpAdapter() *inputHttpEntry.DefaultHttpInputAdapter {
	entryInputHttpAdapterOpts := inputHttpEntry.DefaultHttpInputAdapterOptions{
		EntryUseCase: graph.GetEntryUseCase(),
	}
	entryInputHttpAdapter, err := inputHttpEntry.NewDefaultHttpInputAdapter(entryInputHttpAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryInputHttpAdapter = entryInputHttpAdapter

	return entryInputHttpAdapter
}

func (graph *Graph) provideEntryInputAsyncAdapter() *inputAsyncEntry.DefaultAsyncInputAdapter {
	entryInputAsyncAdapterOpts := inputAsyncEntry.DefaultAsyncInputAdapterOptions{
		EntryUseCase: graph.GetEntryUseCase(),
	}
	entryInputAsyncAdapter, err := inputAsyncEntry.NewDefaultAsyncInputAdapter(entryInputAsyncAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.EntryInputAsyncAdapter = entryInputAsyncAdapter

	return entryInputAsyncAdapter
}

func (graph *Graph) provideEntryEndpoints() {
	var options entryInputInfra.DefaultEntriesHttpHandlerOptions
	options.Adapter = graph.GetEntryInputHttpAdapter()
	handler, err := entryInputInfra.NewDefaultEntriesHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterEntryRoutes(graph.GetRouter())
}

/////// ACCOUNTS ///////

func (graph *Graph) provideAccountDB() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
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
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
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

/////// SMART CONTRACT CONTABLOCK ///////

func (graph *Graph) provideContablockClient() {
	client, err := ethclient.Dial(os.Getenv("SC_CLIENT_NETWORK"))
	if err != nil {
		panic("could not connect to Ethereum gateway")
	}

	graph.ContablockClient = client
}

func (graph *Graph) provideContablockContractAdapter() contablock.ContaBlockContractAdapter {
	contaBlockContractAdapterOpts := contablockOutputHttpAdapter.DefaultContaBlockContractOptions{
		Client: graph.GetContablockClient(),
	}
	contaBlockContractAdapter, err := contablockOutputHttpAdapter.NewDefaultContaBlockContract(contaBlockContractAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.ContaBlockContractAdapter = contaBlockContractAdapter

	return contaBlockContractAdapter
}

func (graph *Graph) provideContablockUseCase() {
	contablockUseCaseOpts := contablockUseCase.DefaultContaBlockUseCaseOptions{
		ContaBlockContractAdapter: graph.GetContablockAdapter(),
	}
	contablockUseCase, err := contablockUseCase.NewDefaultContaBlockUseCase(contablockUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.ContaBlockUseCase = contablockUseCase
}

/////// DAILY LEDGER ///////

func (graph *Graph) provideDailyLedgerUseCase() {
	dailyledgerUseCaseOpts := dailyledger.DefaultDailyLedgerUseCaseOptions{
		EntryUseCase:      graph.GetEntryUseCase(),
		AccountUseCase:    graph.GetAccountUseCase(),
		ContaBlockUseCase: graph.GetContaBlockUseCase(),
		ContractUseCase:   graph.GetContractUseCase(),
		UserUseCase:       graph.GetUserUseCase(),
	}
	dailyledgerUseCase, err := dailyledger.NewDefaultDailyLedgerUseCase(dailyledgerUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.DailyLedgerUseCase = dailyledgerUseCase
}

func (graph *Graph) provideDailyLedgerInputAdapter() *inputDailyLedger.DefaultHttpInputAdapter {
	dailyledgerInputAdapterOpts := inputDailyLedger.DefaultHttpInputAdapterOptions{
		DailyLedgerUseCase: graph.GetDailyLedgerUseCase(),
	}
	dailyledgerInputAdapter, err := inputDailyLedger.NewDefaultHttpInputAdapter(dailyledgerInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.DailyLedgerInputAdapter = dailyledgerInputAdapter

	return dailyledgerInputAdapter
}

func (graph *Graph) provideDailyLedgerEndpoints() {
	var options dailyLedgerInputInfra.DefaultDailyLedgerHttpHandlerOptions
	options.Adapter = graph.GetDailyLedgerInputAdapter()
	handler, err := dailyLedgerInputInfra.NewDefaultDailyLedgerHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterDailyLedgerRoutes(graph.GetRouter())
}

/////// GENERAL LEDGER ///////

func (graph *Graph) provideGeneralLedgerUseCase() {
	generalledgerUseCaseOpts := generalledger.DefaultGeneralLedgerUseCaseOptions{
		EntryUseCase:      graph.GetEntryUseCase(),
		AccountUseCase:    graph.GetAccountUseCase(),
		ContaBlockUseCase: graph.GetContaBlockUseCase(),
		ContractUseCase:   graph.GetContractUseCase(),
		UserUseCase:       graph.GetUserUseCase(),
	}
	generalledgerUseCase, err := generalledger.NewDefaultGeneralLedgerUseCase(generalledgerUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.GeneralLedgerUseCase = generalledgerUseCase
}

func (graph *Graph) provideGeneralLedgerInputAdapter() *inputGeneralLedger.DefaultHttpInputAdapter {
	generalledgerInputAdapterOpts := inputGeneralLedger.DefaultHttpInputAdapterOptions{
		GeneralLedgerUseCase: graph.GetGeneralLedgerUseCase(),
	}
	generalledgerInputAdapter, err := inputGeneralLedger.NewDefaultHttpInputAdapter(generalledgerInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.GeneralLedgerInputAdapter = generalledgerInputAdapter

	return generalledgerInputAdapter
}

func (graph *Graph) provideGeneralLedgerEndpoints() {
	var options generalLedgerInputInfra.DefaultGeneralLedgerHttpHandlerOptions
	options.Adapter = graph.GetGeneralLedgerInputAdapter()
	handler, err := generalLedgerInputInfra.NewDefaultGeneralLedgerHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterGeneralLedgerRoutes(graph.GetRouter())
}

/////// BALANCES ///////

func (graph *Graph) provideBalanceDB() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.BalanceDb = *db
}

func (graph *Graph) provideBalanceStorage() balance.BalanceStorage {
	balanceStorageOpts := balanceOutputDbAdapter.BalanceRepositoryOptions{
		DB: graph.GetBalanceDB(),
	}
	balanceStorage, err := balanceOutputDbAdapter.NewBalanceRepository(balanceStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.BalanceStorage = balanceStorage

	return balanceStorage
}

func (graph *Graph) provideBalanceUseCase() {
	balanceUseCaseOpts := balance.DefaultBalanceUseCaseOptions{
		BalanceStorage: graph.GetBalanceStorage(),
	}
	balanceUseCase, err := balance.NewDefaultBalanceUseCase(balanceUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.BalanceUseCase = balanceUseCase
}

func (graph *Graph) provideBalanceInputAdapter() *inputBalance.DefaultHttpInputAdapter {
	balanceInputAdapterOpts := inputBalance.DefaultHttpInputAdapterOptions{
		BalanceUseCase: graph.GetBalanceUseCase(),
	}
	balanceInputAdapter, err := inputBalance.NewDefaultHttpInputAdapter(balanceInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.BalanceInputAdapter = balanceInputAdapter

	return balanceInputAdapter
}

func (graph *Graph) provideBalanceEndpoints() {
	var options balanceInputInfra.DefaultBalancesHttpHandlerOptions
	options.Adapter = graph.GetBalanceInputAdapter()
	handler, err := balanceInputInfra.NewDefaultBalancesHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterBalanceRoutes(graph.GetRouter())
}

/////// USERS ///////

func (graph *Graph) provideUserDB() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.UserDb = *db
}

func (graph *Graph) provideUserStorage() user.UserStorage {
	userStorageOpts := userOutputDbAdapter.UserRepositoryOptions{
		DB: graph.GetUserDB(),
	}
	userStorage, err := userOutputDbAdapter.NewUserRepository(userStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.UserStorage = userStorage

	return userStorage
}

func (graph *Graph) provideUserUseCase() {
	userUseCaseOpts := user.DefaultUserUseCaseOptions{
		UserStorage:       graph.GetUserStorage(),
		ContaBlockUseCase: graph.GetContaBlockUseCase(),
		ContractUseCase:   graph.GetContractUseCase(),
	}
	userUseCase, err := user.NewDefaultUserUseCase(userUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.UserUseCase = userUseCase
}

func (graph *Graph) provideUserInputAdapter() *inputUser.DefaultHttpInputAdapter {
	userInputAdapterOpts := inputUser.DefaultHttpInputAdapterOptions{
		UserUseCase: graph.GetUserUseCase(),
	}
	userInputAdapter, err := inputUser.NewDefaultHttpInputAdapter(userInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.UserInputAdapter = userInputAdapter

	return userInputAdapter
}

func (graph *Graph) provideUserEndpoints() {
	var options userInputInfra.DefaultUsersHttpHandlerOptions
	options.Adapter = graph.GetUserInputAdapter()
	handler, err := userInputInfra.NewDefaultUsersHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterUserRoutes(graph.GetRouter())
}

/////// CONTRACTS ///////

func (graph *Graph) provideContractDB() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.ContractDb = *db
}

func (graph *Graph) provideContractStorage() contract.ContractStorage {
	contractStorageOpts := contractOutputDbAdapter.ContractRepositoryOptions{
		DB: graph.GetContractDB(),
	}
	contractStorage, err := contractOutputDbAdapter.NewContractRepository(contractStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.ContractStorage = contractStorage

	return contractStorage
}

func (graph *Graph) provideContractUseCase() {
	contractUseCaseOpts := contract.DefaultContractUseCaseOptions{
		ContractStorage:   graph.GetContractStorage(),
		ContaBlockUseCase: graph.GetContaBlockUseCase(),
	}
	contractUseCase, err := contract.NewDefaultContractUseCase(contractUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.ContractUseCase = contractUseCase
}

func (graph *Graph) provideContractInputAdapter() *inputContract.DefaultHttpInputAdapter {
	contractInputAdapterOpts := inputContract.DefaultHttpInputAdapterOptions{
		ContractUseCase: graph.GetContractUseCase(),
	}
	contractInputAdapter, err := inputContract.NewDefaultHttpInputAdapter(contractInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.ContractInputAdapter = contractInputAdapter

	return contractInputAdapter
}

func (graph *Graph) provideContractEndpoints() {
	var options contractInputInfra.DefaultContractsHttpHandlerOptions
	options.Adapter = graph.GetContractInputAdapter()
	handler, err := contractInputInfra.NewDefaultContractsHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterContractRoutes(graph.GetRouter())
}

/////// EVENT SUBSCRIPTIONS ///////

func (graph *Graph) provideEventSubscriptionDB() {
	db, err := sql.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME"))
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	graph.EventSubscriptionDb = *db
}

func (graph *Graph) provideEventSubscriptionStorage() eventSubscription.EventSubscriptionStorage {
	eventSubscriptionStorageOpts := eventSubscriptionOutputDbAdapter.EventSubscriptionRepositoryOptions{
		DB: graph.GetEventSubscriptionDB(),
	}
	eventSubscriptionStorage, err := eventSubscriptionOutputDbAdapter.NewEventSubscriptionRepository(eventSubscriptionStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.EventSubscriptionStorage = eventSubscriptionStorage

	return eventSubscriptionStorage
}

func (graph *Graph) provideEventSubscriptionUseCase() {
	eventSubscriptionUseCaseOpts := eventSubscription.DefaultEventSubscriptionUseCaseOptions{
		EventSubscriptionStorage: graph.GetEventSubscriptionStorage(),
	}
	eventSubscriptionUseCase, err := eventSubscription.NewDefaultEventSubscriptionUseCase(eventSubscriptionUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.EventSubscriptionUseCase = eventSubscriptionUseCase
}

func (graph *Graph) provideEventSubscriptionInputAdapter() *inputEventSubscription.DefaultHttpInputAdapter {
	eventSubscriptionInputAdapterOpts := inputEventSubscription.DefaultHttpInputAdapterOptions{
		EventSubscriptionUseCase: graph.GetEventSubscriptionUseCase(),
	}
	eventSubscriptionInputAdapter, err := inputEventSubscription.NewDefaultHttpInputAdapter(eventSubscriptionInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.EventSubscriptionInputAdapter = eventSubscriptionInputAdapter

	return eventSubscriptionInputAdapter
}

func (graph *Graph) provideEventSubscriptionEndpoints() {
	var options eventSubscriptionInputInfra.DefaultEventSubscriptionsHttpHandlerOptions
	options.Adapter = graph.GetEventSubscriptionInputAdapter()
	handler, err := eventSubscriptionInputInfra.NewDefaultEventSubscriptionsHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterEventSubscriptionRoutes(graph.GetRouter())
}
