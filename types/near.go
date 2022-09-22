package types

import "math/big"

const oneNearString = "1000000000000000000000000"

type Near struct {
    value *big.Int
}

func NewNear(amount int64) *Near {
    yotto, _ := big.NewInt(0).SetString(oneNearString, 10)
    yotto.Mul(yotto, big.NewInt(amount))
    return &Near{yotto}
}

func (n *Near) bigInt() *big.Int {
    return n.value
}
