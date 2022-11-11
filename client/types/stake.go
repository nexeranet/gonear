package types

import (
	"math/big"
)

type Stake struct {
	Stake     big.Int
	PublicKey PublicKey
}

func StakeAction(stake big.Int, key PublicKey) Action {
	return Action{
		Enum: StakeEnum,
		Stake: Stake{
			Stake:     stake,
			PublicKey: key,
		},
	}
}
