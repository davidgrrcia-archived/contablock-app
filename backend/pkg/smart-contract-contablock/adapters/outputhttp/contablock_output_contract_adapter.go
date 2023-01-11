package outputhttp

import (
	"errors"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

func (adapter DefaultContaBlockContract) CreateEntry(input contablockUseCase.CreateEntryAdapterInput) error {
	//input

	//client
	_, err := adapter.client.CreateEntry(adapter.auth, input.ID, input.AccountID, input.MultipleEntryID, input.Currency, input.Concept, input.Date, input.DebitBalance, input.CreditBalance)
	if err != nil {
		return err
	}

	//output
	return nil
}
