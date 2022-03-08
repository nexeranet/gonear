package types

import "math/big"

type SignedTxTransfer struct {
	Transaction TxTransfer
	Signature   Signature
}

type TxTransfer struct {
	SignerId   string
	PublicKey  PublicKey
	Nonce      uint64
	ReceiverId string
	BlockHash  [32]byte
	Actions    []TransferAction
}

type Transfer struct {
	Deposit big.Int
}

type TransferAction struct {
	Enum     ActionEnum
	Transfer Transfer
}
