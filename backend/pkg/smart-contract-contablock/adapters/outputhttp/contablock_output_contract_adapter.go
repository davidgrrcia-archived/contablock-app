package outputhttp

import (
	"context"
	"errors"
	"fmt"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"math/big"
	"strconv"
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
	client *ethclient.Client
}

type DefaultContaBlockContractOptions struct {
	Client *ethclient.Client
}

var _ contablockUseCase.ContaBlockContractAdapter = (*DefaultContaBlockContract)(nil)

func (adapter DefaultContaBlockContract) GrantAccess(input contablockUseCase.GrantAccessAdapterInput) error {
	//input
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.AdminAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return err
	}

	userAddress := common.HexToAddress(input.UserAddress)

	//client
	_, err = session.Contract.GrantAccess(&session.TransactOpts, userAddress)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) CreateEntry(input contablockUseCase.CreateEntryAdapterInput) error {
	//input
	debitBalance := transformFloat64ToBigInt(input.DebitBalance)
	creditBalance := transformFloat64ToBigInt(input.CreditBalance)
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return err
	}

	//client
	_, err = session.Contract.CreateEntry(&session.TransactOpts, input.ID, input.AccountID, input.MultipleEntryID, input.Currency, input.Concept, input.Date, debitBalance, creditBalance)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) GetEntryData(input contablockUseCase.GetEntryDataAdapterInput) (*contablockUseCase.GetEntryDataAdapterOutput, error) {
	//input
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return nil, err
	}

	//client
	entryData, err := session.Contract.GetEntryData(&session.CallOpts, input.ID)
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

func (adapter DefaultContaBlockContract) ListMultipleEntriesIds(input contablockUseCase.ListMultipleEntriesIdsAdapterInput) (*contablockUseCase.ListMultipleEntriesIdsAdapterOutput, error) {
	//input
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return nil, err
	}

	fromString := session.CallOpts.From.String()

	_ = fromString

	//client
	multipleEntriesIds, err := session.Contract.ListMultipleEntryIds(&session.CallOpts)
	if err != nil {
		return nil, err
	}

	var output contablockUseCase.ListMultipleEntriesIdsAdapterOutput
	output.MultipleEntriesIds = append(output.MultipleEntriesIds, multipleEntriesIds...)

	//output
	return &output, nil
}

func (adapter DefaultContaBlockContract) ListAccountIds(input contablockUseCase.ListAccountIdsAdapterInput) (*contablockUseCase.ListAccountIdsAdapterOutput, error) {
	//input
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return nil, err
	}

	//client
	accountIds, err := session.Contract.ListAccountIds(&session.CallOpts)
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
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return err
	}

	//client
	_, err = session.Contract.PublishEntry(&session.TransactOpts, input.ID)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) RemoveEntry(input contablockUseCase.RemoveEntryAdapterInput) error {
	//input
	var clientSessionInput getClientSessionInput
	clientSessionInput.ContractAddress = input.ContractAddress
	clientSessionInput.FromAddress = input.UserAddress
	clientSessionInput.PrivateKey = input.UserPrivateKey
	session, err := adapter.getClientSession(clientSessionInput)
	if err != nil {
		return err
	}

	//client
	_, err = session.Contract.RemoveEntry(&session.TransactOpts, input.ID)
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

type getClientSessionInput struct {
	ContractAddress string
	PrivateKey      string
	FromAddress     string
}

func (adapter DefaultContaBlockContract) getClientSession(input getClientSessionInput) (*contablockClient.ContablockSession, error) {
	var (
		client     = adapter.client
		privateKey = crypto.ToECDSAUnsafe(common.FromHex(input.PrivateKey))
		address    = common.HexToAddress(input.ContractAddress)
		value      = new(big.Int).Mul(big.NewInt(0), big.NewInt(params.Ether)) // in wei (0 eth)
		sender     = common.HexToAddress(input.FromAddress)
		gasLimit   = uint64(30000000) // in units
	)

	// Retrieve the chainid (needed for signer)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	// Retrieve the pending nonce
	nonce, err := client.PendingNonceAt(context.Background(), sender)
	if err != nil {
		return nil, err
	}

	instance, err := contablockClient.NewContablock(address, client)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = value
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	var output contablockClient.ContablockSession
	output.Contract = instance
	output.TransactOpts = *auth
	output.CallOpts = bind.CallOpts{
		Pending: true,
		From:    sender,
		Context: context.Background(),
	}

	return &output, nil
}
