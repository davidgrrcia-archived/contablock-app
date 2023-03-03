package outputhttp

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	contablockClient "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/infra/outputhttp/generated"
	contablockUseCase "github.com/davidgrcia/tfg-informatica-nextjs/pkg/smart-contract-contablock/usecases"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
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
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return err
	}

	userAddress := common.HexToAddress(input.UserAddress)

	//client
	_, err = client.Instance.GrantAccess(client.Auth, userAddress)
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
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return err
	}

	//client
	_, err = client.Instance.CreateEntry(client.Auth, input.ID, input.AccountID, input.MultipleEntryID, input.Currency, input.Concept, input.Date, debitBalance, creditBalance)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) GetEntryData(input contablockUseCase.GetEntryDataAdapterInput) (*contablockUseCase.GetEntryDataAdapterOutput, error) {
	//input
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return nil, err
	}

	//client
	entryData, err := client.Instance.GetEntryData(nil, input.ID)
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
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return nil, err
	}

	//client
	multipleEntriesIds, err := client.Instance.ListMultipleEntryIds(nil)
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
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return nil, err
	}

	//client
	accountIds, err := client.Instance.ListAccountIds(nil)
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
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return err
	}

	//client
	_, err = client.Instance.PublishEntry(client.Auth, input.ID)
	if err != nil {
		return err
	}

	//output
	return nil
}

func (adapter DefaultContaBlockContract) RemoveEntry(input contablockUseCase.RemoveEntryAdapterInput) error {
	//input
	var clientInstanceInput getClientInstanceInput
	clientInstanceInput.ContractAddress = input.ContractAddress
	clientInstanceInput.FromAddress = input.UserAddress
	clientInstanceInput.PrivateKey = input.UserPrivateKey
	client, err := adapter.getClientInstance(clientInstanceInput)
	if err != nil {
		return err
	}

	//client
	_, err = client.Instance.RemoveEntry(client.Auth, input.ID)
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

type getClientInstanceOutput struct {
	Instance *contablockClient.Contablock
	Auth     *bind.TransactOpts
}

type getClientInstanceInput struct {
	ContractAddress string
	PrivateKey      string
	FromAddress     string
}

func (adapter DefaultContaBlockContract) getClientInstance(input getClientInstanceInput) (*getClientInstanceOutput, error) {
	client := adapter.client

	_, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(input.ContractAddress)
	instance, err := contablockClient.NewContablock(address, client)
	if err != nil {
		return nil, err
	}

	privateKeyString := strings.ReplaceAll(input.PrivateKey, "0x", "")
	privateKey, err := crypto.HexToECDSA(privateKeyString)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(30000000) // in units
	auth.GasPrice = gasPrice
	auth.From = common.HexToAddress(input.FromAddress) // todo review this bc input is not used

	var output getClientInstanceOutput
	output.Instance = instance
	output.Auth = auth

	return &output, nil
}
