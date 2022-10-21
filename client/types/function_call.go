package types

import "math/big"

type FunctionCall struct {
	MethodName string
	Args       []byte
	Gas        uint64
	Deposit    big.Int
}

func FunctionCallAction(name string, args []byte, gas uint64, deposit big.Int) Action {
	return Action{
		Enum: FunctionCallEnum,
		FunctionCall: FunctionCall{
			MethodName: name,
			Args:       args,
			Gas:        gas,
			Deposit:    deposit,
		},
	}
}
