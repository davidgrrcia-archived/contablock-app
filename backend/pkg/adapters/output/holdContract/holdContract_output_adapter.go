package holdContract

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"github.com/davidevloper/HoldContract-go/eth-gateway/holdContract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

/* INTERFACE ADAPTER */

type HoldContractAdapter interface {
	BalanceOf(ctx context.Context, address string) (int64, error)
	BalanceOnHoldOf(ctx context.Context, address string) (int64, error)
	Transfer(ctx context.Context, recipient string, amount int64) error
	TransferFrom(ctx context.Context, sender string, recipient string, amount int64) error
	Hold(ctx context.Context, recipient string, amount int64, id int64) error
	HoldFrom(ctx context.Context, sender string, recipient string, amount int64, id int64) error
	ExecuteHold(ctx context.Context, id int64) error
	RemoveHold(ctx context.Context, id int64) error
}

/* DEFAULT USE CASE TYPES */

type DefaultHoldContractAdapter struct {
	holdContract holdContract.HoldContract
}

type DefaultHoldContractAdapterOptions struct {
	HoldContract holdContract.HoldContract
}

var auth *bind.TransactOpts

/* CONSTRUCTOR */

func NewDefaultHoldContractAdapter() (*DefaultHoldContractAdapter, error) {

	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
		return nil, err
	} else {
		fmt.Println("Connected to ethereum client")
	}

	privateKey, err := crypto.HexToECDSA("53fc1da06fb832df4f28ddfa72c7856c70c990b42eff5b93cbc2e1994541b156")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth = bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	//get contract address from json
	jsonString := string(getContractAddressFromJSON())
	contractAddressString := strings.Replace(jsonString, "\"", "", -1)

	// create instance of the deployed contract to interact with it
	contractAddress := common.HexToAddress(contractAddressString)

	instance, err := holdContract.NewHoldContract(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	return &DefaultHoldContractAdapter{
		holdContract: *instance,
	}, nil
}

/* AUX VAR TO VIEW MISSING METHODS IMPLEMENTATION */

var _ HoldContractAdapter = (*DefaultHoldContractAdapter)(nil)

func (instance *DefaultHoldContractAdapter) BalanceOf(ctx context.Context, address string) (int64, error) {

	walletAddress := common.HexToAddress(address)

	amount, err := instance.holdContract.BalanceOf(nil, walletAddress)
	if err != nil {
		log.Fatal(err)
	}

	return amount.Int64(), nil
}

func (instance *DefaultHoldContractAdapter) BalanceOnHoldOf(ctx context.Context, address string) (int64, error) {

	walletAddress := common.HexToAddress(address)

	amount, err := instance.holdContract.BalanceOnHoldOf(nil, walletAddress)
	if err != nil {
		log.Fatal(err)
	}

	return amount.Int64(), nil
}

//todo fix transfer

func (instance *DefaultHoldContractAdapter) Transfer(ctx context.Context, recipient string, amount int64) error {

	recipientAddress := common.HexToAddress(recipient)

	_, err := instance.holdContract.Transfer(auth, recipientAddress, big.NewInt(amount))
	if err != nil {
		return err
	}

	return nil
}

func (instance *DefaultHoldContractAdapter) TransferFrom(ctx context.Context, sender string, recipient string, amount int64) error {

	senderAddress := common.HexToAddress(sender)
	recipientAddress := common.HexToAddress(recipient)

	_, err := instance.holdContract.TransferFrom(nil, senderAddress, recipientAddress, big.NewInt(amount))
	if err != nil {
		return err
	}

	return nil
}

func (instance *DefaultHoldContractAdapter) Hold(ctx context.Context, recipient string, amount int64, id int64) error {
	recipientAddress := common.HexToAddress(recipient)

	_, err := instance.holdContract.Hold(auth, recipientAddress, big.NewInt(amount), big.NewInt(id))
	if err != nil {
		return err
	}

	return nil
}

func (instance *DefaultHoldContractAdapter) HoldFrom(ctx context.Context, sender string, recipient string, amount int64, id int64) error {
	senderAddress := common.HexToAddress(sender)
	recipientAddress := common.HexToAddress(recipient)

	_, err := instance.holdContract.HoldFrom(nil, senderAddress, recipientAddress, big.NewInt(amount), big.NewInt(id))
	if err != nil {
		return err
	}

	return nil
}

func (instance *DefaultHoldContractAdapter) ExecuteHold(ctx context.Context, id int64) error {
	_, err := instance.holdContract.ExecuteHold(nil, big.NewInt(id))
	if err != nil {
		return err
	}

	return nil
}

func (instance *DefaultHoldContractAdapter) RemoveHold(ctx context.Context, id int64) error {
	_, err := instance.holdContract.RemoveHold(nil, big.NewInt(id))
	if err != nil {
		return err
	}

	return nil
}

func getContractAddressFromJSON() string {
	// Open the truffle generated json file
	jsonFile, err := os.Open("./smart-contracts/build/contracts/HoldContract.json")

	if err != nil {
		fmt.Println(err)
		return ""
	}

	// defer the closing of our json file
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}

	json.Unmarshal([]byte(byteValue), &result)

	addressJsonBytes, err := json.Marshal(result["networks"].(map[string]interface{})["5777"].(map[string]interface{})["address"])
	addressJsonString := string(addressJsonBytes)

	return addressJsonString
}
