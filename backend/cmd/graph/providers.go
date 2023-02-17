package graph

import (
	"context"
	"crypto/ecdsa"
	"database/sql"
	"log"
	"math/big"
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
	inputDailyLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/adapters/input/http"
	dailyLedgerInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/infra/input"
	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
	inputEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	entryOutputDbAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/output/database"
	entryInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/infra/input"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	inputGeneralLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/adapters/input/http"
	generalLedgerInputInfra "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/infra/input"
	generalledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/usecases"
	contablockOutputHttpAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/adapters/outputhttp"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	_, err = client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	address := common.HexToAddress(os.Getenv("SC_CONTABLOCK_CONTRACT_ADDRESS"))
	instance, err := contablockClient.NewContablock(address, client)
	if err != nil {
		panic(err)
	}

	privateKey, err := crypto.HexToECDSA(os.Getenv("SC_PRIVATE_KEY_AUTH"))
	if err != nil {
		panic(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		panic(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		panic(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice

	graph.ContablockClient = *instance
	graph.ContablockAuth = *auth
	graph.ContablockNonce = *auth.Nonce
}

func (graph *Graph) provideContablockContractAdapter() contablock.ContaBlockContractAdapter {
	contaBlockContractAdapterOpts := contablockOutputHttpAdapter.DefaultContaBlockContractOptions{
		Client: graph.GetContablockClient(),
		Auth:   graph.GetContablockAuth(),
		Nonce:  graph.GetContablockNonce(),
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
