package holdContract

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type HoldContractInterface interface {
	BalanceOf(account common.Address) (*big.Int, error)
	BalanceOnHoldOf(account common.Address) (*big.Int, error)
	Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error)
}

type DefaultHoldContract struct {
	holdContract HoldContract
}

type DefaultHoldContractOptions struct {
	HoldContract HoldContract
}

/*
func NewDefaultHoldContract() (*DefaultHoldContract, error) {
	//todo
}*/

func (d DefaultHoldContract) BalanceOf(account common.Address) (*big.Int, error) {
	return d.BalanceOf(account)
}

func (d DefaultHoldContract) BalanceOnHoldOf(account common.Address) (*big.Int, error) {
	return d.BalanceOnHoldOf(account)
}

func (d DefaultHoldContract) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return d.Transfer(recipient, amount)
}
