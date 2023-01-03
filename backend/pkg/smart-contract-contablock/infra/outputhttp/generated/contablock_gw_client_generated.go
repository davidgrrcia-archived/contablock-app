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
	ABI: "[{\"inputs\":[],\"name\":\"retrieve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"store\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610150806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610059575b600080fd5b610043610075565b60405161005091906100a1565b60405180910390f35b610073600480360381019061006e91906100ed565b61007e565b005b60008054905090565b8060008190555050565b6000819050919050565b61009b81610088565b82525050565b60006020820190506100b66000830184610092565b92915050565b600080fd5b6100ca81610088565b81146100d557600080fd5b50565b6000813590506100e7816100c1565b92915050565b600060208284031215610103576101026100bc565b5b6000610111848285016100d8565b9150509291505056fea2646970667358221220b0bfdaddb28e27261ee68839a78b62598c7bd03c0bc399ff307b184100c6744964736f6c63430008110033",
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

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Contablock *ContablockCaller) Retrieve(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contablock.contract.Call(opts, &out, "retrieve")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Contablock *ContablockSession) Retrieve() (*big.Int, error) {
	return _Contablock.Contract.Retrieve(&_Contablock.CallOpts)
}

// Retrieve is a free data retrieval call binding the contract method 0x2e64cec1.
//
// Solidity: function retrieve() view returns(uint256)
func (_Contablock *ContablockCallerSession) Retrieve() (*big.Int, error) {
	return _Contablock.Contract.Retrieve(&_Contablock.CallOpts)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Contablock *ContablockTransactor) Store(opts *bind.TransactOpts, num *big.Int) (*types.Transaction, error) {
	return _Contablock.contract.Transact(opts, "store", num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Contablock *ContablockSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.Store(&_Contablock.TransactOpts, num)
}

// Store is a paid mutator transaction binding the contract method 0x6057361d.
//
// Solidity: function store(uint256 num) returns()
func (_Contablock *ContablockTransactorSession) Store(num *big.Int) (*types.Transaction, error) {
	return _Contablock.Contract.Store(&_Contablock.TransactOpts, num)
}
