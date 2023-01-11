// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contablock

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContablockMetaData contains all meta data concerning the Contablock contract.
var ContablockMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"}],\"name\":\"EntryCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_creditBalance\",\"type\":\"uint256\"}],\"name\":\"createEntry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"entries\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"multipleEntryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"currency\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"concept\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"debitBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"creditBalance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d32806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c8063e521d3f11461003b578063f8756c0114610057575b600080fd5b610055600480360381019061005091906106ee565b61008e565b005b610071600480360381019061006c919061084c565b6101d0565b604051610085989796959493929190610923565b60405180910390f35b604051806101000160405280898152602001888152602001878152602001868152602001858152602001848152602001838152602001828152506000896040516100d89190610a07565b908152602001604051809103902060008201518160000190816100fb9190610c2a565b5060208201518160010190816101119190610c2a565b5060408201518160020190816101279190610c2a565b50606082015181600301908161013d9190610c2a565b5060808201518160040190816101539190610c2a565b5060a08201518160050190816101699190610c2a565b5060c0820151816006015560e082015181600701559050507f5b1db6dd582e363533bc294d439583c525a3de484a31bb721352020728afb11388888888888888886040516101be989796959493929190610923565b60405180910390a15050505050505050565b60008180516020810182018051848252602083016020850120818352809550505050505060009150905080600001805461020990610a4d565b80601f016020809104026020016040519081016040528092919081815260200182805461023590610a4d565b80156102825780601f1061025757610100808354040283529160200191610282565b820191906000526020600020905b81548152906001019060200180831161026557829003601f168201915b50505050509080600101805461029790610a4d565b80601f01602080910402602001604051908101604052809291908181526020018280546102c390610a4d565b80156103105780601f106102e557610100808354040283529160200191610310565b820191906000526020600020905b8154815290600101906020018083116102f357829003601f168201915b50505050509080600201805461032590610a4d565b80601f016020809104026020016040519081016040528092919081815260200182805461035190610a4d565b801561039e5780601f106103735761010080835404028352916020019161039e565b820191906000526020600020905b81548152906001019060200180831161038157829003601f168201915b5050505050908060030180546103b390610a4d565b80601f01602080910402602001604051908101604052809291908181526020018280546103df90610a4d565b801561042c5780601f106104015761010080835404028352916020019161042c565b820191906000526020600020905b81548152906001019060200180831161040f57829003601f168201915b50505050509080600401805461044190610a4d565b80601f016020809104026020016040519081016040528092919081815260200182805461046d90610a4d565b80156104ba5780601f1061048f576101008083540402835291602001916104ba565b820191906000526020600020905b81548152906001019060200180831161049d57829003601f168201915b5050505050908060050180546104cf90610a4d565b80601f01602080910402602001604051908101604052809291908181526020018280546104fb90610a4d565b80156105485780601f1061051d57610100808354040283529160200191610548565b820191906000526020600020905b81548152906001019060200180831161052b57829003601f168201915b5050505050908060060154908060070154905088565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6105c58261057c565b810181811067ffffffffffffffff821117156105e4576105e361058d565b5b80604052505050565b60006105f761055e565b905061060382826105bc565b919050565b600067ffffffffffffffff8211156106235761062261058d565b5b61062c8261057c565b9050602081019050919050565b82818337600083830152505050565b600061065b61065684610608565b6105ed565b90508281526020810184848401111561067757610676610577565b5b610682848285610639565b509392505050565b600082601f83011261069f5761069e610572565b5b81356106af848260208601610648565b91505092915050565b6000819050919050565b6106cb816106b8565b81146106d657600080fd5b50565b6000813590506106e8816106c2565b92915050565b600080600080600080600080610100898b03121561070f5761070e610568565b5b600089013567ffffffffffffffff81111561072d5761072c61056d565b5b6107398b828c0161068a565b985050602089013567ffffffffffffffff81111561075a5761075961056d565b5b6107668b828c0161068a565b975050604089013567ffffffffffffffff8111156107875761078661056d565b5b6107938b828c0161068a565b965050606089013567ffffffffffffffff8111156107b4576107b361056d565b5b6107c08b828c0161068a565b955050608089013567ffffffffffffffff8111156107e1576107e061056d565b5b6107ed8b828c0161068a565b94505060a089013567ffffffffffffffff81111561080e5761080d61056d565b5b61081a8b828c0161068a565b93505060c061082b8b828c016106d9565b92505060e061083c8b828c016106d9565b9150509295985092959890939650565b60006020828403121561086257610861610568565b5b600082013567ffffffffffffffff8111156108805761087f61056d565b5b61088c8482850161068a565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108cf5780820151818401526020810190506108b4565b60008484015250505050565b60006108e682610895565b6108f081856108a0565b93506109008185602086016108b1565b6109098161057c565b840191505092915050565b61091d816106b8565b82525050565b600061010082019050818103600083015261093e818b6108db565b90508181036020830152610952818a6108db565b9050818103604083015261096681896108db565b9050818103606083015261097a81886108db565b9050818103608083015261098e81876108db565b905081810360a08301526109a281866108db565b90506109b160c0830185610914565b6109be60e0830184610914565b9998505050505050505050565b600081905092915050565b60006109e182610895565b6109eb81856109cb565b93506109fb8185602086016108b1565b80840191505092915050565b6000610a1382846109d6565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a6557607f821691505b602082108103610a7857610a77610a1e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610ae07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610aa3565b610aea8683610aa3565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610b27610b22610b1d846106b8565b610b02565b6106b8565b9050919050565b6000819050919050565b610b4183610b0c565b610b55610b4d82610b2e565b848454610ab0565b825550505050565b600090565b610b6a610b5d565b610b75818484610b38565b505050565b5b81811015610b9957610b8e600082610b62565b600181019050610b7b565b5050565b601f821115610bde57610baf81610a7e565b610bb884610a93565b81016020851015610bc7578190505b610bdb610bd385610a93565b830182610b7a565b50505b505050565b600082821c905092915050565b6000610c0160001984600802610be3565b1980831691505092915050565b6000610c1a8383610bf0565b9150826002028217905092915050565b610c3382610895565b67ffffffffffffffff811115610c4c57610c4b61058d565b5b610c568254610a4d565b610c61828285610b9d565b600060209050601f831160018114610c945760008415610c82578287015190505b610c8c8582610c0e565b865550610cf4565b601f198416610ca286610a7e565b60005b82811015610cca57848901518255600182019150602085019450602081019050610ca5565b86831015610ce75784890151610ce3601f891682610bf0565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220354de15bc440572e967f70a8573f0c85ea028c614b162ed5867a268f6524de1064736f6c63430008110033",
}

