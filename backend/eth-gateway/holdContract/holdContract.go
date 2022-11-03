// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package holdContract

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

// HoldContractMetaData contains all meta data concerning the HoldContract contract.
var HoldContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_initialAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOnHoldOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"hold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"holdFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"executeHold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"holdId\",\"type\":\"uint256\"}],\"name\":\"removeHold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162001ce838038062001ce88339818101604052810190620000379190620000e1565b33600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555080600081905550506200013c565b600081519050620000db8162000122565b92915050565b600060208284031215620000fa57620000f96200011d565b5b60006200010a84828501620000ca565b91505092915050565b6000819050919050565b600080fd5b6200012d8162000113565b81146200013957600080fd5b50565b611b9c806200014c6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806323b872dd1161007157806323b872dd1461018c5780633ca22cb1146101bc57806370a08231146101ec578063743d3e5b1461021c578063a9059cbb1461024c578063dd62ed3e1461027c576100a9565b806308e53826146100ae578063095ea7b3146100de5780630cfd12bc1461010e57806312c383641461013e57806318160ddd1461016e575b600080fd5b6100c860048036038101906100c391906115c8565b6102ac565b6040516100d59190611708565b60405180910390f35b6100f860048036038101906100f39190611535565b61054a565b6040516101059190611708565b60405180910390f35b61012860048036038101906101239190611575565b61063c565b6040516101359190611708565b60405180910390f35b610158600480360381019061015391906115c8565b6108d5565b6040516101659190611708565b60405180910390f35b610176610b4b565b6040516101839190611803565b60405180910390f35b6101a660048036038101906101a1919061147b565b610b54565b6040516101b39190611708565b60405180910390f35b6101d660048036038101906101d191906114ce565b610e14565b6040516101e39190611708565b60405180910390f35b6102066004803603810190610201919061140e565b61112c565b6040516102139190611803565b60405180910390f35b6102366004803603810190610231919061140e565b611175565b6040516102439190611803565b60405180910390f35b61026660048036038101906102619190611535565b6111be565b6040516102739190611708565b60405180910390f35b6102966004803603810190610291919061143b565b61135d565b6040516102a39190611803565b60405180910390f35b6000806007600084815260200190815260200160002054905060006005600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060006006600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156103a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039c90611763565b60405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610435576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161042c90611743565b60405180910390fd5b6000831415610479576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610470906117a3565b60405180910390fd5b82600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546104c89190611885565b9250508190555082600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461051e919061182f565b925050819055506000600760008781526020019081526020016000208190555060019350505050919050565b600081600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258460405161062a9190611803565b60405180910390a36001905092915050565b600082600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410156106c0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106b790611723565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff166005600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610762576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610759906117e3565b60405180910390fd5b82600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546107b19190611885565b9250508190555082600460008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610807919061182f565b92505081905550336005600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550836006600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826007600084815260200190815260200160002081905550600190509392505050565b6000806007600084815260200190815260200160002054905060006006600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610996576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098d90611763565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166005600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610a37576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a2e90611783565b60405180910390fd5b6000821415610a7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a72906117a3565b60405180910390fd5b81600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610aca9190611885565b9250508190555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610b20919061182f565b9250508190555060006007600086815260200190815260200160002081905550600192505050919050565b60008054905090565b600080600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410158015610c255750828110155b610c64576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5b906117c3565b60405180910390fd5b82600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610cb3919061182f565b9250508190555082600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d099190611885565b9250508190555082600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610d9c9190611885565b925050819055508373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef85604051610e009190611803565b60405180910390a360019150509392505050565b600080600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905083600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410158015610ee55750838110155b610f24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f1b906117c3565b60405180910390fd5b83600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f73919061182f565b9250508190555083600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610fc99190611885565b9250508190555083600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461105c9190611885565b92505081905550856005600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550846006600085815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508360076000858152602001908152602001600020819055506001915050949350505050565b6000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600081600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020541015611242576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123990611723565b60405180910390fd5b81600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112919190611885565b9250508190555081600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546112e7919061182f565b925050819055508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161134b9190611803565b60405180910390a36001905092915050565b6000600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000813590506113f381611b38565b92915050565b60008135905061140881611b4f565b92915050565b60006020828403121561142457611423611930565b5b6000611432848285016113e4565b91505092915050565b6000806040838503121561145257611451611930565b5b6000611460858286016113e4565b9250506020611471858286016113e4565b9150509250929050565b60008060006060848603121561149457611493611930565b5b60006114a2868287016113e4565b93505060206114b3868287016113e4565b92505060406114c4868287016113f9565b9150509250925092565b600080600080608085870312156114e8576114e7611930565b5b60006114f6878288016113e4565b9450506020611507878288016113e4565b9350506040611518878288016113f9565b9250506060611529878288016113f9565b91505092959194509250565b6000806040838503121561154c5761154b611930565b5b600061155a858286016113e4565b925050602061156b858286016113f9565b9150509250929050565b60008060006060848603121561158e5761158d611930565b5b600061159c868287016113e4565b93505060206115ad868287016113f9565b92505060406115be868287016113f9565b9150509250925092565b6000602082840312156115de576115dd611930565b5b60006115ec848285016113f9565b91505092915050565b6115fe816118cb565b82525050565b6000611611602f8361181e565b915061161c82611935565b604082019050919050565b6000611634602e8361181e565b915061163f82611984565b604082019050919050565b6000611657600f8361181e565b9150611662826119d3565b602082019050919050565b600061167a60428361181e565b9150611685826119fc565b606082019050919050565b600061169d60218361181e565b91506116a882611a71565b604082019050919050565b60006116c060398361181e565b91506116cb82611ac0565b604082019050919050565b60006116e360168361181e565b91506116ee82611b0f565b602082019050919050565b611702816118f7565b82525050565b600060208201905061171d60008301846115f5565b92915050565b6000602082019050818103600083015261173c81611604565b9050919050565b6000602082019050818103600083015261175c81611627565b9050919050565b6000602082019050818103600083015261177c8161164a565b9050919050565b6000602082019050818103600083015261179c8161166d565b9050919050565b600060208201905081810360008301526117bc81611690565b9050919050565b600060208201905081810360008301526117dc816116b3565b9050919050565b600060208201905081810360008301526117fc816116d6565b9050919050565b600060208201905061181860008301846116f9565b92915050565b600082825260208201905092915050565b600061183a826118f7565b9150611845836118f7565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561187a57611879611901565b5b828201905092915050565b6000611890826118f7565b915061189b836118f7565b9250828210156118ae576118ad611901565b5b828203905092915050565b60006118c4826118d7565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600080fd5b7f746f6b656e2062616c616e6365206973206c6f776572207468616e207468652060008201527f76616c7565207265717565737465640000000000000000000000000000000000602082015250565b7f69742063616e206f6e6c7920626520657865637574656420627920746865206360008201527f6f6e7472616374206f776e65722e000000000000000000000000000000000000602082015250565b7f696e76616c696420686f6c642069640000000000000000000000000000000000600082015250565b7f69742063616e206f6e6c7920626520657865637574656420627920746865207360008201527f616d652061646472657373207468617420637265617465642074686520686f6c60208201527f642e000000000000000000000000000000000000000000000000000000000000604082015250565b7f616d6f756e74206f6e20686f6c6420616c7265616479207472616e736665726560008201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b7f746f6b656e2062616c616e6365206f7220616c6c6f77616e6365206973206c6f60008201527f776572207468616e20616d6f756e742072657175657374656400000000000000602082015250565b7f686f6c64206964206d75737420626520756e6971756500000000000000000000600082015250565b611b41816118b9565b8114611b4c57600080fd5b50565b611b58816118f7565b8114611b6357600080fd5b5056fea264697066735822122047ea98f98d2f8e828c5fc9dfde94a370cac5225d135b3827160e0478aef5fafe64736f6c63430008070033",
}

