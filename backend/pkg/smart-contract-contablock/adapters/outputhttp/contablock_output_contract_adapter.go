package outputhttp

import (
	"math/big"

	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
)

type DefaultContaBlockContract struct {
	client contablockClient.ContablockSession
}

type DefaultContaBlockContractOptions struct {
	Client contablockClient.ContablockSession
}

func NewDefaultContaBlockContract(options DefaultContaBlockContractOptions) (*DefaultContaBlockContract, error) {
	return &DefaultContaBlockContract{
		client: options.Client,
	}, nil
}

var _ contablockUseCase.ContaBlockContractAdapter = (*DefaultContaBlockContract)(nil)

func (c DefaultContaBlockContract) Store(input contablockUseCase.StoreAdapterInput) error {
	//input
	var num *big.Int
	num = input.Num

	//client
	_, err := c.client.Store(num)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (c DefaultContaBlockContract) Retrieve() (*contablockUseCase.RetrieveAdapterOutput, error) {
	//input

	//client
	num, err := c.client.Retrieve()
	if err != nil {
		return nil, err
	}

	//output
	var output contablockUseCase.RetrieveAdapterOutput
	output.Num = num

	return &output, nil
}
