package graph

import (
	inputBet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/bet"
	inputGame "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/game"
	inputUserWallet "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user-wallet"
	betOutputDbAdapter "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/database/bet"
	gameOutputDbAdapter "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/database/game"
	"github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/database/user"
	"github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/database/user-wallet"
	holdContractAdapter "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/output/holdContract"
	"github.com/davidevloper/HoldContract-go/backend/pkg/entities"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/bet"
	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/game"
	userWallet "github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user-wallet"
	"log"
	"time"

	"github.com/davidevloper/HoldContract-go/backend/pkg/infra/input"

	inputUser "github.com/davidevloper/HoldContract-go/backend/pkg/adapters/input/http/user"

	"github.com/davidevloper/HoldContract-go/backend/pkg/usecases/user"
	"github.com/gorilla/mux"
)

func (graph *Graph) provideRouter() {
	graph.Router = mux.NewRouter().StrictSlash(true)
}

/* USERS */

func (graph *Graph) provideUserDB() {
	graph.UserDb = user.UserCollection{
		{
			User: entities.User{
				ID:       "1",
				Email:    "test@test.com",
				Username: "david",
				Password: "123",
			},
			CreationDate: time.Now(),
			LastUpdate:   time.Now(),
		},
		{
			User: entities.User{
				ID:       "2",
				Email:    "test@test.com",
				Username: "maria",
				Password: "123",
			},
			CreationDate: time.Now(),
			LastUpdate:   time.Now(),
		},
	}
}

