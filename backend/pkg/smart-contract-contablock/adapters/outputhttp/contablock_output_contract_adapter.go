package outputhttp

import (
	"errors"
	"log"

	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

func NewDefaultContaBlockContract(options DefaultContaBlockContractOptions) (*DefaultContaBlockContract, error) {
	if options.Client == nil {
		return nil, errors.New("cannot create an instance of ContaBlockContract with a nil client")
	}

	return &DefaultContaBlockContract{
		client: options.Client,
	}, nil
}

type DefaultContaBlockContract struct {
	client *contablockClient.Contablock
}

type DefaultContaBlockContractOptions struct {
	Client *contablockClient.Contablock
}

var _ contablockUseCase.ContaBlockContractAdapter = (*DefaultContaBlockContract)(nil)

func (instance DefaultContaBlockContract) Store(input contablockUseCase.StoreAdapterInput) error {
	//input
	//var num *big.Int
	//num = input.Num
	//
	////client
	//_, err := instance.client.Store(num)
	//if err != nil {
	//	return err
	//}

	//output
	return nil
}

func (instance DefaultContaBlockContract) Retrieve() (*contablockUseCase.RetrieveAdapterOutput, error) {
	//input

	//client
	value, err := instance.client.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}

	//output
	var output contablockUseCase.RetrieveAdapterOutput
	output.Num = value

	return &output, nil
}
