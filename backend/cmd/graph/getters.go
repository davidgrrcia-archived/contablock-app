package graph

import (
	inputAdapterBet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/bet"
	inputAdapterGame "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/game"
	inputAdapterUser "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user"
	inputUserWallet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user-wallet"
	holdContractAdapter "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/holdContract"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/game"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
	userWallet "github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user-wallet"
	"github.com/gorilla/mux"
)

func (graph *Graph) GetRouter() *mux.Router {
	if graph.Router == nil {
		panic("graph Router is nil")
	}
	return graph.Router
}

/*
func (graph *Graph) GetHoldContract() *holdContract.HoldContract {
	return graph.HoldContract
}*/

func (graph *Graph) GetHoldContractAdapter() holdContractAdapter.HoldContractAdapter {
	return graph.HoldContractAdapter
}

/* USERS */

func (graph *Graph) GetUserInputAdapter() inputAdapterUser.UserInputAdapter {
	return graph.UserInputAdapter
}

func (graph *Graph) GetUserDB() user.UserCollection {
	return graph.UserDb
}

func (graph *Graph) GetUserStorage() user.UserStorage {
	if graph.UserStorage == nil {
		panic("graph UserWalletStorage is nil")
	}
	return graph.UserStorage
}

func (graph *Graph) GetUserUseCase() user.UserUseCase {
	if graph.UserUseCase == nil {
		panic("graph UserWalletUseCase case is nil")
	}
	return graph.UserUseCase
}

/* WALLET USER */

func (graph *Graph) GetUserWalletInputAdapter() inputUserWallet.UserWalletInputAdapter {
	return graph.UserWalletInputAdapter
}

func (graph *Graph) GetUserWalletDB() userWallet.UserWalletCollection {
	if graph.UserDb == nil {
		panic("graph UserDb is nil")
	}
	return graph.UserWalletDb
}

func (graph *Graph) GetUserWalletStorage() userWallet.UserWalletStorage {
	if graph.UserStorage == nil {
		panic("graph UserWalletStorage is nil")
	}
	return graph.UserWalletStorage
}

func (graph *Graph) GetUserWalletUseCase() userWallet.UserWalletUseCase {
	if graph.UserUseCase == nil {
		panic("graph UserWalletUseCase case is nil")
	}
	return graph.UserWalletUseCase
}

/* GAMES */

func (graph *Graph) GetGameInputAdapter() inputAdapterGame.GameInputAdapter {
	return graph.GameInputAdapter
}

func (graph *Graph) GetGameDB() game.GameCollection {
	return graph.GameDb
}

func (graph *Graph) GetGameStorage() game.GameStorage {
	if graph.GameStorage == nil {
		panic("graph GameWalletStorage is nil")
	}
	return graph.GameStorage
}

func (graph *Graph) GetGameUseCase() game.GameUseCase {
	if graph.GameUseCase == nil {
		panic("graph GameWalletUseCase case is nil")
	}
	return graph.GameUseCase
}

/* BETS */

func (graph *Graph) GetBetInputAdapter() inputAdapterBet.BetInputAdapter {
	return graph.BetInputAdapter
}

func (graph *Graph) GetBetDB() bet.BetCollection {
	return graph.BetDb
}

func (graph *Graph) GetBetStorage() bet.BetStorage {
	if graph.BetStorage == nil {
		panic("graph BetWalletStorage is nil")
	}
	return graph.BetStorage
}

func (graph *Graph) GetBetUseCase() bet.BetUseCase {
	if graph.BetUseCase == nil {
		panic("graph BetWalletUseCase case is nil")
	}
	return graph.BetUseCase
}