// ContablockABI is the input ABI used to generate the binding from.
// Deprecated: Use ContablockMetaData.ABI instead.
var ContablockABI = ContablockMetaData.ABI

// ContablockBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContablockMetaData.Bin instead.
var ContablockBin = ContablockMetaData.Bin

// DeployContablock deploys a new Ethereum contract, binding an instance of Contablock to it.
func DeployContablock(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contablock, error) {
	parsed, err := ContablockMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContablockBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contablock{ContablockCaller: ContablockCaller{contract: contract}, ContablockTransactor: ContablockTransactor{contract: contract}, ContablockFilterer: ContablockFilterer{contract: contract}}, nil
}

// Contablock is an auto generated Go binding around an Ethereum contract.
type Contablock struct {
	ContablockCaller     // Read-only binding to the contract
	ContablockTransactor // Write-only binding to the contract
	ContablockFilterer   // Log filterer for contract events
}

// ContablockCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContablockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContablockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContablockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContablockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContablockSession struct {
	Contract     *Contablock       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContablockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContablockCallerSession struct {
	Contract *ContablockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ContablockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContablockTransactorSession struct {
	Contract     *ContablockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ContablockRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContablockRaw struct {
	Contract *Contablock // Generic contract binding to access the raw methods on
}

// ContablockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContablockCallerRaw struct {
	Contract *ContablockCaller // Generic read-only contract binding to access the raw methods on
}

// ContablockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContablockTransactorRaw struct {
	Contract *ContablockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContablock creates a new instance of Contablock, bound to a specific deployed contract.
func NewContablock(address common.Address, backend bind.ContractBackend) (*Contablock, error) {
	contract, err := bindContablock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contablock{ContablockCaller: ContablockCaller{contract: contract}, ContablockTransactor: ContablockTransactor{contract: contract}, ContablockFilterer: ContablockFilterer{contract: contract}}, nil
}

// NewContablockCaller creates a new read-only instance of Contablock, bound to a specific deployed contract.
func NewContablockCaller(address common.Address, caller bind.ContractCaller) (*ContablockCaller, error) {
	contract, err := bindContablock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContablockCaller{contract: contract}, nil
}

// NewContablockTransactor creates a new write-only instance of Contablock, bound to a specific deployed contract.
func NewContablockTransactor(address common.Address, transactor bind.ContractTransactor) (*ContablockTransactor, error) {
	contract, err := bindContablock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContablockTransactor{contract: contract}, nil
}

// NewContablockFilterer creates a new log filterer instance of Contablock, bound to a specific deployed contract.
func NewContablockFilterer(address common.Address, filterer bind.ContractFilterer) (*ContablockFilterer, error) {
	contract, err := bindContablock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContablockFilterer{contract: contract}, nil
}

// bindContablock binds a generic wrapper to an already deployed contract.
func bindContablock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContablockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contablock *ContablockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contablock.Contract.ContablockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contablock *ContablockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contablock.Contract.ContablockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contablock *ContablockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contablock.Contract.ContablockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contablock *ContablockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contablock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contablock *ContablockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contablock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contablock *ContablockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contablock.Contract.contract.Transact(opts, method, params...)
}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockCaller) Entries(opts *bind.CallOpts, arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
}, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "entries", arg0)

	outstruct := new(struct {
		Id              string
		AccountId       string
		MultipleEntryId string
		Currency        string
		Concept         string
		Date            string
		DebitBalance    *big.Int
		CreditBalance   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.AccountId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.MultipleEntryId = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Currency = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Concept = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Date = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.DebitBalance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CreditBalance = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockSession) Entries(arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
}, error) {
	return _Contablock.Contract.Entries(&_Contablock.CallOpts, arg0)
}

