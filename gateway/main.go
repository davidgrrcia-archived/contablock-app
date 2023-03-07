package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	"github.com/davidgrcia/tfg-informatica-nextjs/cmd/graph"
	entryInputAsyncAdapter "github.com/davidgrcia/tfg-informatica-nextjs/pkg/entry/adapters/input/async"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
	"golang.org/x/sync/errgroup"
	"gopkg.in/yaml.v2"
)

const (
	eventSubscriptionYamlPath = "./config/event-subscriptions.yaml"
)

type YamlFile struct {
	EventSubscriptions []EventSubscription `yaml:"eventSubscriptions"`
}

type EventSubscription struct {
	Name               string   `yaml:"name"`
	ProjectsSubscribed []string `yaml:"projectsSubscribed"`
}

func main() {
	err := godotenv.Load("../.env.local")
	if err != nil {
		log.Println(err)
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

	yaml := readYamlFile(eventSubscriptionYamlPath)

	eg := new(errgroup.Group)
	eg.Go(func() error {
		var input watchEntryCreatedInput
		input.instance = instance
		input.handler = entriesHandler
		input.yaml = yaml
		return watchEntryCreated(context.Background(), input)
	})
	if err = eg.Wait(); err != nil {
		panic(err)
	}

}

type watchEntryCreatedInput struct {
	instance *contablockClient.Contablock
	handler  entryInputAsyncAdapter.EntryInputAsyncAdapter
	yaml     YamlFile
}

func watchEntryCreated(ctx context.Context, input watchEntryCreatedInput) error {
	eventsChannel := make(chan *contablockClient.ContablockEntryCreated)
	opts := &bind.WatchOpts{
		Start:   nil,
		Context: ctx,
	}
	subscription, err := input.instance.WatchEntryCreated(opts, eventsChannel)
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

			log.Println("EntryCreated event received:", event.Entry)
			projects := getProjectsSubscribedToEvent("EntryCreated", input.yaml)
			for _, project := range projects {
				err = input.handler.AdaptEntryCreatedEvent(ctx, project, *event)
				if err != nil {
					log.Fatal(err)
				}
			}

		}
	}
}

func getProjectsSubscribedToEvent(event string, yaml YamlFile) []string {
	for _, eventSubscription := range yaml.EventSubscriptions {
		if eventSubscription.Name == event {
			return eventSubscription.ProjectsSubscribed
		}
	}
	return []string{}
}

func readYamlFile(path string) YamlFile {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var yamlFile YamlFile
	err = yaml.Unmarshal(data, &yamlFile)
	if err != nil {
		log.Fatal(err)
	}

	return yamlFile
}
