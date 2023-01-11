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

func (adapter DefaultContaBlockContract) GetEntryData(input contablockUseCase.GetEntryDataAdapterInput) (*contablockUseCase.GetEntryDataAdapterOutput, error) {
	//input

	//client
	entryData, err := adapter.client.GetEntryData(nil, input.ID)
	if err != nil {
		return nil, err
	}

	var output contablockUseCase.GetEntryDataAdapterOutput
	output.ID = entryData.Id
	output.AccountID = entryData.AccountId
	output.MultipleEntryID = entryData.MultipleEntryId
	output.Currency = entryData.Currency
	output.Concept = entryData.Concept
	output.Date = entryData.Date
	output.DebitBalance = entryData.DebitBalance
	output.CreditBalance = entryData.CreditBalance
	output.State = entryData.State

	//output
	return &output, nil
}

func (adapter DefaultContaBlockContract) ListMultipleEntriesIds() (*contablockUseCase.ListMultipleEntriesIdsAdapterOutput, error) {
	//input

	//client
	multipleEntriesIds, err := adapter.client.ListMultipleEntryIds(nil)
	if err != nil {
		return nil, err
	}

	var output contablockUseCase.ListMultipleEntriesIdsAdapterOutput
	for _, item := range multipleEntriesIds {
		output.MultipleEntriesIds = append(output.MultipleEntriesIds, item)
	}

	//output
	return &output, nil
}

func (adapter DefaultContaBlockContract) ListAccountIds() (*contablockUseCase.ListAccountIdsAdapterOutput, error) {
	//input

	//client
	accountIds, err := adapter.client.ListAccountIds(nil)
	if err != nil {
		return nil, err
	}

	var output contablockUseCase.ListAccountIdsAdapterOutput
	for _, item := range accountIds {
		output.AccountIds = append(output.AccountIds, item)
	}

	//output
	return &output, nil
}

func (adapter DefaultContaBlockContract) PublishEntry(input contablockUseCase.PublishEntryAdapterInput) error {
	//input

	//client
	_, err := adapter.client.PublishEntry(adapter.auth, input.ID)
	if err != nil {
		return err
	}

	//output
	return nil
}
