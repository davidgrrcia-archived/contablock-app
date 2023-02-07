package outputhttp

import (
	"errors"
	"fmt"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"strconv"
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
	debitBalance := transformFloat64ToBigInt(input.DebitBalance)
	creditBalance := transformFloat64ToBigInt(input.CreditBalance)

	//client
	_, err := adapter.client.CreateEntry(adapter.auth, input.ID, input.AccountID, input.MultipleEntryID, input.Currency, input.Concept, input.Date, debitBalance, creditBalance)
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
	debitBalance, err := transformBigIntToFloat64(entryData.DebitBalance)
	if err != nil {
		return nil, err
	}
	output.DebitBalance = debitBalance
	creditBalance, err := transformBigIntToFloat64(entryData.CreditBalance)
	if err != nil {
		return nil, err
	}
	output.CreditBalance = creditBalance
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
	output.MultipleEntriesIds = append(output.MultipleEntriesIds, multipleEntriesIds...)

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
	output.AccountIds = append(output.AccountIds, accountIds...)

	//output
	return &output, nil
}

func (adapter DefaultContaBlockContract) PublishEntry(input contablockUseCase.PublishEntryAdapterInput) error {
	//input
	adapter.auth.Nonce.Add(adapter.auth.Nonce, big.NewInt(0)) //todo review

	//client
	_, err := adapter.client.PublishEntry(adapter.auth, input.ID)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) RemoveEntry(input contablockUseCase.RemoveEntryAdapterInput) error {
	//input

	//client
	_, err := adapter.client.RemoveEntry(adapter.auth, input.ID)
	if err != nil {
		return err
	}

	//output
	return nil
}

func transformBigIntToFloat64(num *big.Int) (*float64, error) {
	if num == nil {
		return nil, nil
	}

	string := num.String()

	float, err := strconv.ParseFloat(string, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &float, err
}

func transformFloat64ToBigInt(num *float64) *big.Int {
	if num == nil {
		zero := 0.00
		num = &zero
	}

	bigFloat := big.NewFloat(*num)
	bigInt, _ := bigFloat.Int(nil)

	return bigInt
}