// HoldContractABI is the input ABI used to generate the binding from.
// Deprecated: Use HoldContractMetaData.ABI instead.
var HoldContractABI = HoldContractMetaData.ABI

// HoldContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use HoldContractMetaData.Bin instead.
var HoldContractBin = HoldContractMetaData.Bin

// DeployHoldContract deploys a new Ethereum contract, binding an instance of HoldContract to it.
func DeployHoldContract(auth *bind.TransactOpts, backend bind.ContractBackend, _initialAmount *big.Int) (common.Address, *types.Transaction, *HoldContract, error) {
	parsed, err := HoldContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(HoldContractBin), backend, _initialAmount)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HoldContract{HoldContractCaller: HoldContractCaller{contract: contract}, HoldContractTransactor: HoldContractTransactor{contract: contract}, HoldContractFilterer: HoldContractFilterer{contract: contract}}, nil
}

// HoldContract is an auto generated Go binding around an Ethereum contract.
type HoldContract struct {
	HoldContractCaller     // Read-only binding to the contract
	HoldContractTransactor // Write-only binding to the contract
	HoldContractFilterer   // Log filterer for contract events
}

// HoldContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type HoldContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoldContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HoldContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoldContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HoldContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HoldContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HoldContractSession struct {
	Contract     *HoldContract     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HoldContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HoldContractCallerSession struct {
	Contract *HoldContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// HoldContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HoldContractTransactorSession struct {
	Contract     *HoldContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// HoldContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type HoldContractRaw struct {
	Contract *HoldContract // Generic contract binding to access the raw methods on
}

// HoldContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HoldContractCallerRaw struct {
	Contract *HoldContractCaller // Generic read-only contract binding to access the raw methods on
}

// HoldContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HoldContractTransactorRaw struct {
	Contract *HoldContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHoldContract creates a new instance of HoldContract, bound to a specific deployed contract.
func NewHoldContract(address common.Address, backend bind.ContractBackend) (*HoldContract, error) {
	contract, err := bindHoldContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HoldContract{HoldContractCaller: HoldContractCaller{contract: contract}, HoldContractTransactor: HoldContractTransactor{contract: contract}, HoldContractFilterer: HoldContractFilterer{contract: contract}}, nil
}

// NewHoldContractCaller creates a new read-only instance of HoldContract, bound to a specific deployed contract.
func NewHoldContractCaller(address common.Address, caller bind.ContractCaller) (*HoldContractCaller, error) {
	contract, err := bindHoldContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HoldContractCaller{contract: contract}, nil
}

// NewHoldContractTransactor creates a new write-only instance of HoldContract, bound to a specific deployed contract.
func NewHoldContractTransactor(address common.Address, transactor bind.ContractTransactor) (*HoldContractTransactor, error) {
	contract, err := bindHoldContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HoldContractTransactor{contract: contract}, nil
}

// NewHoldContractFilterer creates a new log filterer instance of HoldContract, bound to a specific deployed contract.
func NewHoldContractFilterer(address common.Address, filterer bind.ContractFilterer) (*HoldContractFilterer, error) {
	contract, err := bindHoldContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HoldContractFilterer{contract: contract}, nil
}

// bindHoldContract binds a generic wrapper to an already deployed contract.
func bindHoldContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HoldContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoldContract *HoldContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoldContract.Contract.HoldContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoldContract *HoldContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoldContract.Contract.HoldContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoldContract *HoldContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoldContract.Contract.HoldContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HoldContract *HoldContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _HoldContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HoldContract *HoldContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HoldContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HoldContract *HoldContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HoldContract.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_HoldContract *HoldContractCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HoldContract.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_HoldContract *HoldContractSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HoldContract.Contract.Allowance(&_HoldContract.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_HoldContract *HoldContractCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _HoldContract.Contract.Allowance(&_HoldContract.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HoldContract.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HoldContract.Contract.BalanceOf(&_HoldContract.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _HoldContract.Contract.BalanceOf(&_HoldContract.CallOpts, account)
}

// BalanceOnHoldOf is a free data retrieval call binding the contract method 0x743d3e5b.
//
// Solidity: function balanceOnHoldOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractCaller) BalanceOnHoldOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _HoldContract.contract.Call(opts, &out, "balanceOnHoldOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOnHoldOf is a free data retrieval call binding the contract method 0x743d3e5b.
//
// Solidity: function balanceOnHoldOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractSession) BalanceOnHoldOf(account common.Address) (*big.Int, error) {
	return _HoldContract.Contract.BalanceOnHoldOf(&_HoldContract.CallOpts, account)
}

// BalanceOnHoldOf is a free data retrieval call binding the contract method 0x743d3e5b.
//
// Solidity: function balanceOnHoldOf(address account) view returns(uint256 balance)
func (_HoldContract *HoldContractCallerSession) BalanceOnHoldOf(account common.Address) (*big.Int, error) {
	return _HoldContract.Contract.BalanceOnHoldOf(&_HoldContract.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HoldContract *HoldContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _HoldContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HoldContract *HoldContractSession) TotalSupply() (*big.Int, error) {
	return _HoldContract.Contract.TotalSupply(&_HoldContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_HoldContract *HoldContractCallerSession) TotalSupply() (*big.Int, error) {
	return _HoldContract.Contract.TotalSupply(&_HoldContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Approve(&_HoldContract.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Approve(&_HoldContract.TransactOpts, spender, amount)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactor) ExecuteHold(opts *bind.TransactOpts, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "executeHold", holdId)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractSession) ExecuteHold(holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.ExecuteHold(&_HoldContract.TransactOpts, holdId)
}

// ExecuteHold is a paid mutator transaction binding the contract method 0x12c38364.
//
// Solidity: function executeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) ExecuteHold(holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.ExecuteHold(&_HoldContract.TransactOpts, holdId)
}

// Hold is a paid mutator transaction binding the contract method 0x0cfd12bc.
//
// Solidity: function hold(address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactor) Hold(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "hold", recipient, amount, holdId)
}

// Hold is a paid mutator transaction binding the contract method 0x0cfd12bc.
//
// Solidity: function hold(address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractSession) Hold(recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Hold(&_HoldContract.TransactOpts, recipient, amount, holdId)
}

// Hold is a paid mutator transaction binding the contract method 0x0cfd12bc.
//
// Solidity: function hold(address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) Hold(recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Hold(&_HoldContract.TransactOpts, recipient, amount, holdId)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x3ca22cb1.
//
// Solidity: function holdFrom(address sender, address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactor) HoldFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "holdFrom", sender, recipient, amount, holdId)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x3ca22cb1.
//
// Solidity: function holdFrom(address sender, address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractSession) HoldFrom(sender common.Address, recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.HoldFrom(&_HoldContract.TransactOpts, sender, recipient, amount, holdId)
}