// Entries is a free data retrieval call binding the contract method 0xf8756c01.
//
// Solidity: function entries(string ) view returns(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockCallerSession) Entries(arg0 string) (struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
}, error) {
	return _Contablock.Contract.Entries(&_Contablock.CallOpts, arg0)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockTransactor) CreateEntry(opts *bind.TransactOpts, _id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.contract.Transact(opts, "createEntry", _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockSession) CreateEntry(_id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.CreateEntry(&_Contablock.TransactOpts, _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// CreateEntry is a paid mutator transaction binding the contract method 0xe521d3f1.
//
// Solidity: function createEntry(string _id, string _accountId, string _multipleEntryId, string _currency, string _concept, string _date, uint256 _debitBalance, uint256 _creditBalance) returns()
func (_Contablock *ContablockTransactorSession) CreateEntry(_id string, _accountId string, _multipleEntryId string, _currency string, _concept string, _date string, _debitBalance *big.Int, _creditBalance *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.CreateEntry(&_Contablock.TransactOpts, _id, _accountId, _multipleEntryId, _currency, _concept, _date, _debitBalance, _creditBalance)
}

// ContablockEntryCreatedIterator is returned from FilterEntryCreated and is used to iterate over the raw logs and unpacked data for EntryCreated events raised by the Contablock contract.
type ContablockEntryCreatedIterator struct {
	Event *ContablockEntryCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContablockEntryCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContablockEntryCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContablockEntryCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContablockEntryCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContablockEntryCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContablockEntryCreated represents a EntryCreated event raised by the Contablock contract.
type ContablockEntryCreated struct {
	Id              string
	AccountId       string
	MultipleEntryId string
	Currency        string
	Concept         string
	Date            string
	DebitBalance    *big.Int
	CreditBalance   *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterEntryCreated is a free log retrieval operation binding the contract event 0x5b1db6dd582e363533bc294d439583c525a3de484a31bb721352020728afb113.
//
// Solidity: event EntryCreated(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockFilterer) FilterEntryCreated(opts *bind.FilterOpts) (*ContablockEntryCreatedIterator, error) {

	logs, sub, err := _Contablock.contract.FilterLogs(opts, "EntryCreated")
	if err != nil {
		return nil, err
	}
	return &ContablockEntryCreatedIterator{contract: _Contablock.contract, event: "EntryCreated", logs: logs, sub: sub}, nil
}

// WatchEntryCreated is a free log subscription operation binding the contract event 0x5b1db6dd582e363533bc294d439583c525a3de484a31bb721352020728afb113.
//
// Solidity: event EntryCreated(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockFilterer) WatchEntryCreated(opts *bind.WatchOpts, sink chan<- *ContablockEntryCreated) (event.Subscription, error) {

	logs, sub, err := _Contablock.contract.WatchLogs(opts, "EntryCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContablockEntryCreated)
				if err := _Contablock.contract.UnpackLog(event, "EntryCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEntryCreated is a log parse operation binding the contract event 0x5b1db6dd582e363533bc294d439583c525a3de484a31bb721352020728afb113.
//
// Solidity: event EntryCreated(string id, string accountId, string multipleEntryId, string currency, string concept, string date, uint256 debitBalance, uint256 creditBalance)
func (_Contablock *ContablockFilterer) ParseEntryCreated(log types.Log) (*ContablockEntryCreated, error) {
	event := new(ContablockEntryCreated)
	if err := _Contablock.contract.UnpackLog(event, "EntryCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
