package contablock

import (
	"math/big"
)

type StoreInput struct {
	Num *big.Int
}

type RetrieveOutput struct {
	Num *big.Int
}
