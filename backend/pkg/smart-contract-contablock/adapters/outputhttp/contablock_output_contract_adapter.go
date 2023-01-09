package outputhttp

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"math/big"

	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

func NewDefaultContaBlockContract(options DefaultContaBlockContractOptions) (*DefaultContaBlockContract, error) {
	if options.Client == nil {
		return nil, errors.New("cannot create an instance of ContaBlockContract with a nil client")
	}
	if options.Auth == nil {
		return nil, errors.New("cannot create an instance of ContaBlockContract with a nil auth")
	}

	return &DefaultContaBlockContract{
		client: options.Client,
		auth:   options.Auth,
	}, nil
}

type DefaultContaBlockContract struct {
	client *contablockClient.Contablock
	auth   *bind.TransactOpts
}

type DefaultContaBlockContractOptions struct {
	Client *contablockClient.Contablock
	Auth   *bind.TransactOpts
}

var _ contablockUseCase.ContaBlockContractAdapter = (*DefaultContaBlockContract)(nil)

func (adapter DefaultContaBlockContract) Store(input contablockUseCase.StoreAdapterInput) error {
	//input
	var num *big.Int
	num = input.Num

	//client
	_, err := adapter.client.Store(adapter.auth, num)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) Retrieve() (*contablockUseCase.RetrieveAdapterOutput, error) {
	//input

	//client
	value, err := adapter.client.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}

	//output
	var output contablockUseCase.RetrieveAdapterOutput
	output.Num = value

	return &output, nil
}
