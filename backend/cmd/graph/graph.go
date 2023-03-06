package graph

import (
	"database/sql"
	inputAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputBalance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/adapters/input/http"
	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
	inputBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	business "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputContract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/adapters/input/http"
	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
	inputDailyLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/adapters/input/http"
	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
	inputAsyncEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/async"
	inputHttpEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	inputEventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/adapters/input/http"
	eventSubscription "github.com/davidgrcia/tfg-informatica-nextjs/pkg/event-subscription/usecases"
	inputGeneralLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/adapters/input/http"
	generalledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	inputUser "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/adapters/input/http"
	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

type Graph struct {
	*mux.Router

	//contablock
	ContablockClient *ethclient.Client
	contablock.ContaBlockContractAdapter
	contablock.ContaBlockUseCase

	//event subscriptions
	EventSubscriptionDb sql.DB
	eventSubscription.EventSubscriptionStorage
	eventSubscription.EventSubscriptionUseCase
	inputEventSubscription.EventSubscriptionInputAdapter

	//contracts
	ContractDb sql.DB
	contract.ContractStorage
	contract.ContractUseCase
	inputContract.ContractInputAdapter

	//users
	UserDb sql.DB
	user.UserStorage
	user.UserUseCase
	inputUser.UserInputAdapter

	//entries
	EntryDb sql.DB
	entry.EntryStorage
	entry.EntryUseCase
	inputHttpEntry.EntryInputHttpAdapter
	inputAsyncEntry.EntryInputAsyncAdapter

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

	//daily ledger
	dailyledger.DailyLedgerUseCase
	inputDailyLedger.DailyLedgerInputAdapter

	//general ledger
	generalledger.GeneralLedgerUseCase
	inputGeneralLedger.GeneralLedgerInputAdapter

	//balances
	BalanceDb sql.DB
	balance.BalanceStorage
	balance.BalanceUseCase
	inputBalance.BalanceInputAdapter
}

func NewGraph() Graph {
	return Graph{}
}

func (graph *Graph) Build() {
	// output infra
	graph.provideContablockClient()
	graph.provideContractDB()
	graph.provideEventSubscriptionDB()
	graph.provideUserDB()
	graph.provideAccountDB()
	graph.provideBusinessDB()
	graph.provideEntryDB()
	graph.provideBalanceDB()

	// output adapters
	graph.provideContablockContractAdapter()
	graph.provideContractStorage()
	graph.provideEventSubscriptionStorage()
	graph.provideUserStorage()
	graph.provideAccountStorage()
	graph.provideBusinessStorage()
	graph.provideEntryStorage()
	graph.provideBalanceStorage()

	// use cases
	graph.provideContablockUseCase()
	graph.provideContractUseCase()
	graph.provideEventSubscriptionUseCase()
	graph.provideUserUseCase()
	graph.provideBalanceUseCase()
	graph.provideAccountUseCase()
	graph.provideBusinessUseCase()
	graph.provideEntryUseCase()
	graph.provideDailyLedgerUseCase()
	graph.provideGeneralLedgerUseCase()

	// input adapters
	graph.provideAccountInputAdapter()
	graph.provideContractInputAdapter()
	graph.provideEventSubscriptionInputAdapter()
	graph.provideUserInputAdapter()
	graph.provideBusinessInputAdapter()
	graph.provideEntryInputHttpAdapter()
	graph.provideEntryInputAsyncAdapter()
	graph.provideDailyLedgerInputAdapter()
	graph.provideGeneralLedgerInputAdapter()
	graph.provideBalanceInputAdapter()

	// input infra
	graph.provideRouter()
	graph.provideContractEndpoints()
	graph.provideEventSubscriptionEndpoints()
	graph.provideAccountEndpoints()
	graph.provideUserEndpoints()
	graph.provideBusinessEndpoints()
	graph.provideEntryEndpoints()
	graph.provideDailyLedgerEndpoints()
	graph.provideGeneralLedgerEndpoints()
	graph.provideBalanceEndpoints()
}