func (graph *Graph) provideUserStorage() user.UserStorage {
	userStorageOpts := userOutputDbAdapter.InMemoryUserStorageOptions{
		UserDatabase: graph.GetUserDB(),
	}
	userStorage, err := userOutputDbAdapter.NewInMemoryUserStorage(userStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.UserStorage = userStorage

	return userStorage
}

func (graph *Graph) provideUserUseCase() {
	userUseCaseOpts := user.DefaultUserUseCaseOptions{
		UserStorage: graph.GetUserStorage(),
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
	var options input.DefaultUsersHttpHandlerOptions
	options.Adapter = graph.GetUserInputAdapter()
	handler, err := input.NewDefaultUsersHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterUserRoutes(graph.GetRouter())
}

/* USER WALLET */

func (graph *Graph) provideUserWalletDB() {
	graph.UserWalletDb = userWallet.UserWalletCollection{}
}

func (graph *Graph) provideUserWalletStorage() userWallet.UserWalletStorage {
	userWalletStorageOpts := userWalletOutputDbAdapter.InMemoryUserWalletStorageOptions{
		UserWalletDatabase: graph.GetUserWalletDB(),
	}
	userWalletStorage, err := userWalletOutputDbAdapter.NewInMemoryUserWalletStorage(userWalletStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.UserWalletStorage = userWalletStorage

	return userWalletStorage
}

func (graph *Graph) provideUserWalletUseCase() {
	userWalletUseCaseOpts := userWallet.DefaultUserWalletUseCaseOptions{
		UserWalletStorage: graph.GetUserWalletStorage(),
		HoldContract:      graph.GetHoldContractAdapter(),
	}
	userWalletUseCase, err := userWallet.NewDefaultUserWalletUseCase(userWalletUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.UserWalletUseCase = userWalletUseCase
}

func (graph *Graph) ProvideUserWalletInputAdapter() *inputUserWallet.DefaultHttpInputAdapter {
	userWalletInputAdapterOpts := inputUserWallet.DefaultHttpInputAdapterOptions{
		UserWalletUseCase: graph.GetUserWalletUseCase(),
	}
	userWalletInputAdapter, err := inputUserWallet.NewDefaultHttpInputAdapter(userWalletInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.UserWalletInputAdapter = userWalletInputAdapter

	return userWalletInputAdapter
}

func (graph *Graph) provideUserWalletEndpoints() {
	var options input.DefaultUserWalletsHttpHandlerOptions
	options.Adapter = graph.GetUserWalletInputAdapter()
	handler, err := input.NewDefaultUserWalletsHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterUserWalletRoutes(graph.GetRouter())
}

func (graph *Graph) provideHoldContractInputAdapter() {

	holdContractInputAdapter, err := holdContractAdapter.NewDefaultHoldContractAdapter()
	if err != nil {
		panic(err)
	}

	graph.HoldContractAdapter = holdContractInputAdapter
}

/* GAMES */

func (graph *Graph) provideGameDB() {
	graph.GameDb = game.GameCollection{}
}

func (graph *Graph) provideGameStorage() game.GameStorage {
	gameStorageOpts := gameOutputDbAdapter.InMemoryGameStorageOptions{
		GameDatabase: graph.GetGameDB(),
	}
	gameStorage, err := gameOutputDbAdapter.NewInMemoryGameStorage(gameStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.GameStorage = gameStorage

	return gameStorage
}

func (graph *Graph) provideGameUseCase() {
	gameUseCaseOpts := game.DefaultGameUseCaseOptions{
		GameStorage:  graph.GetGameStorage(),
		UserUseCase:  graph.GetUserUseCase(),
		BetUseCase:   graph.GetBetUseCase(),
		HoldContract: graph.GetHoldContractAdapter(),
	}
	gameUseCase, err := game.NewDefaultGameUseCase(gameUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.GameUseCase = gameUseCase
}

func (graph *Graph) provideGameInputAdapter() *inputGame.DefaultHttpInputAdapter {
	gameInputAdapterOpts := inputGame.DefaultHttpInputAdapterOptions{
		GameUseCase: graph.GetGameUseCase(),
	}
	gameInputAdapter, err := inputGame.NewDefaultHttpInputAdapter(gameInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.GameInputAdapter = gameInputAdapter

	return gameInputAdapter
}

func (graph *Graph) provideGameEndpoints() {
	var options input.DefaultGamesHttpHandlerOptions
	options.Adapter = graph.GetGameInputAdapter()
	handler, err := input.NewDefaultGamesHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterGameRoutes(graph.GetRouter())
}

/* BETS */

func (graph *Graph) provideBetDB() {
	graph.BetDb = bet.BetCollection{}
}

func (graph *Graph) provideBetStorage() bet.BetStorage {
	betStorageOpts := betOutputDbAdapter.InMemoryBetStorageOptions{
		BetDatabase: graph.GetBetDB(),
	}
	betStorage, err := betOutputDbAdapter.NewInMemoryBetStorage(betStorageOpts)
	if err != nil {
		panic(err)
	}

	graph.BetStorage = betStorage

	return betStorage
}

func (graph *Graph) provideBetUseCase() {
	betUseCaseOpts := bet.DefaultBetUseCaseOptions{
		BetStorage: graph.GetBetStorage(),
	}
	betUseCase, err := bet.NewDefaultBetUseCase(betUseCaseOpts)
	if err != nil {
		panic(err)
	}

	graph.BetUseCase = betUseCase
}

func (graph *Graph) provideBetInputAdapter() *inputBet.DefaultHttpInputAdapter {
	betInputAdapterOpts := inputBet.DefaultHttpInputAdapterOptions{
		BetUseCase: graph.GetBetUseCase(),
	}
	betInputAdapter, err := inputBet.NewDefaultHttpInputAdapter(betInputAdapterOpts)
	if err != nil {
		panic(err)
	}

	graph.BetInputAdapter = betInputAdapter

	return betInputAdapter
}

func (graph *Graph) provideBetEndpoints() {
	var options input.DefaultBetHttpHandlerOptions
	options.Adapter = graph.GetBetInputAdapter()
	handler, err := input.NewDefaultBetHttpHandler(options)
	if err != nil {
		log.Fatal(err)
	}
	handler.RegisterBetRoutes(graph.GetRouter())
}
