package contablock

import (
	"math/big"
)

type ContaBlockContractAdapter interface {
	Store(input StoreAdapterInput) error
	Retrieve() (*RetrieveAdapterOutput, error)
}

type StoreAdapterInput struct {
	Num *big.Int
}

type RetrieveAdapterOutput struct {
	Num *big.Int
}