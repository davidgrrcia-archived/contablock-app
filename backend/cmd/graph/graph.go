package graph

import (
	inputUserWallet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user-wallet"
	holdContractAdapter "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/holdContract"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/game"
	userWallet "github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user-wallet"
	"github.com/davidevloper/HoldContract-go/eth-gateway/holdContract"
	"github.com/gorilla/mux"

	inputBet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/bet"
	inputGame "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/game"
	inputUser "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
)

type Graph struct {
	*mux.Router

	//holdContract
	holdContractAdapter.HoldContractAdapter
	holdContract.HoldContractInterface

	//users
	UserDb user.UserCollection
	user.UserStorage
	user.UserUseCase
	inputUser.UserInputAdapter

	//wallet users
	UserWalletDb userWallet.UserWalletCollection
	userWallet.UserWalletStorage
	userWallet.UserWalletUseCase
	inputUserWallet.UserWalletInputAdapter

	//games
	GameDb game.GameCollection
	game.GameStorage
	game.GameUseCase
	inputGame.GameInputAdapter

	//bets
	BetDb bet.BetCollection
	bet.BetStorage
	bet.BetUseCase
	inputBet.BetInputAdapter
}

func NewGraph() Graph {
	return Graph{}
}

func (graph *Graph) Build() {
	// output infra
	graph.provideUserDB()
	graph.provideUserWalletDB()
	graph.provideGameDB()
	graph.provideBetDB()

	// output adapters
	graph.provideUserStorage()
	graph.provideUserWalletStorage()
	graph.provideGameStorage()
	graph.provideBetStorage()
	graph.provideHoldContractInputAdapter()

	// use cases
	graph.provideUserUseCase()
	graph.provideUserWalletUseCase()
	graph.provideBetUseCase()
	graph.provideGameUseCase()

	// input adapters
	graph.provideUserInputAdapter()
	graph.ProvideUserWalletInputAdapter()
	graph.provideGameInputAdapter()
	graph.provideBetInputAdapter()

	// input infra
	graph.provideRouter()
	graph.provideUserEndpoints()
	graph.provideUserWalletEndpoints()
	graph.provideGameEndpoints()
	graph.provideBetEndpoints()
}
