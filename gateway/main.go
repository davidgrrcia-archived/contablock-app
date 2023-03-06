package main

import (
	"context"
	entryInputAsyncAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/async"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"log"
	"os"

	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
)

func main() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		panic(err)
	}

	websocketClient, err := ethclient.Dial(os.Getenv("SC_CLIENT_NETWORK_WEBSOCKET"))
	if err != nil {
		panic(err)
	}

	contractAddress := common.HexToAddress(os.Getenv("CONTRACT_ADDRESS"))

	instance, err := contablockClient.NewContablock(contractAddress, websocketClient)
	if err != nil {
		panic(err)
	}

	var backendAppGraph = graph.NewGraph()
	backendAppGraph.Build()
	entriesHandler := backendAppGraph.GetEntryInputAsyncAdapter()

	eg := new(errgroup.Group)
	eg.Go(func() error {
		return watchEntryCreated(context.Background(), instance, entriesHandler)
	})
	if err = eg.Wait(); err != nil {
		panic(err)
	}

}

func watchEntryCreated(ctx context.Context, instance *contablockClient.Contablock, handler entryInputAsyncAdapter.EntryInputAsyncAdapter) error {
	eventsChannel := make(chan *contablockClient.ContablockEntryCreated)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	subscription, err := instance.WatchEntryCreated(opts, eventsChannel)
	if err != nil {
		return err
	}
	defer subscription.Unsubscribe()

	for {
		select {
		case <-ctx.Done():
			return nil
		case errChan := <-subscription.Err():
			return errChan
		case event := <-eventsChannel:
			err = handler.AdaptEntryCreatedEvent(ctx, *event)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
