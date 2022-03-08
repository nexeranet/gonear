package types

import "math/big"

type SignedTxFunctionCall struct {
	Transaction TxFunctionCall
	Signature   Signature
}

type TxFunctionCall struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    []FunctionCallAction
}

type FunctionCallAction struct {
	Enum         ActionEnum
	FunctionCall FunctionCall
}

type FunctionCall struct {
	MethodName string
	Args       []byte
	Gas        big.Int
	Deposit    big.Int
}