// HoldFrom is a paid mutator transaction binding the contract method 0x3ca22cb1.
//
// Solidity: function holdFrom(address sender, address recipient, uint256 amount, uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) HoldFrom(sender common.Address, recipient common.Address, amount *big.Int, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.HoldFrom(&_HoldContract.TransactOpts, sender, recipient, amount, holdId)
}

// RemoveHold is a paid mutator transaction binding the contract method 0x08e53826.
//
// Solidity: function removeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactor) RemoveHold(opts *bind.TransactOpts, holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "removeHold", holdId)
}

// RemoveHold is a paid mutator transaction binding the contract method 0x08e53826.
//
// Solidity: function removeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractSession) RemoveHold(holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.RemoveHold(&_HoldContract.TransactOpts, holdId)
}

// RemoveHold is a paid mutator transaction binding the contract method 0x08e53826.
//
// Solidity: function removeHold(uint256 holdId) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) RemoveHold(holdId *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.RemoveHold(&_HoldContract.TransactOpts, holdId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Transfer(&_HoldContract.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.Transfer(&_HoldContract.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.TransferFrom(&_HoldContract.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool success)
func (_HoldContract *HoldContractTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _HoldContract.Contract.TransferFrom(&_HoldContract.TransactOpts, sender, recipient, amount)
}

// HoldContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the HoldContract contract.
type HoldContractApprovalIterator struct {
	Event *HoldContractApproval // Event containing the contract specifics and raw log

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
func (it *HoldContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoldContractApproval)
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
		it.Event = new(HoldContractApproval)
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
func (it *HoldContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoldContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoldContractApproval represents a Approval event raised by the HoldContract contract.
type HoldContractApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HoldContract *HoldContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*HoldContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HoldContract.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &HoldContractApprovalIterator{contract: _HoldContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HoldContract *HoldContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *HoldContractApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _HoldContract.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoldContractApproval)
				if err := _HoldContract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_HoldContract *HoldContractFilterer) ParseApproval(log types.Log) (*HoldContractApproval, error) {
	event := new(HoldContractApproval)
	if err := _HoldContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// HoldContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the HoldContract contract.
type HoldContractTransferIterator struct {
	Event *HoldContractTransfer // Event containing the contract specifics and raw log

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
func (it *HoldContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HoldContractTransfer)
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
		it.Event = new(HoldContractTransfer)
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
func (it *HoldContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HoldContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HoldContractTransfer represents a Transfer event raised by the HoldContract contract.
type HoldContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HoldContract *HoldContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*HoldContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HoldContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &HoldContractTransferIterator{contract: _HoldContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HoldContract *HoldContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *HoldContractTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _HoldContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HoldContractTransfer)
				if err := _HoldContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_HoldContract *HoldContractFilterer) ParseTransfer(log types.Log) (*HoldContractTransfer, error) {
	event := new(HoldContractTransfer)
	if err := _HoldContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
