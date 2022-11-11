package types

import "math/big"

type Transfer struct {
	Deposit big.Int
}

func TransferAction(deposit big.Int) Action {
    return Action{
        Enum: TransferEnum,
        Transfer: Transfer{
            Deposit: deposit,
        },
    }
}
