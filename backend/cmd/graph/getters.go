package graph

import (
	"database/sql"
	inputAdapterAccount "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/adapters/input/http"
	account "github.com/davidgrcia/tfg-informatica-nextjs/pkg/account/usecases"
	inputAdapterBalance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/adapters/input/http"
	balance "github.com/davidgrcia/tfg-informatica-nextjs/pkg/balance/usecases"
	inputAdapterBusiness "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/adapters/input/http"
	business "github.com/davidgrcia/tfg-informatica-nextjs/pkg/business/usecases"
	inputAdapterContract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/adapters/input/http"
	contract "github.com/davidgrcia/tfg-informatica-nextjs/pkg/contract/usecases"
	inputAdapterDailyLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/adapters/input/http"
	dailyledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/daily-ledger/usecases"
	inputAsyncAdapterEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/async"
	inputHttpAdapterEntry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/http"
	entry "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/usecases"
	inputAdapterGeneralLedger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/adapters/input/http"
	generalledger "github.com/davidgrcia/tfg-informatica-nextjs/pkg/general-ledger/usecases"
	contablock "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	inputAdapterUser "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/adapters/input/http"
	user "github.com/davidgrcia/tfg-informatica-nextjs/pkg/user/usecases"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
)

func (graph *Graph) GetRouter() *mux.Router {
	if graph.Router == nil {
		panic("graph Router is nil")
	}
	return graph.Router
}

/////// SMART CONTRACT CONTABLOCK ///////

func (graph *Graph) GetContablockClient() *ethclient.Client {
	return graph.ContablockClient
}

func (graph *Graph) GetContablockAdapter() contablock.ContaBlockContractAdapter {
	return graph.ContaBlockContractAdapter
}

func (graph *Graph) GetContaBlockUseCase() contablock.ContaBlockUseCase {
	return graph.ContaBlockUseCase
}

/////// ENTRIES ///////

func (graph *Graph) GetEntryInputHttpAdapter() inputHttpAdapterEntry.EntryInputHttpAdapter {
	return graph.EntryInputHttpAdapter
}

func (graph *Graph) GetEntryInputAsyncAdapter() inputAsyncAdapterEntry.EntryInputAsyncAdapter {
	return graph.EntryInputAsyncAdapter
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

/////// DAILY LEDGER ///////

func (graph *Graph) GetDailyLedgerInputAdapter() inputAdapterDailyLedger.DailyLedgerInputAdapter {
	return graph.DailyLedgerInputAdapter
}

func (graph *Graph) GetDailyLedgerUseCase() dailyledger.DailyLedgerUseCase {
	if graph.DailyLedgerUseCase == nil {
		panic("graph DailyLedgerUseCase case is nil")
	}
	return graph.DailyLedgerUseCase
}

/////// GENERAL LEDGER ///////

func (graph *Graph) GetGeneralLedgerInputAdapter() inputAdapterGeneralLedger.GeneralLedgerInputAdapter {
	return graph.GeneralLedgerInputAdapter
}

func (graph *Graph) GetGeneralLedgerUseCase() generalledger.GeneralLedgerUseCase {
	if graph.GeneralLedgerUseCase == nil {
		panic("graph GeneralLedgerUseCase case is nil")
	}
	return graph.GeneralLedgerUseCase
}

/////// BALANCES ///////

func (graph *Graph) GetBalanceInputAdapter() inputAdapterBalance.BalanceInputAdapter {
	return graph.BalanceInputAdapter
}

func (graph *Graph) GetBalanceDB() *sql.DB {
	return &graph.BalanceDb
}

func (graph *Graph) GetBalanceStorage() balance.BalanceStorage {
	if graph.BalanceStorage == nil {
		panic("graph BalanceStorage is nil")
	}
	return graph.BalanceStorage
}

func (graph *Graph) GetBalanceUseCase() balance.BalanceUseCase {
	if graph.BalanceUseCase == nil {
		panic("graph BalanceUseCase case is nil")
	}
	return graph.BalanceUseCase
}

/////// USERS ///////

func (graph *Graph) GetUserInputAdapter() inputAdapterUser.UserInputAdapter {
	return graph.UserInputAdapter
}

func (graph *Graph) GetUserDB() *sql.DB {
	return &graph.UserDb
}

func (graph *Graph) GetUserStorage() user.UserStorage {
	if graph.UserStorage == nil {
		panic("graph UserStorage is nil")
	}
	return graph.UserStorage
}

func (graph *Graph) GetUserUseCase() user.UserUseCase {
	if graph.UserUseCase == nil {
		panic("graph UserUseCase case is nil")
	}
	return graph.UserUseCase
}

/////// CONTRACTS ///////

func (graph *Graph) GetContractInputAdapter() inputAdapterContract.ContractInputAdapter {
	return graph.ContractInputAdapter
}

func (graph *Graph) GetContractDB() *sql.DB {
	return &graph.ContractDb
}

func (graph *Graph) GetContractStorage() contract.ContractStorage {
	if graph.ContractStorage == nil {
		panic("graph ContractStorage is nil")
	}
	return graph.ContractStorage
}

func (graph *Graph) GetContractUseCase() contract.ContractUseCase {
	if graph.ContractUseCase == nil {
		panic("graph ContractUseCase case is nil")
	}
	return graph.ContractUseCase
}
